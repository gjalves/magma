// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolver

import (
	"testing"

	"github.com/facebookincubator/symphony/graph/ent/user"
	"github.com/facebookincubator/symphony/graph/viewer"
	"github.com/facebookincubator/symphony/graph/viewer/viewertest"
	"github.com/stretchr/testify/require"
)

func TestUserOwner(t *testing.T) {
	r := newTestResolver(t)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)
	vr := r.Viewer()

	permissions, err := vr.Permissions(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, false, permissions.AdminPolicy.CanRead)
	require.Equal(t, false, permissions.CanWrite)
}

func TestUserOwnerInWriteGroup(t *testing.T) {
	r := newTestResolver(t)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)
	vr := r.Viewer()

	u, err := viewer.UserFromContext(ctx)
	require.NoError(t, err)
	_, err = r.client.UsersGroup.Create().SetName(viewer.WritePermissionGroupName).AddMembers(u).Save(ctx)
	require.NoError(t, err)

	require.NoError(t, err)
	permissions, err := vr.Permissions(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, false, permissions.AdminPolicy.CanRead)
	require.Equal(t, true, permissions.CanWrite)
}

func TestAdminViewer(t *testing.T) {
	r := newTestResolver(t)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)
	vr := r.Viewer()

	u, err := viewer.UserFromContext(ctx)
	require.NoError(t, err)
	_, err = r.client.User.UpdateOne(u).SetRole(user.RoleADMIN).Save(ctx)
	require.NoError(t, err)
	permissions, err := vr.Permissions(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, true, permissions.AdminPolicy.CanRead)
	require.Equal(t, false, permissions.CanWrite)
}

func TestOwnerViewer(t *testing.T) {
	r := newTestResolver(t)
	defer r.drv.Close()
	ctx := viewertest.NewContext(r.client)
	vr := r.Viewer()

	u, err := viewer.UserFromContext(ctx)
	require.NoError(t, err)
	_, err = r.client.User.UpdateOne(u).SetRole(user.RoleOWNER).Save(ctx)
	require.NoError(t, err)
	permissions, err := vr.Permissions(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, true, permissions.AdminPolicy.CanRead)
	require.Equal(t, true, permissions.CanWrite)
}
