/**
 * @generated
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 **/

 /**
 * @flow
 */

/* eslint-disable */

'use strict';

/*::
import type { ReaderFragment } from 'relay-runtime';
type ServiceEndpointsView_endpoints$ref = any;
type ServiceLinksView_links$ref = any;
export type DiscoveryMethod = "INVENTORY" | "%future added value";
export type ServiceStatus = "DISCONNECTED" | "IN_SERVICE" | "MAINTENANCE" | "PENDING" | "%future added value";
import type { FragmentReference } from "relay-runtime";
declare export opaque type ServicePanel_service$ref: FragmentReference;
declare export opaque type ServicePanel_service$fragmentType: ServicePanel_service$ref;
export type ServicePanel_service = {|
  +id: string,
  +name: string,
  +externalId: ?string,
  +status: ServiceStatus,
  +customer: ?{|
    +name: string
  |},
  +serviceType: {|
    +name: string,
    +discoveryMethod: ?DiscoveryMethod,
    +endpointDefinitions: $ReadOnlyArray<{|
      +id: string,
      +name: string,
      +role: ?string,
      +equipmentType: {|
        +id: string,
        +name: string,
      |},
    |}>,
  |},
  +links: $ReadOnlyArray<?{|
    +id: string,
    +$fragmentRefs: ServiceLinksView_links$ref,
  |}>,
  +endpoints: $ReadOnlyArray<?{|
    +id: string,
    +definition: {|
      +id: string,
      +name: string,
    |},
    +$fragmentRefs: ServiceEndpointsView_endpoints$ref,
  |}>,
  +$refType: ServicePanel_service$ref,
|};
export type ServicePanel_service$data = ServicePanel_service;
export type ServicePanel_service$key = {
  +$data?: ServicePanel_service$data,
  +$fragmentRefs: ServicePanel_service$ref,
  ...
};
*/


const node/*: ReaderFragment*/ = (function(){
var v0 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "id",
  "args": null,
  "storageKey": null
},
v1 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "name",
  "args": null,
  "storageKey": null
},
v2 = [
  (v0/*: any*/),
  (v1/*: any*/)
];
return {
  "kind": "Fragment",
  "name": "ServicePanel_service",
  "type": "Service",
  "metadata": null,
  "argumentDefinitions": [],
  "selections": [
    (v0/*: any*/),
    (v1/*: any*/),
    {
      "kind": "ScalarField",
      "alias": null,
      "name": "externalId",
      "args": null,
      "storageKey": null
    },
    {
      "kind": "ScalarField",
      "alias": null,
      "name": "status",
      "args": null,
      "storageKey": null
    },
    {
      "kind": "LinkedField",
      "alias": null,
      "name": "customer",
      "storageKey": null,
      "args": null,
      "concreteType": "Customer",
      "plural": false,
      "selections": [
        (v1/*: any*/)
      ]
    },
    {
      "kind": "LinkedField",
      "alias": null,
      "name": "serviceType",
      "storageKey": null,
      "args": null,
      "concreteType": "ServiceType",
      "plural": false,
      "selections": [
        (v1/*: any*/),
        {
          "kind": "ScalarField",
          "alias": null,
          "name": "discoveryMethod",
          "args": null,
          "storageKey": null
        },
        {
          "kind": "LinkedField",
          "alias": null,
          "name": "endpointDefinitions",
          "storageKey": null,
          "args": null,
          "concreteType": "ServiceEndpointDefinition",
          "plural": true,
          "selections": [
            (v0/*: any*/),
            (v1/*: any*/),
            {
              "kind": "ScalarField",
              "alias": null,
              "name": "role",
              "args": null,
              "storageKey": null
            },
            {
              "kind": "LinkedField",
              "alias": null,
              "name": "equipmentType",
              "storageKey": null,
              "args": null,
              "concreteType": "EquipmentType",
              "plural": false,
              "selections": (v2/*: any*/)
            }
          ]
        }
      ]
    },
    {
      "kind": "LinkedField",
      "alias": null,
      "name": "links",
      "storageKey": null,
      "args": null,
      "concreteType": "Link",
      "plural": true,
      "selections": [
        (v0/*: any*/),
        {
          "kind": "FragmentSpread",
          "name": "ServiceLinksView_links",
          "args": null
        }
      ]
    },
    {
      "kind": "LinkedField",
      "alias": null,
      "name": "endpoints",
      "storageKey": null,
      "args": null,
      "concreteType": "ServiceEndpoint",
      "plural": true,
      "selections": [
        (v0/*: any*/),
        {
          "kind": "LinkedField",
          "alias": null,
          "name": "definition",
          "storageKey": null,
          "args": null,
          "concreteType": "ServiceEndpointDefinition",
          "plural": false,
          "selections": (v2/*: any*/)
        },
        {
          "kind": "FragmentSpread",
          "name": "ServiceEndpointsView_endpoints",
          "args": null
        }
      ]
    }
  ]
};
})();
// prettier-ignore
(node/*: any*/).hash = 'ff14440bf865c0fada807dcbe20fc346';
module.exports = node;
