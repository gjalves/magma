// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolverutil

import (
	"fmt"

	"github.com/facebookincubator/symphony/graph/graphql/models"
	"github.com/facebookincubator/symphony/pkg/ent"
	"github.com/facebookincubator/symphony/pkg/ent/location"
	"github.com/facebookincubator/symphony/pkg/ent/predicate"
	"github.com/facebookincubator/symphony/pkg/ent/user"
	"github.com/facebookincubator/symphony/pkg/ent/workorder"
	"github.com/facebookincubator/symphony/pkg/ent/workordertype"

	"github.com/pkg/errors"
)

func handleWorkOrderFilter(q *ent.WorkOrderQuery, filter *models.WorkOrderFilterInput) (*ent.WorkOrderQuery, error) {
	if filter.FilterType == models.WorkOrderFilterTypeWorkOrderName {
		return nameFilter(q, filter)
	}
	if filter.FilterType == models.WorkOrderFilterTypeWorkOrderStatus {
		return statusFilter(q, filter)
	}
	if filter.FilterType == models.WorkOrderFilterTypeWorkOrderOwnedBy {
		return ownedByFilter(q, filter)
	}
	if filter.FilterType == models.WorkOrderFilterTypeWorkOrderType {
		return typeFilter(q, filter)
	}
	if filter.FilterType == models.WorkOrderFilterTypeWorkOrderAssignedTo {
		return assignedToFilter(q, filter)
	}
	if filter.FilterType == models.WorkOrderFilterTypeWorkOrderCreationDate {
		return creationDateFilter(q, filter)
	}
	if filter.FilterType == models.WorkOrderFilterTypeWorkOrderCloseDate {
		return closeDateFilter(q, filter)
	}
	if filter.FilterType == models.WorkOrderFilterTypeWorkOrderLocationInst {
		return locationInstFilter(q, filter)
	}
	if filter.FilterType == models.WorkOrderFilterTypeWorkOrderPriority {
		return priorityFilter(q, filter)
	}
	return nil, errors.Errorf("filter type is not supported: %s", filter.FilterType)
}

func nameFilter(q *ent.WorkOrderQuery, filter *models.WorkOrderFilterInput) (*ent.WorkOrderQuery, error) {
	if filter.Operator == models.FilterOperatorContains && filter.StringValue != nil {
		return q.Where(workorder.NameContainsFold(*filter.StringValue)), nil
	}
	return nil, errors.Errorf("operation is not supported: %s", filter.Operator)
}

func statusFilter(q *ent.WorkOrderQuery, filter *models.WorkOrderFilterInput) (*ent.WorkOrderQuery, error) {
	if filter.Operator != models.FilterOperatorIsOneOf {
		return nil, errors.Errorf("operation %q is not supported", filter.Operator)
	}
	statuses := make([]workorder.Status, 0, len(filter.StringSet))
	for _, str := range filter.StringSet {
		status := workorder.Status(str)
		if err := workorder.StatusValidator(status); err != nil {
			return nil, fmt.Errorf("%s is not a valid work order status", str)
		}
		statuses = append(statuses, status)
	}
	return q.Where(workorder.StatusIn(statuses...)), nil
}

func ownedByFilter(q *ent.WorkOrderQuery, filter *models.WorkOrderFilterInput) (*ent.WorkOrderQuery, error) {
	if filter.Operator == models.FilterOperatorIsOneOf {
		return q.Where(workorder.HasOwnerWith(user.IDIn(filter.IDSet...))), nil
	}
	return nil, errors.Errorf("operation is not supported: %s", filter.Operator)
}

func typeFilter(q *ent.WorkOrderQuery, filter *models.WorkOrderFilterInput) (*ent.WorkOrderQuery, error) {
	if filter.Operator == models.FilterOperatorIsOneOf {
		return q.Where(workorder.HasTypeWith(workordertype.IDIn(filter.IDSet...))), nil
	}
	return nil, errors.Errorf("operation is not supported: %s", filter.Operator)
}

func assignedToFilter(q *ent.WorkOrderQuery, filter *models.WorkOrderFilterInput) (*ent.WorkOrderQuery, error) {
	if filter.Operator == models.FilterOperatorIsOneOf {
		return q.Where(workorder.HasAssigneeWith(user.IDIn(filter.IDSet...))), nil
	}
	return nil, errors.Errorf("operation is not supported: %s", filter.Operator)
}

