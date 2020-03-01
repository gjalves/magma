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

from gql.gql.enum_utils import enum_field
from .service_endpoint_role_enum import ServiceEndpointRole


@dataclass_json
@dataclass
class AddServiceLinkMutation:
    __QUERY__ = """
    mutation AddServiceLinkMutation($id: ID!, $linkId: ID!) {
  addServiceLink(id: $id, linkId: $linkId) {
    id
    name
    externalId
    customer {
      id
      name
      externalId
    }
    endpoints {
      id
      port {
        id
      }
      role
    }
    links {
      id
    }
  }
}

    """

    @dataclass_json
    @dataclass
    class AddServiceLinkMutationData:
        @dataclass_json
        @dataclass
        class Service:
            @dataclass_json
            @dataclass
            class Customer:
                id: str
                name: str
                externalId: Optional[str] = None

            @dataclass_json
            @dataclass
            class ServiceEndpoint:
                @dataclass_json
                @dataclass
                class EquipmentPort:
                    id: str

                id: str
                port: EquipmentPort
                role: ServiceEndpointRole = enum_field(ServiceEndpointRole)

            @dataclass_json
            @dataclass
            class Link:
                id: str

            id: str
            name: str
            endpoints: List[ServiceEndpoint]
            links: List[Link]
            externalId: Optional[str] = None
            customer: Optional[Customer] = None

        addServiceLink: Optional[Service] = None

    data: Optional[AddServiceLinkMutationData] = None
    errors: Optional[Any] = None

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, id: str, linkId: str):
        # fmt: off
        variables = {"id": id, "linkId": linkId}
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data
