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

from .add_equipment_input import AddEquipmentInput


@dataclass
class AddEquipmentMutation(DataClassJsonMixin):
    @dataclass
    class AddEquipmentMutationData(DataClassJsonMixin):
        @dataclass
        class Equipment(DataClassJsonMixin):
            id: str
            name: str

        addEquipment: Equipment

    data: AddEquipmentMutationData

    __QUERY__: str = """
    mutation AddEquipmentMutation($input: AddEquipmentInput!) {
  addEquipment(input: $input) {
    id
    name
  }
}

    """

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, input: AddEquipmentInput) -> AddEquipmentMutationData:
        # fmt: off
        variables = {"input": input}
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data
