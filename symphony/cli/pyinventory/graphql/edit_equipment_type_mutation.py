#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!
# pyre-strict

from dataclasses import dataclass
from datetime import datetime
from gql.gql.datetime_utils import DATETIME_FIELD
from gql.gql.graphql_client import GraphqlClient
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import DataClassJsonMixin

from gql.gql.enum_utils import enum_field
from .property_kind_enum import PropertyKind

from .edit_equipment_type_input import EditEquipmentTypeInput


@dataclass
class EditEquipmentTypeMutation(DataClassJsonMixin):
    @dataclass
    class EditEquipmentTypeMutationData(DataClassJsonMixin):
        @dataclass
        class EquipmentType(DataClassJsonMixin):
            @dataclass
            class PropertyType(DataClassJsonMixin):
                id: str
                name: str
                type: PropertyKind = enum_field(PropertyKind)
                index: Optional[int] = None
                stringValue: Optional[str] = None
                intValue: Optional[int] = None
                booleanValue: Optional[bool] = None
                floatValue: Optional[Number] = None
                latitudeValue: Optional[Number] = None
                longitudeValue: Optional[Number] = None
                isEditable: Optional[bool] = None
                isInstanceProperty: Optional[bool] = None

            @dataclass
            class EquipmentPositionDefinition(DataClassJsonMixin):
                id: str
                name: str
                index: Optional[int] = None
                visibleLabel: Optional[str] = None

            @dataclass
            class EquipmentPortDefinition(DataClassJsonMixin):
                id: str
                name: str
                index: Optional[int] = None
                visibleLabel: Optional[str] = None

            id: str
            name: str
            propertyTypes: List[PropertyType]
            positionDefinitions: List[EquipmentPositionDefinition]
            portDefinitions: List[EquipmentPortDefinition]
            category: Optional[str] = None

        editEquipmentType: EquipmentType

    data: EditEquipmentTypeMutationData

    __QUERY__: str = """
    mutation EditEquipmentTypeMutation($input: EditEquipmentTypeInput!) {
  editEquipmentType(input: $input) {
    id
    name
    category
    propertyTypes {
      id
      name
      type
      index
      stringValue
      intValue
      booleanValue
      floatValue
      latitudeValue
      longitudeValue
      isEditable
      isInstanceProperty
    }
    positionDefinitions {
      id
      name
      index
      visibleLabel
    }
    portDefinitions {
      id
      name
      index
      visibleLabel
    }
  }
}

    """

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, input: EditEquipmentTypeInput) -> EditEquipmentTypeMutationData:
        # fmt: off
        variables = {"input": input}
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data
