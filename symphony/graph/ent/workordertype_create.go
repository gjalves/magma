// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/symphony/graph/ent/checklistcategory"
	"github.com/facebookincubator/symphony/graph/ent/checklistitemdefinition"
	"github.com/facebookincubator/symphony/graph/ent/propertytype"
	"github.com/facebookincubator/symphony/graph/ent/workorder"
	"github.com/facebookincubator/symphony/graph/ent/workorderdefinition"
	"github.com/facebookincubator/symphony/graph/ent/workordertype"
)

// WorkOrderTypeCreate is the builder for creating a WorkOrderType entity.
type WorkOrderTypeCreate struct {
	config
	create_time            *time.Time
	update_time            *time.Time
	name                   *string
	description            *string
	work_orders            map[int]struct{}
	property_types         map[int]struct{}
	definitions            map[int]struct{}
	check_list_categories  map[int]struct{}
	check_list_definitions map[int]struct{}
}

// SetCreateTime sets the create_time field.
func (wotc *WorkOrderTypeCreate) SetCreateTime(t time.Time) *WorkOrderTypeCreate {
	wotc.create_time = &t
	return wotc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (wotc *WorkOrderTypeCreate) SetNillableCreateTime(t *time.Time) *WorkOrderTypeCreate {
	if t != nil {
		wotc.SetCreateTime(*t)
	}
	return wotc
}

// SetUpdateTime sets the update_time field.
func (wotc *WorkOrderTypeCreate) SetUpdateTime(t time.Time) *WorkOrderTypeCreate {
	wotc.update_time = &t
	return wotc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (wotc *WorkOrderTypeCreate) SetNillableUpdateTime(t *time.Time) *WorkOrderTypeCreate {
	if t != nil {
		wotc.SetUpdateTime(*t)
	}
	return wotc
}

// SetName sets the name field.
func (wotc *WorkOrderTypeCreate) SetName(s string) *WorkOrderTypeCreate {
	wotc.name = &s
	return wotc
}

// SetDescription sets the description field.
func (wotc *WorkOrderTypeCreate) SetDescription(s string) *WorkOrderTypeCreate {
	wotc.description = &s
	return wotc
}

// SetNillableDescription sets the description field if the given value is not nil.
func (wotc *WorkOrderTypeCreate) SetNillableDescription(s *string) *WorkOrderTypeCreate {
	if s != nil {
		wotc.SetDescription(*s)
	}
	return wotc
}

// AddWorkOrderIDs adds the work_orders edge to WorkOrder by ids.
func (wotc *WorkOrderTypeCreate) AddWorkOrderIDs(ids ...int) *WorkOrderTypeCreate {
	if wotc.work_orders == nil {
		wotc.work_orders = make(map[int]struct{})
	}
	for i := range ids {
		wotc.work_orders[ids[i]] = struct{}{}
	}
	return wotc
}

// AddWorkOrders adds the work_orders edges to WorkOrder.
func (wotc *WorkOrderTypeCreate) AddWorkOrders(w ...*WorkOrder) *WorkOrderTypeCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wotc.AddWorkOrderIDs(ids...)
}

// AddPropertyTypeIDs adds the property_types edge to PropertyType by ids.
func (wotc *WorkOrderTypeCreate) AddPropertyTypeIDs(ids ...int) *WorkOrderTypeCreate {
	if wotc.property_types == nil {
		wotc.property_types = make(map[int]struct{})
	}
	for i := range ids {
		wotc.property_types[ids[i]] = struct{}{}
	}
	return wotc
}

// AddPropertyTypes adds the property_types edges to PropertyType.
func (wotc *WorkOrderTypeCreate) AddPropertyTypes(p ...*PropertyType) *WorkOrderTypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return wotc.AddPropertyTypeIDs(ids...)
}

// AddDefinitionIDs adds the definitions edge to WorkOrderDefinition by ids.
func (wotc *WorkOrderTypeCreate) AddDefinitionIDs(ids ...int) *WorkOrderTypeCreate {
	if wotc.definitions == nil {
		wotc.definitions = make(map[int]struct{})
	}
	for i := range ids {
		wotc.definitions[ids[i]] = struct{}{}
	}
	return wotc
}

