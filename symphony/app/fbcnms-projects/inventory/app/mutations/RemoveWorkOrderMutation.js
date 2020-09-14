/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import RelayEnvironment from '../common/RelayEnvironment.js';
import {commitMutation, graphql} from 'react-relay';
import type {MutationCallbacks} from './MutationCallbacks.js';
import type {
  RemoveWorkOrderMutation,
  // $FlowFixMe (T62907961) Relay flow types
  RemoveWorkOrderMutationMutationResponse,
  // $FlowFixMe (T62907961) Relay flow types
  RemoveWorkOrderMutationMutationVariables,
} from './__generated__/RemoveWorkOrderMutation.graphql';
import type {StoreUpdater} from '../common/RelayEnvironment';

const mutation = graphql`
  mutation RemoveWorkOrderMutation($id: ID!) {
    removeWorkOrder(id: $id)
  }
`;

export default (
  variables: RemoveWorkOrderMutationMutationVariables,
  callbacks?: MutationCallbacks<RemoveWorkOrderMutationMutationResponse>,
  updater?: StoreUpdater,
) => {
  const {onCompleted, onError} = callbacks ? callbacks : {};
  commitMutation<RemoveWorkOrderMutation>(RelayEnvironment, {
    mutation,
    variables,
    updater,
    onCompleted,
    onError,
  });
};
