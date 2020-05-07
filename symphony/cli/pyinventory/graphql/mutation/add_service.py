#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass
from datetime import datetime
from gql.gql.datetime_utils import DATETIME_FIELD
from gql.gql.graphql_client import GraphqlClient
from gql.gql.client import OperationException
from gql.gql.reporter import FailedOperationException
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional
from time import perf_counter
from dataclasses_json import DataClassJsonMixin

from ..fragment.customer import CustomerFragment, QUERY as CustomerFragmentQuery
from ..fragment.link import LinkFragment, QUERY as LinkFragmentQuery
from ..fragment.property import PropertyFragment, QUERY as PropertyFragmentQuery
from ..input.service_create_data import ServiceCreateData


QUERY: List[str] = CustomerFragmentQuery + LinkFragmentQuery + PropertyFragmentQuery + ["""
mutation AddServiceMutation($data: ServiceCreateData!) {
  addService(data: $data) {
    id
    name
    externalId
    customer {
      ...CustomerFragment
    }
    endpoints {
      id
      port {
        id
        properties {
          ...PropertyFragment
        }
        definition {
          id
          name
          portType {
            id
            name
          }
        }
        link {
          ...LinkFragment
        }
      }
      definition {
        role
      }
    }
    links {
      ...LinkFragment
    }
  }
}

"""]

@dataclass
class AddServiceMutation(DataClassJsonMixin):
    @dataclass
    class AddServiceMutationData(DataClassJsonMixin):
        @dataclass
        class Service(DataClassJsonMixin):
            @dataclass
            class Customer(CustomerFragment):
                pass

            @dataclass
            class ServiceEndpoint(DataClassJsonMixin):
                @dataclass
                class EquipmentPort(DataClassJsonMixin):
                    @dataclass
                    class Property(PropertyFragment):
                        pass

                    @dataclass
                    class EquipmentPortDefinition(DataClassJsonMixin):
                        @dataclass
                        class EquipmentPortType(DataClassJsonMixin):
                            id: str
                            name: str

                        id: str
                        name: str
                        portType: Optional[EquipmentPortType]

                    @dataclass
                    class Link(LinkFragment):
                        pass

                    id: str
                    properties: List[Property]
                    definition: EquipmentPortDefinition
                    link: Optional[Link]

                @dataclass
                class ServiceEndpointDefinition(DataClassJsonMixin):
                    role: Optional[str]

                id: str
                definition: ServiceEndpointDefinition
                port: Optional[EquipmentPort]

            @dataclass
            class Link(LinkFragment):
                pass

            id: str
            name: str
            endpoints: List[ServiceEndpoint]
            links: List[Link]
            externalId: Optional[str]
            customer: Optional[Customer]

        addService: Service

    data: AddServiceMutationData

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, data: ServiceCreateData) -> AddServiceMutationData.Service:
        # fmt: off
        variables = {"data": data}
        try:
            start_time = perf_counter()
            response_text = client.call(''.join(set(QUERY)), variables=variables)
            res = cls.from_json(response_text).data
            elapsed_time = perf_counter() - start_time
            client.reporter.log_successful_operation("AddServiceMutation", variables, elapsed_time)
            return res.addService
        except OperationException as e:
            raise FailedOperationException(
                client.reporter,
                e.err_msg,
                e.err_id,
                "AddServiceMutation",
                variables,
            )
