// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolver

import (
	"context"
	"fmt"
	"strings"

	"github.com/facebookincubator/symphony/graph/graphql/models"
	"github.com/facebookincubator/symphony/pkg/ent"
	"github.com/facebookincubator/symphony/pkg/ent/user"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type userResolver struct{}

func (userResolver) ProfilePhoto(ctx context.Context, user *ent.User) (*ent.File, error) {
	photo, err := user.Edges.ProfilePhotoOrErr()
	if ent.IsNotLoaded(err) {
		photo, err = user.QueryProfilePhoto().Only(ctx)
	}
	return photo, ent.MaskNotFound(err)
}

func (r queryResolver) User(ctx context.Context, authID string) (*ent.User, error) {
	u, err := r.ClientFrom(ctx).User.Query().
		Where(user.AuthID(authID)).
		Only(ctx)
	return u, ent.MaskNotFound(err)
}

func (userResolver) Groups(ctx context.Context, user *ent.User) ([]*ent.UsersGroup, error) {
	return user.QueryGroups().All(ctx)
}

func (userResolver) Name(_ context.Context, user *ent.User) (string, error) {
	parts := make([]string, 0, 2)
	if user.FirstName != "" {
		parts = append(parts, user.FirstName)
	}
	if user.LastName != "" {
		parts = append(parts, user.LastName)
	}
	if len(parts) > 0 {
		return strings.Join(parts, " "), nil
	}
	return user.Email, nil
}

func (r mutationResolver) EditUser(ctx context.Context, input models.EditUserInput) (*ent.User, error) {
	client := ent.FromContext(ctx)
	u, err := client.User.Get(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("querying User %q: %w", input.ID, err)
	}

	nameFieldsAreEmpty :=
		input.FirstName != nil && *input.FirstName == "" || input.LastName != nil && *input.LastName == ""
	if nameFieldsAreEmpty {
		return nil, gqlerror.Errorf("User must have non empty first and last name values")
	}

	upd := client.User.
		UpdateOne(u).
		SetNillableFirstName(input.FirstName).
		SetNillableLastName(input.LastName).
		SetNillableStatus(input.Status).
		SetNillableRole(input.Role)

	return upd.Save(ctx)
}

func (r mutationResolver) UpdateUserGroups(ctx context.Context, input models.UpdateUserGroupsInput) (*ent.User, error) {
	return r.ClientFrom(ctx).User.
		UpdateOneID(input.ID).
		AddGroupIDs(input.AddGroupIds...).
		RemoveGroupIDs(input.RemoveGroupIds...).
		Save(ctx)
}