// AddDefinitions adds the definitions edges to WorkOrderDefinition.
func (wotc *WorkOrderTypeCreate) AddDefinitions(w ...*WorkOrderDefinition) *WorkOrderTypeCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wotc.AddDefinitionIDs(ids...)
}

// AddCheckListCategoryIDs adds the check_list_categories edge to CheckListCategory by ids.
func (wotc *WorkOrderTypeCreate) AddCheckListCategoryIDs(ids ...int) *WorkOrderTypeCreate {
	if wotc.check_list_categories == nil {
		wotc.check_list_categories = make(map[int]struct{})
	}
	for i := range ids {
		wotc.check_list_categories[ids[i]] = struct{}{}
	}
	return wotc
}

// AddCheckListCategories adds the check_list_categories edges to CheckListCategory.
func (wotc *WorkOrderTypeCreate) AddCheckListCategories(c ...*CheckListCategory) *WorkOrderTypeCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return wotc.AddCheckListCategoryIDs(ids...)
}

// AddCheckListDefinitionIDs adds the check_list_definitions edge to CheckListItemDefinition by ids.
func (wotc *WorkOrderTypeCreate) AddCheckListDefinitionIDs(ids ...int) *WorkOrderTypeCreate {
	if wotc.check_list_definitions == nil {
		wotc.check_list_definitions = make(map[int]struct{})
	}
	for i := range ids {
		wotc.check_list_definitions[ids[i]] = struct{}{}
	}
	return wotc
}

// AddCheckListDefinitions adds the check_list_definitions edges to CheckListItemDefinition.
func (wotc *WorkOrderTypeCreate) AddCheckListDefinitions(c ...*CheckListItemDefinition) *WorkOrderTypeCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return wotc.AddCheckListDefinitionIDs(ids...)
}

// Save creates the WorkOrderType in the database.
func (wotc *WorkOrderTypeCreate) Save(ctx context.Context) (*WorkOrderType, error) {
	if wotc.create_time == nil {
		v := workordertype.DefaultCreateTime()
		wotc.create_time = &v
	}
	if wotc.update_time == nil {
		v := workordertype.DefaultUpdateTime()
		wotc.update_time = &v
	}
	if wotc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	return wotc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (wotc *WorkOrderTypeCreate) SaveX(ctx context.Context) *WorkOrderType {
	v, err := wotc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wotc *WorkOrderTypeCreate) sqlSave(ctx context.Context) (*WorkOrderType, error) {
	var (
		wot   = &WorkOrderType{config: wotc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: workordertype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workordertype.FieldID,
			},
		}
	)
	if value := wotc.create_time; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: workordertype.FieldCreateTime,
		})
		wot.CreateTime = *value
	}
	if value := wotc.update_time; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: workordertype.FieldUpdateTime,
		})
		wot.UpdateTime = *value
	}
	if value := wotc.name; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: workordertype.FieldName,
		})
		wot.Name = *value
	}
	if value := wotc.description; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: workordertype.FieldDescription,
		})
		wot.Description = *value
	}
	if nodes := wotc.work_orders; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   workordertype.WorkOrdersTable,
			Columns: []string{workordertype.WorkOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workorder.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wotc.property_types; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workordertype.PropertyTypesTable,
			Columns: []string{workordertype.PropertyTypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: propertytype.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wotc.definitions; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   workordertype.DefinitionsTable,
			Columns: []string{workordertype.DefinitionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workorderdefinition.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wotc.check_list_categories; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workordertype.CheckListCategoriesTable,
			Columns: []string{workordertype.CheckListCategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: checklistcategory.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wotc.check_list_definitions; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workordertype.CheckListDefinitionsTable,
			Columns: []string{workordertype.CheckListDefinitionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: checklistitemdefinition.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, wotc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	wot.ID = int(id)
	return wot, nil
}
