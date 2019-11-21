"""
Copyright (c) 2019-present, Arcangelli.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""

from typing import Optional, Any, Callable, Dict, List, Type
from magma.common.service import MagmaService
from magma.enodebd.data_models.data_model import TrParam, DataModel
from magma.enodebd.data_models.data_model_parameters import ParameterName, \
    TrParameterType
from magma.enodebd.data_models import transform_for_magma, transform_for_enb
from magma.enodebd.device_config.enodeb_config_postprocessor import \
    EnodebConfigurationPostProcessor
from magma.enodebd.device_config.enodeb_configuration import \
    EnodebConfiguration
from magma.enodebd.devices.mikrotik_intercell_10 import \
    MikrotikIntercell10GetObjectParametersState, \
    MikrotikIntercell10WaitGetTransientParametersState
from magma.enodebd.devices.device_utils import EnodebDeviceName
from magma.enodebd.state_machines.enb_acs_impl import \
    BasicEnodebAcsStateMachine
from magma.enodebd.state_machines.enb_acs_states import \
    WaitInformState, SendGetTransientParametersState, \
    GetParametersState, WaitGetParametersState, DeleteObjectsState, \
    AddObjectsState, SetParameterValuesState, WaitSetParameterValuesState, \
    WaitRebootResponseState, WaitInformMRebootState, EnodebAcsState, \
    WaitEmptyMessageState, ErrorState, \
    EndSessionState, MikrotikIntercell10SendRebootState, GetRPCMethodsState


class MikrotikIntercell10Handler(BasicEnodebAcsStateMachine):
    def __init__(
        self,
        service: MagmaService,
    ) -> None:
        self._state_map = {}
        super().__init__(service)

    def reboot_asap(self) -> None:
        self.transition('reboot')

    def is_enodeb_connected(self) -> bool:
        return not isinstance(self.state, WaitInformState)

    def _init_state_map(self) -> None:
        self._state_map = {
            'wait_inform': WaitInformState(self, when_done='get_rpc_methods'),
            'get_rpc_methods': GetRPCMethodsState(self, when_done='wait_empty', when_skip='get_transient_params'),
            'wait_empty': WaitEmptyMessageState(self, when_done='get_transient_params'),
            'get_transient_params': SendGetTransientParametersState(self, when_done='wait_get_transient_params'),
            'wait_get_transient_params': MikrotikIntercell10WaitGetTransientParametersState(self, when_get='get_params', when_get_obj_params='get_obj_params', when_delete='delete_objs', when_add='add_objs', when_set='set_params', when_skip='end_session'),
            'get_params': GetParametersState(self, when_done='wait_get_params'),
            'wait_get_params': WaitGetParametersState(self, when_done='get_obj_params'),
            'get_obj_params': MikrotikIntercell10GetObjectParametersState(self, when_delete='delete_objs', when_add='add_objs', when_set='set_params', when_skip='end_session'),
            'delete_objs': DeleteObjectsState(self, when_add='add_objs', when_skip='set_params'),
            'add_objs': AddObjectsState(self, when_done='set_params'),
            'set_params': SetParameterValuesState(self, when_done='wait_set_params'),
            'wait_set_params': WaitSetParameterValuesState(self, when_done='check_get_params', when_apply_invasive='check_get_params'),
            'check_get_params': GetParametersState(self, when_done='check_wait_get_params', request_all_params=True),
            'check_wait_get_params': WaitGetParametersState(self, when_done='end_session'),
            'end_session': EndSessionState(self),
            # These states are only entered through manual user intervention
            'reboot': MikrotikIntercell10SendRebootState(self, when_done='wait_reboot'),
            'wait_reboot': WaitRebootResponseState(self, when_done='wait_post_reboot_inform'),
            'wait_post_reboot_inform': WaitInformMRebootState(self, when_done='wait_empty', when_timeout='wait_inform'),
            # The states below are entered when an unexpected message type is
            # received
            'unexpected_fault': ErrorState(self)
        }

    @property
    def device_name(self) -> str:
        return EnodebDeviceName.MIKROTIK_INTERCELL_10

    @property
    def data_model_class(self) -> Type[DataModel]:
        return MikrotikIntercell10TrDataModel

    @property
    def config_postprocessor(self) -> EnodebConfigurationPostProcessor:
        return MikrotikIntercell10TrConfigurationInitializer()

    @property
    def state_map(self) -> Dict[str, EnodebAcsState]:
        return self._state_map

    @property
    def disconnected_state_name(self) -> str:
        return 'wait_inform'

    @property
    def unexpected_fault_state_name(self) -> str:
        return 'unexpected_fault'


class MikrotikIntercell10TrDataModel(DataModel):
    """
    Class to represent relevant data model parameters from TR-196/TR-098.
    This class is effectively read-only.

    This model specifically targets Mikrotik Intercell 10 units.

    These models have these idiosyncrasies (on account of running TR098):

    - Parameter content root is different (InternetGatewayDevice)
    - GetParameter queries with a wildcard e.g. InternetGatewayDevice. do
      not respond with the full tree (we have to query all parameters)
    - MME status is not exposed - we assume the MME is connected if
      the eNodeB is transmitting (OpState=true)
    - Parameters such as band capability/duplex config
      are rooted under `boardconf.` and not the device config root
    - Parameters like Admin state, CellReservedForOperatorUse,
      Duplex mode, DL bandwidth and Band capability have different
      formats from Intel-based Mikrotik units, necessitating,
      formatting before configuration and transforming values
      read from eNodeB state.
    - Num PLMNs is not reported by these units
    """
    # Mapping of TR parameter paths to aliases
    DEVICE_PATH = 'Device.'
    FAP_PATH = DEVICE_PATH + 'FAP.'
    FAPSERVICE_PATH = DEVICE_PATH + 'Services.FAPService.1.'
    EEPROM_PATH = 'boardconf.status.eepromInfo.'
    PARAMETERS = {
        # Top-level objects
        ParameterName.DEVICE: TrParam(DEVICE_PATH, True, TrParameterType.OBJECT, False),
        ParameterName.FAP_SERVICE: TrParam(FAPSERVICE_PATH, True, TrParameterType.OBJECT, False),

        # General information
        ParameterName.MME_STATUS: TrParam(FAPSERVICE_PATH + 'FAPControl.LTE.OpState', True, TrParameterType.BOOLEAN, False),
        ParameterName.GPS_LAT: TrParam(FAP_PATH + 'GPS.ContinuousGPSStatus.Latitude', True, TrParameterType.STRING, False),
        ParameterName.GPS_LONG: TrParam(FAP_PATH + 'GPS.ContinuousGPSStatus.Longitude', True, TrParameterType.STRING, False),
        ParameterName.SW_VERSION: TrParam(DEVICE_PATH + 'DeviceInfo.SoftwareVersion', True, TrParameterType.STRING, False),
        ParameterName.SERIAL_NUMBER: TrParam(DEVICE_PATH + 'DeviceInfo.SerialNumber', True, TrParameterType.STRING, False), # Missing

        # Capabilities
        ParameterName.DUPLEX_MODE_CAPABILITY: TrParam(EEPROM_PATH + 'div_multiple', True, TrParameterType.STRING, False), # Missing
        ParameterName.BAND_CAPABILITY: TrParam(EEPROM_PATH + 'work_mode', True, TrParameterType.STRING, False), # Missing

        # RF-related parameters
        ParameterName.EARFCNDL: TrParam(FAPSERVICE_PATH + 'CellConfig.LTE.RAN.RF.EARFCNDL', True, TrParameterType.UNSIGNED_INT, False),
        ParameterName.EARFCNUL: TrParam(FAPSERVICE_PATH + 'CellConfig.LTE.RAN.RF.EARFCNUL', True, TrParameterType.UNSIGNED_INT, False),
        ParameterName.BAND: TrParam(FAPSERVICE_PATH + 'CellConfig.LTE.RAN.RF.FreqBandIndicator', True, TrParameterType.UNSIGNED_INT, False),
        ParameterName.PCI: TrParam(FAPSERVICE_PATH + 'CellConfig.LTE.RAN.RF.PhyCellID', True, TrParameterType.STRING, False),
        ParameterName.DL_BANDWIDTH: TrParam(DEVICE_PATH + 'Services.RfConfig.1.RfCarrierCommon.carrierBwMhz', True, TrParameterType.INT, False), # Missing
        ParameterName.CELL_ID: TrParam(FAPSERVICE_PATH + 'CellConfig.LTE.RAN.Common.CellIdentity', True, TrParameterType.UNSIGNED_INT, False),
        ParameterName.SUBFRAME_ASSIGNMENT: TrParam(FAPSERVICE_PATH + 'CellConfig.LTE.RAN.PHY.TDDFrame.SubFrameAssignment', True, TrParameterType.INT, False),
        ParameterName.SPECIAL_SUBFRAME_PATTERN: TrParam(FAPSERVICE_PATH + 'CellConfig.LTE.RAN.PHY.TDDFrame.SpecialSubframePatterns', True, TrParameterType.INT, False),

        # Other LTE parameters
        ParameterName.ADMIN_STATE: TrParam(FAPSERVICE_PATH + 'FAPControl.LTE.AdminState', False, TrParameterType.UNSIGNED_INT, False),
        ParameterName.OP_STATE: TrParam(FAPSERVICE_PATH + 'FAPControl.LTE.OpState', True, TrParameterType.UNSIGNED_INT, False),
        ParameterName.RF_TX_STATUS: TrParam(FAPSERVICE_PATH + 'FAPControl.LTE.RFTxStatus', True, TrParameterType.UNSIGNED_INT, False),

        # Core network parameters
        ParameterName.MME_IP: TrParam(FAPSERVICE_PATH + 'FAPControl.LTE.Gateway.S1SigLinkServerList', True, TrParameterType.STRING, False),
        ParameterName.MME_PORT: TrParam(FAPSERVICE_PATH + 'FAPControl.LTE.Gateway.S1SigLinkPort', True, TrParameterType.UNSIGNED_INT, False),
        ParameterName.NUM_PLMNS: TrParam(FAPSERVICE_PATH + 'CellConfig.LTE.EPC.PLMNListNumberOfEntries', True, TrParameterType.UNSIGNED_INT, False),
        ParameterName.TAC: TrParam(FAPSERVICE_PATH + 'CellConfig.LTE.EPC.TAC', True, TrParameterType.UNSIGNED_INT, False),
        ParameterName.IP_SEC_ENABLE: TrParam(DEVICE_PATH + 'IPsec.Enable', False, TrParameterType.UNSIGNED_INT, False),

        # Management server parameters
        ParameterName.PERIODIC_INFORM_ENABLE: TrParam(DEVICE_PATH + 'ManagementServer.PeriodicInformEnable', False, TrParameterType.UNSIGNED_INT, False),
        ParameterName.PERIODIC_INFORM_INTERVAL: TrParam(DEVICE_PATH + 'ManagementServer.PeriodicInformInterval', False, TrParameterType.INT, False),

        # Performance management parameters
        ParameterName.PERF_MGMT_ENABLE: TrParam(FAP_PATH + 'PerfMgmt.Config.1.Enable', False, TrParameterType.UNSIGNED_INT, False),
        ParameterName.PERF_MGMT_UPLOAD_INTERVAL: TrParam(FAP_PATH + 'PerfMgmt.Config.1.PeriodicUploadInterval', False, TrParameterType.UNSIGNED_INT, False),
        ParameterName.PERF_MGMT_UPLOAD_URL: TrParam(FAP_PATH + 'PerfMgmt.Config.1.URL', False, TrParameterType.STRING, False),
        ParameterName.PERF_MGMT_USER: TrParam(FAP_PATH + 'PerfMgmt.Config.1.Username', False, TrParameterType.STRING, False),
        ParameterName.PERF_MGMT_PASSWORD: TrParam(FAP_PATH + 'PerfMgmt.Config.1.Password', False, TrParameterType.STRING, False),
    }

    NUM_PLMNS_IN_CONFIG = 6
    TRANSFORMS_FOR_ENB = {
        ParameterName.CELL_BARRED: transform_for_enb.invert_cell_barred,
    }
    for i in range(1, NUM_PLMNS_IN_CONFIG + 1):
        TRANSFORMS_FOR_ENB[ParameterName.PLMN_N_CELL_RESERVED % i] = transform_for_enb.cell_reserved
        PARAMETERS[ParameterName.PLMN_N % i] = TrParam(FAPSERVICE_PATH + 'CellConfig.LTE.EPC.PLMNList.%d.' % i, True, TrParameterType.STRING, False)
        PARAMETERS[ParameterName.PLMN_N_CELL_RESERVED % i] = TrParam(FAPSERVICE_PATH + 'CellConfig.LTE.EPC.PLMNList.%d.CellReservedForOperatorUse' % i, True, TrParameterType.STRING, False)
        PARAMETERS[ParameterName.PLMN_N_ENABLE % i] = TrParam(FAPSERVICE_PATH + 'CellConfig.LTE.EPC.PLMNList.%d.Enable' % i, True, TrParameterType.BOOLEAN, False)
        PARAMETERS[ParameterName.PLMN_N_PRIMARY % i] = TrParam(FAPSERVICE_PATH + 'CellConfig.LTE.EPC.PLMNList.%d.IsPrimary' % i, True, TrParameterType.BOOLEAN, False)
        PARAMETERS[ParameterName.PLMN_N_PLMNID % i] = TrParam(FAPSERVICE_PATH + 'CellConfig.LTE.EPC.PLMNList.%d.PLMNID' % i, True, TrParameterType.STRING, False)

    TRANSFORMS_FOR_ENB[ParameterName.ADMIN_STATE] = transform_for_enb.admin_state
    TRANSFORMS_FOR_MAGMA = {
        # We don't set these parameters
        ParameterName.BAND_CAPABILITY: transform_for_magma.band_capability,
        ParameterName.DUPLEX_MODE_CAPABILITY: transform_for_magma.duplex_mode
    }

    @classmethod
    def get_parameter(cls, param_name: ParameterName) -> Optional[TrParam]:
        return cls.PARAMETERS.get(param_name)

    @classmethod
    def _get_magma_transforms(
        cls,
    ) -> Dict[ParameterName, Callable[[Any], Any]]:
        return cls.TRANSFORMS_FOR_MAGMA

    @classmethod
    def _get_enb_transforms(cls) -> Dict[ParameterName, Callable[[Any], Any]]:
        return cls.TRANSFORMS_FOR_ENB

    @classmethod
    def get_load_parameters(cls) -> List[ParameterName]:
        """
        Load all the parameters instead of a subset.
        """
        return list(cls.PARAMETERS.keys())

    @classmethod
    def get_num_plmns(cls) -> int:
        return cls.NUM_PLMNS_IN_CONFIG

    @classmethod
    def get_parameter_names(cls) -> List[ParameterName]:
        excluded_params = [str(ParameterName.DEVICE),
                           str(ParameterName.FAP_SERVICE)]
        names = list(filter(lambda x: (not str(x).startswith('PLMN'))
                                      and (str(x) not in excluded_params),
                            cls.PARAMETERS.keys()))
        return names

    @classmethod
    def get_numbered_param_names(
        cls,
    ) -> Dict[ParameterName, List[ParameterName]]:
        names = {}
        for i in range(1, cls.NUM_PLMNS_IN_CONFIG + 1):
            params = []
            params.append(ParameterName.PLMN_N_CELL_RESERVED % i)
            params.append(ParameterName.PLMN_N_ENABLE % i)
            params.append(ParameterName.PLMN_N_PRIMARY % i)
            params.append(ParameterName.PLMN_N_PLMNID % i)
            names[ParameterName.PLMN_N % i] = params

        return names


class MikrotikIntercell10TrConfigurationInitializer(EnodebConfigurationPostProcessor):
    def postprocess(self, desired_cfg: EnodebConfiguration) -> None:

        desired_cfg.delete_parameter(ParameterName.ADMIN_STATE)
