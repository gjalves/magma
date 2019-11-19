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
from magma.enodebd.devices.baicells_qafb import \
    BaicellsQafbGetObjectParametersState, \
    BaicellsQafbWaitGetTransientParametersState
from magma.enodebd.devices.device_utils import EnodebDeviceName
from magma.enodebd.state_machines.enb_acs_impl import \
    BasicEnodebAcsStateMachine
from magma.enodebd.state_machines.enb_acs_states import \
    WaitInformState, SendGetTransientParametersState, \
    GetParametersState, WaitGetParametersState, DeleteObjectsState, \
    AddObjectsState, SetParameterValuesState, WaitSetParameterValuesState, \
    WaitRebootResponseState, WaitInformMRebootState, EnodebAcsState, \
    WaitEmptyMessageState, ErrorState, \
    EndSessionState, BaicellsSendRebootState, GetRPCMethodsState

class MikrotikIntercell10Handler(BasicEnodebAcsStateMachine):
    def __init__(
        self,
        service: MagmaService,
    ) -> None:
        self._state_map = {}
        super().__init__(service)

    def reboot_asap(self) -> None:
        self.transition('reboot')
