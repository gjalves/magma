#!/usr/bin/env python3
# pyre-strict

from dataclasses import asdict
from typing import Dict, List, Optional, Tuple

from dacite import Config, from_dict

from .._utils import PropertyValue, format_properties, get_graphql_property_inputs
from ..consts import (
    Customer,
    EquipmentPort,
    Link,
    Service,
    ServiceEndpoint,
    ServiceType,
)
from ..graphql.add_service_endpoint_input import AddServiceEndpointInput
from ..graphql.add_service_endpoint_mutation import AddServiceEndpointMutation
from ..graphql.add_service_link_mutation import AddServiceLinkMutation
from ..graphql.add_service_mutation import AddServiceMutation
from ..graphql.add_service_type_mutation import AddServiceTypeMutation
from ..graphql.property_type_input import PropertyTypeInput
from ..graphql.remove_service_mutation import RemoveServiceMutation
from ..graphql.remove_service_type_mutation import RemoveServiceTypeMutation
from ..graphql.service_create_data_input import ServiceCreateData
from ..graphql.service_details_query import ServiceDetailsQuery
from ..graphql.service_endpoint_role_enum import ServiceEndpointRole
from ..graphql.service_status_enum import ServiceStatus
from ..graphql.service_type_create_data_input import ServiceTypeCreateData
from ..graphql.service_type_services_query import ServiceTypeServicesQuery
from ..graphql.service_types_query import ServiceTypesQuery
from ..graphql_client import GraphqlClient


def _populate_service_types(client: GraphqlClient) -> None:
    edges = ServiceTypesQuery.execute(client).serviceTypes.edges
    for edge in edges:
        node = edge.node
        client.serviceTypes[node.name] = ServiceType(
            name=node.name,
            id=node.id,
            hasCustomer=node.hasCustomer,
            propertyTypes=[asdict(p) for p in node.propertyTypes],
        )


def add_service_type(
    client: GraphqlClient,
    name: str,
    hasCustomer: bool,
    properties: List[Tuple[str, str, PropertyValue, bool]],
) -> ServiceType:

    new_property_types = format_properties(properties)
    result = AddServiceTypeMutation.execute(
        client,
        data=ServiceTypeCreateData(
            name=name,
            hasCustomer=hasCustomer,
            properties=[
                from_dict(
                    data_class=PropertyTypeInput, data=p, config=Config(strict=True)
                )
                for p in new_property_types
            ],
        ),
    ).addServiceType

    service_type = ServiceType(
        name=result.name,
        id=result.id,
        hasCustomer=result.hasCustomer,
        propertyTypes=[asdict(p) for p in result.propertyTypes],
    )
    client.serviceTypes[name] = service_type
    return service_type


def add_service(
    client: GraphqlClient,
    name: str,
    external_id: str,
    service_type: str,
    customer: Optional[Customer],
    properties_dict: Dict[str, PropertyValue],
    links: List[Link],
) -> Service:
    property_types = client.serviceTypes[service_type].propertyTypes
    properties = get_graphql_property_inputs(property_types, properties_dict)
    service_create_data = ServiceCreateData(
        name=name,
        externalId=external_id,
        serviceTypeId=client.serviceTypes[service_type].id,
        status=ServiceStatus.PENDING,
        customerId=customer.id if customer is not None else None,
        properties=properties,
        upstreamServiceIds=[],
    )
    result = AddServiceMutation.execute(client, data=service_create_data).addService
    for l in links:
        result = AddServiceLinkMutation.execute(
            client, id=result.id, linkId=l.id
        ).addServiceLink
    return Service(
        name=result.name,
        id=result.id,
        externalId=result.externalId,
        customer=Customer(
            name=result.customer.name,
            id=result.customer.id,
            externalId=result.customer.externalId,
        )
        if result.customer is not None
        else None,
        endpoints=[
            ServiceEndpoint(id=e.id, port=EquipmentPort(id=e.port.id), role=e.role)
            for e in result.endpoints
        ],
        links=[Link(id=l.id) for l in result.links],
    )


def add_service_endpoint(
    client: GraphqlClient,
    service: Service,
    port: EquipmentPort,
    role: ServiceEndpointRole,
) -> None:
    AddServiceEndpointMutation.execute(
        client, input=AddServiceEndpointInput(id=service.id, portId=port.id, role=role)
    )


def get_service(client: GraphqlClient, id: str) -> Service:
    result = ServiceDetailsQuery.execute(client, id=id).service
    return Service(
        name=result.name,
        id=result.id,
        externalId=result.externalId,
        customer=Customer(
            name=result.customer.name,
            id=result.customer.id,
            externalId=result.customer.externalId,
        )
        if result.customer is not None
        else None,
        endpoints=[
            ServiceEndpoint(id=e.id, port=EquipmentPort(id=e.port.id), role=e.role)
            for e in result.endpoints
        ],
        links=[Link(id=l.id) for l in result.links],
    )


def delete_service_type_with_services(
    client: GraphqlClient, service_type: ServiceType
) -> None:
    services = ServiceTypeServicesQuery.execute(
        client, id=service_type.id
    ).serviceType.services
    for service in services:
        RemoveServiceMutation.execute(client, id=service.id)
    RemoveServiceTypeMutation.execute(client, id=service_type.id)
