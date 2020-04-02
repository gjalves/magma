#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass
from datetime import datetime
from gql.gql.datetime_utils import DATETIME_FIELD
from gql.gql.graphql_client import GraphqlClient
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import DataClassJsonMixin

from .equipment_port_type_fragment import EquipmentPortTypeFragment, QUERY as EquipmentPortTypeFragmentQuery
from .add_equipment_port_type_input import AddEquipmentPortTypeInput


@dataclass
class AddEquipmentPortTypeMutation(DataClassJsonMixin):
    @dataclass
    class AddEquipmentPortTypeMutationData(DataClassJsonMixin):
        @dataclass
        class EquipmentPortType(EquipmentPortTypeFragment):
            pass

        addEquipmentPortType: EquipmentPortType

    data: AddEquipmentPortTypeMutationData

    __QUERY__: str = EquipmentPortTypeFragmentQuery + """
    mutation AddEquipmentPortTypeMutation($input: AddEquipmentPortTypeInput!) {
  addEquipmentPortType(input: $input) {
    ...EquipmentPortTypeFragment
  }
}

    """

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, input: AddEquipmentPortTypeInput) -> AddEquipmentPortTypeMutationData:
        # fmt: off
        variables = {"input": input}
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data
