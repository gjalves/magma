#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass
from datetime import datetime
from gql.gql.datetime_utils import DATETIME_FIELD
from gql.gql.graphql_client import GraphqlClient
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import dataclass_json


@dataclass_json
@dataclass
class EquipmentTypeAndPropertiesQuery:
    __QUERY__ = """
    query EquipmentTypeAndPropertiesQuery($id: ID!) {
  equipment: node(id: $id) {
    ... on Equipment {
      equipmentType {
        name
      }
      properties {
        propertyType {
          id
        }
        stringValue
        intValue
        booleanValue
        floatValue
        latitudeValue
        longitudeValue
      }
    }
  }
}

    """

    @dataclass_json
    @dataclass
    class EquipmentTypeAndPropertiesQueryData:
        @dataclass_json
        @dataclass
        class Node:
            @dataclass_json
            @dataclass
            class EquipmentType:
                name: str

            @dataclass_json
            @dataclass
            class Property:
                @dataclass_json
                @dataclass
                class PropertyType:
                    id: str

                propertyType: PropertyType
                stringValue: Optional[str] = None
                intValue: Optional[int] = None
                booleanValue: Optional[bool] = None
                floatValue: Optional[Number] = None
                latitudeValue: Optional[Number] = None
                longitudeValue: Optional[Number] = None

            equipmentType: EquipmentType
            properties: List[Property]

        equipment: Optional[Node] = None

    data: Optional[EquipmentTypeAndPropertiesQueryData] = None
    errors: Optional[Any] = None

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, id: str):
        # fmt: off
        variables = {"id": id}
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data
