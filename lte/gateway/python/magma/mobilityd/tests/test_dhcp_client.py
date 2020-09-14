"""
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""
import datetime
import logging
import os
import subprocess
import sys
import threading
import time
import unittest
from freezegun import freeze_time

from magma.pipelined.bridge_util import BridgeTools

from magma.mobilityd.dhcp_client import DHCPClient
from magma.mobilityd.mac import MacAddress
from magma.mobilityd.dhcp_desc import DHCPState, DHCPDescriptor
from magma.mobilityd.uplink_gw import UplinkGatewayInfo
from scapy.layers.dhcp import DHCP
from scapy.layers.l2 import Ether
from scapy.sendrecv import AsyncSniffer

LOG = logging.getLogger('mobilityd.dhcp.test')
LOG.isEnabledFor(logging.DEBUG)

logging.basicConfig(stream=sys.stderr, level=logging.DEBUG)
SCRIPT_PATH = "/home/vagrant/magma/lte/gateway/python/magma/mobilityd/"
DHCP_IFACE = "cl1_dhcp0"
PKT_CAPTURE_WAIT = 2

"""
Test dhclient class independent of IP allocator.
"""


class DhcpClient(unittest.TestCase):
    def setUp(self):
        self._br = "dh_br0"
        try:
            subprocess.check_call(["pkill", "dnsmasq"])
        except subprocess.CalledProcessError:
            pass

        setup_dhcp_server = SCRIPT_PATH + "scripts/setup-test-dhcp-srv.sh"
        subprocess.check_call([setup_dhcp_server, "cl1"])

        setup_uplink_br = [SCRIPT_PATH + "scripts/setup-uplink-br.sh",
                           self._br,
                           "cl1uplink_p0",
                           DHCP_IFACE]
        subprocess.check_call(setup_uplink_br)

        self.dhcp_wait = threading.Condition()
        self.dhcp_store = {}
        self.gw_info_map = {}
        self.gw_info = UplinkGatewayInfo(self.gw_info_map)
        self._dhcp_client = DHCPClient(dhcp_wait=self.dhcp_wait,
                                       dhcp_store=self.dhcp_store,
                                       gw_info=self.gw_info,
                                       iface=DHCP_IFACE,
                                       lease_renew_wait_min=1)
        self._dhcp_client.run()

    def tearDown(self):
        self._dhcp_client.stop()
        BridgeTools.destroy_bridge(self._br)

    @unittest.skipIf(os.getuid(), reason="needs root user")
    def test_dhcp_lease1(self):
        self.pkt_list_lock = threading.Condition()

        self._setup_sniffer()
        mac1 = MacAddress("11:22:33:44:55:66")
        self._alloc_ip_address_from_dhcp(mac1)
        self._validate_req_state(mac1, DHCPState.REQUEST)
        self._validate_state_as_current(mac1)

        # trigger lease reneval before deadline
        time1 = datetime.datetime.now() + datetime.timedelta(seconds=100)
        self._start_sniffer()
        with freeze_time(time1):
            self._stop_sniffer_and_check(DHCPState.REQUEST, mac1)
            self._validate_req_state(mac1, DHCPState.REQUEST)
            self._validate_state_as_current(mac1)

            # trigger lease after deadline
            time2 = datetime.datetime.now() + datetime.timedelta(seconds=200)
            self._start_sniffer()
            with freeze_time(time2):
                LOG.debug("check discover after lease loss")
                self._stop_sniffer_and_check(DHCPState.DISCOVER, mac1)
                self._validate_req_state(mac1, DHCPState.REQUEST)
                self._validate_state_as_current(mac1)

        self._dhcp_client.release_ip_address(mac1)
        time.sleep(PKT_CAPTURE_WAIT)
        self._validate_req_state(mac1, DHCPState.RELEASE)

    def _validate_req_state(self, mac: MacAddress, state: DHCPState):
        with self.dhcp_wait:
            dhcp1 = self.dhcp_store.get(mac.as_redis_key())
            self.assertEqual(dhcp1.state_requested, state)

    def _validate_state_as_current(self, mac: MacAddress):
        with self.dhcp_wait:
            dhcp1 = self.dhcp_store.get(mac.as_redis_key())
            assert (dhcp1.state == DHCPState.OFFER or dhcp1.state == DHCPState.ACK)

    def _alloc_ip_address_from_dhcp(self, mac: MacAddress) -> DHCPDescriptor:
        retry_count = 0
        with self.dhcp_wait:
            dhcp_desc = None
            while (retry_count < 60 and (dhcp_desc is None or
                                         dhcp_desc.ip_is_allocated() is not True)):
                if retry_count % 5 == 0:
                    self._dhcp_client.send_dhcp_packet(mac, DHCPState.DISCOVER)

                self.dhcp_wait.wait(timeout=1)
                dhcp_desc = self._dhcp_client.get_dhcp_desc(mac)
                retry_count = retry_count + 1

            return dhcp_desc

    def _handle_dhcp_req_packet(self, packet):
        if DHCP not in packet:
            return
        with self.pkt_list_lock:
            self.pkt_list.append(packet)

    def _setup_sniffer(self):
        self._sniffer = AsyncSniffer(iface=DHCP_IFACE,
                                     filter="udp and (port 67 or 68)",
                                     prn=self._handle_dhcp_req_packet)

    def _start_sniffer(self):
        self.pkt_list = []
        self._sniffer.start()
        time.sleep(PKT_CAPTURE_WAIT)

    def _stop_sniffer_and_check(self, state: DHCPState, mac: MacAddress):
        for x in range(30):
            time.sleep(PKT_CAPTURE_WAIT)
            with self.pkt_list_lock:
                for pkt in self.pkt_list:
                    LOG.debug("DHCP pkt %s", pkt.show(dump=True))
                    if DHCP in pkt:
                        if pkt[DHCP].options[0][1] == int(state) and \
                                pkt[Ether].src == str(mac):
                            self._sniffer.stop()
                            return None

        assert 0
