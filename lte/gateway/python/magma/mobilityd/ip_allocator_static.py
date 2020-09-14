"""
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

This is one of ip allocator for ip address manager.
The IP allocator accepts IP blocks (range of IP addresses), and supports
allocating and releasing IP addresses from the assigned IP blocks.
"""

from __future__ import absolute_import, division, print_function, \
    unicode_literals

from ipaddress import ip_address, ip_network
from typing import List

from magma.mobilityd.ip_descriptor import IPDesc, IPState, IPType
from magma.mobilityd.ip_allocator_base import IPAllocator
from magma.mobilityd.subscriberdb_client import SubscriberDbClient

DEFAULT_IP_RECYCLE_INTERVAL = 15


class IPAllocatorStaticWrapper(IPAllocator):

    def __init__(self, subscriberdb_rpc_stub, ip_allocator: IPAllocator):
        """ Initializes a static IP allocator
            This is wrapper around other configured Ip allocator. If subscriber
            does have static IP, it uses underlying IP allocator to allocate IP
            for the subscriber.
        """
        self._subscriber_client = SubscriberDbClient(subscriberdb_rpc_stub)
        self._ip_allocator = ip_allocator

    def add_ip_block(self, ipblock: ip_network):
        """ Add a block of IP addresses to the free IP list
        """
        self._ip_allocator.add_ip_block(ipblock)

    def remove_ip_blocks(self, ipblocks: List[ip_network],
                         _force: bool = False) -> List[ip_network]:
        """ Remove allocated IP blocks.
        """
        return self._ip_allocator.remove_ip_blocks(ipblocks, _force)

    def list_added_ip_blocks(self) -> List[ip_network]:
        """ List IP blocks added to the IP allocator
        Return:
             copy of the list of assigned IP blocks
        """
        return self._ip_allocator.list_added_ip_blocks()

    def list_allocated_ips(self, ipblock: ip_network) -> List[ip_address]:
        """ List IP addresses allocated from a given IP block
        """
        return self._ip_allocator.list_allocated_ips(ipblock)

    def alloc_ip_address(self, sid: str) -> IPDesc:
        """ Check if subscriber has static IP assigned.
        If it is not allocated use IP allocator to assign an IP.
        """
        ip_desc = self._allocate_static_ip(sid)
        if ip_desc is None:
            ip_desc = self._ip_allocator.alloc_ip_address(sid)
        return ip_desc

    def release_ip(self, ip_desc: IPDesc):
        """
        Statically allocated IPs do not need to do any update on
        ip release
        """
        if ip_desc.type != IPType.STATIC:
            self._ip_allocator.release_ip(ip_desc)

    def _allocate_static_ip(self, sid: str) -> IPDesc:
        """
        Check if static IP allocation is enabled and then check
        subscriber DB for assigned static IP for the SID
        """
        ip_addr = self._subscriber_client.get_subscriber_ip(sid)
        if ip_addr is None:
            return None
        return IPDesc(ip=ip_addr, state=IPState.ALLOCATED,
                      sid=sid, ip_block=ip_addr, ip_type=IPType.STATIC)