func creationDateFilter(q *ent.WorkOrderQuery, filter *models.WorkOrderFilterInput) (*ent.WorkOrderQuery, error) {
	switch filter.Operator {
	case models.FilterOperatorDateLessThan:
		return q.Where(workorder.CreationDateLT(*filter.TimeValue)), nil
	case models.FilterOperatorDateLessOrEqualThan:
		return q.Where(workorder.CreationDateLTE(*filter.TimeValue)), nil
	case models.FilterOperatorDateGreaterThan:
		return q.Where(workorder.CreationDateGT(*filter.TimeValue)), nil
	case models.FilterOperatorDateGreaterOrEqualThan:
		return q.Where(workorder.CreationDateGTE(*filter.TimeValue)), nil
	}
	return nil, errors.Errorf("operation is not supported: %s", filter.Operator)
}

func closeDateFilter(q *ent.WorkOrderQuery, filter *models.WorkOrderFilterInput) (*ent.WorkOrderQuery, error) {
	switch filter.Operator {
	case models.FilterOperatorDateLessThan:
		return q.Where(workorder.CloseDateLT(*filter.TimeValue)), nil
	case models.FilterOperatorDateLessOrEqualThan:
		return q.Where(workorder.CloseDateLTE(*filter.TimeValue)), nil
	case models.FilterOperatorDateGreaterThan:
		return q.Where(workorder.CloseDateGT(*filter.TimeValue)), nil
	case models.FilterOperatorDateGreaterOrEqualThan:
		return q.Where(workorder.CloseDateGTE(*filter.TimeValue)), nil
	}
	return nil, errors.Errorf("operation is not supported: %s", filter.Operator)
}

func locationInstFilter(q *ent.WorkOrderQuery, filter *models.WorkOrderFilterInput) (*ent.WorkOrderQuery, error) {
	if filter.Operator == models.FilterOperatorIsOneOf {
		return q.Where(workorder.HasLocationWith(location.IDIn(filter.IDSet...))), nil
	}
	return nil, errors.Errorf("operation is not supported: %s", filter.Operator)
}

func priorityFilter(q *ent.WorkOrderQuery, filter *models.WorkOrderFilterInput) (*ent.WorkOrderQuery, error) {
	if filter.Operator != models.FilterOperatorIsOneOf {
		return nil, fmt.Errorf("operation %q is not supported", filter.Operator)
	}
	priorities := make([]workorder.Priority, 0, len(filter.StringSet))
	for _, str := range filter.StringSet {
		priority := workorder.Priority(str)
		if err := workorder.PriorityValidator(priority); err != nil {
			return nil, fmt.Errorf("%s is not a valid work order priority", str)
		}
		priorities = append(priorities, priority)
	}
	return q.Where(workorder.PriorityIn(priorities...)), nil
}

func handleWOLocationFilter(q *ent.WorkOrderQuery, filter *models.WorkOrderFilterInput) (*ent.WorkOrderQuery, error) {
	switch filter.FilterType {
	case models.WorkOrderFilterTypeLocationInst:
		return woLocationFilter(q, filter)
	case models.WorkOrderFilterTypeLocationInstExternalID:
		return woLocationExternalIDFilter(q, filter)
	}
	return nil, errors.Errorf("filter type is not supported: %s", filter.FilterType)
}

func woLocationExternalIDFilter(q *ent.WorkOrderQuery, filter *models.WorkOrderFilterInput) (*ent.WorkOrderQuery, error) {
	if filter.Operator == models.FilterOperatorContains {
		return q.Where(workorder.HasLocationWith(location.ExternalIDContainsFold(*filter.StringValue))), nil
	}
	return nil, errors.Errorf("operation %s is not supported", filter.Operator)
}

func woLocationFilter(q *ent.WorkOrderQuery, filter *models.WorkOrderFilterInput) (*ent.WorkOrderQuery, error) {
	if filter.Operator == models.FilterOperatorIsOneOf {
		if filter.MaxDepth == nil {
			return nil, errors.New("max depth not supplied to location filter")
		}
		var ps []predicate.WorkOrder
		for _, lid := range filter.IDSet {
			ps = append(ps, workorder.HasLocationWith(BuildLocationAncestorFilter(lid, 1, *filter.MaxDepth)))
		}
		return q.Where(workorder.Or(ps...)), nil
	}
	return nil, errors.Errorf("operation is not supported: %s", filter.Operator)
}
