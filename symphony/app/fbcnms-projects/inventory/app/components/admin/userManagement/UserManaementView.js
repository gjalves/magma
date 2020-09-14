/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */
import type {NavigatableView} from '@fbcnms/ui/components/design-system/View/NavigatableViews';

import * as React from 'react';
import AppContext from '@fbcnms/ui/context/AppContext';
import Button from '@fbcnms/ui/components/design-system/Button';
import InventorySuspense from '../../../common/InventorySuspense';
import NavigatableViews from '@fbcnms/ui/components/design-system/View/NavigatableViews';
import NewUserDialog from './users/NewUserDialog';
import PermissionsGroupCard from './groups/PermissionsGroupCard';
import PermissionsGroupsView, {
  PERMISSION_GROUPS_VIEW_NAME,
} from './groups/PermissionsGroupsView';
import PermissionsPoliciesView, {
  PERMISSION_POLICIES_VIEW_NAME,
} from './policies/PermissionsPoliciesView';
import PermissionsPolicyCard from './policies/PermissionsPolicyCard';
import PopoverMenu from '@fbcnms/ui/components/design-system/Select/PopoverMenu';
import Strings from '@fbcnms/strings/Strings';
import UsersView from './users/UsersView';
import fbt from 'fbt';
import {ALL_USERS_PATH_PARAM, USER_PATH_PARAM} from './users/UsersTable';
import {DialogShowingContextProvider} from '@fbcnms/ui/components/design-system/Dialog/DialogShowingContext';
import {FormContextProvider} from '../../../common/FormContext';
import {NEW_DIALOG_PARAM, POLICY_TYPES} from './utils/UserManagementUtils';
import {useCallback, useContext, useMemo, useState} from 'react';
import {useHistory, useRouteMatch} from 'react-router-dom';

const USERS_HEADER = fbt(
  'Users and Roles',
  'Header for view showing system users settings',
);

const UserManaementForm = () => {
  const history = useHistory();
  const match = useRouteMatch();
  const basePath = match.path;
  const [addingNewUser, setAddingNewUser] = useState(false);
  const gotoGroupsPage = useCallback(() => history.push(`${basePath}/groups`), [
    history,
    basePath,
  ]);
  const gotoPoliciesPage = useCallback(
    () => history.push(`${basePath}/policies`),
    [history, basePath],
  );

  const {isFeatureEnabled} = useContext(AppContext);
  const permissionPoliciesMode = isFeatureEnabled('permission_policies');

  const VIEWS: Array<NavigatableView> = useMemo(() => {
    const views = [
      {
        routingPath: `users/${USER_PATH_PARAM}`,
        targetPath: `users/${ALL_USERS_PATH_PARAM}`,
        menuItem: {
          label: USERS_HEADER,
          tooltip: `${USERS_HEADER}`,
        },
        component: {
          header: {
            title: `${USERS_HEADER}`,
            subtitle:
              'Add and manage users by entering their details and selecting a role.',
            actionButtons: [
              <Button onClick={() => setAddingNewUser(true)}>
                <fbt desc="">Add User</fbt>
              </Button>,
            ],
          },
          children: <UsersView />,
        },
      },
      {
        routingPath: 'groups',
        menuItem: {
          label: PERMISSION_GROUPS_VIEW_NAME,
          tooltip: `${PERMISSION_GROUPS_VIEW_NAME}`,
        },
        component: {
          header: {
            title: `${PERMISSION_GROUPS_VIEW_NAME}`,
            subtitle: 'Create and manage groups and apply policies to them.',
            actionButtons: permissionPoliciesMode
              ? [
                  <Button
                    onClick={() => history.push(`group/${NEW_DIALOG_PARAM}`)}>
                    <fbt desc="">Create Group</fbt>
                  </Button>,
                ]
              : [],
          },
          children: <PermissionsGroupsView />,
        },
      },
      {
        routingPath: 'group/:id',
        component: {
          children: (
            <PermissionsGroupCard
              redirectToGroupsView={gotoGroupsPage}
              onClose={gotoGroupsPage}
            />
          ),
        },
        relatedMenuItemIndex: 1,
      },
    ];

    if (permissionPoliciesMode) {
      views.push(
        {
          routingPath: 'policies',
          menuItem: {
            label: PERMISSION_POLICIES_VIEW_NAME,
            tooltip: `${PERMISSION_POLICIES_VIEW_NAME}`,
          },
          component: {
            header: {
              title: `${PERMISSION_POLICIES_VIEW_NAME}`,
              subtitle: 'Manage policies and apply them to groups.',
              actionButtons: [
                <PopoverMenu
                  options={[
                    POLICY_TYPES.InventoryPolicy,
                    POLICY_TYPES.WorkforcePolicy,
                  ].map(type => ({
                    key: type.key,
                    value: type.key,
                    label: fbt(
                      fbt.param('policy type', type.value) + ' Policy',
                      'create policy of given type',
                    ),
                  }))}
                  skin="primary"
                  onChange={typeKey => {
                    history.push(`policy/${NEW_DIALOG_PARAM}?type=${typeKey}`);
                  }}>
                  <fbt desc="">Create Policy</fbt>
                </PopoverMenu>,
              ],
            },
            children: <PermissionsPoliciesView />,
          },
        },
        {
          routingPath: 'policy/:id',
          component: {
            children: (
              <PermissionsPolicyCard
                redirectToPoliciesView={gotoPoliciesPage}
                onClose={gotoPoliciesPage}
              />
            ),
          },
          relatedMenuItemIndex: 3,
        },
      );
    }

    return views;
  }, [gotoGroupsPage, gotoPoliciesPage, history, permissionPoliciesMode]);

  return (
    <>
      <FormContextProvider permissions={{adminRightsRequired: true}}>
        <NavigatableViews
          header={Strings.admin.users.viewHeader}
          views={VIEWS}
          routingBasePath={basePath}
        />
      </FormContextProvider>
      {addingNewUser && (
        <NewUserDialog onClose={() => setAddingNewUser(false)} />
      )}
    </>
  );
};

const UserManaementView = () => (
  <InventorySuspense isTopLevel={true}>
    <DialogShowingContextProvider>
      <UserManaementForm />
    </DialogShowingContextProvider>
  </InventorySuspense>
);

export default UserManaementView;
