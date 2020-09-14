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
	"github.com/facebookincubator/symphony/graph/ent/checklistitemdefinition"
	"github.com/facebookincubator/symphony/graph/ent/workordertype"
)

// CheckListItemDefinitionCreate is the builder for creating a CheckListItemDefinition entity.
type CheckListItemDefinitionCreate struct {
	config
	create_time     *time.Time
	update_time     *time.Time
	title           *string
	_type           *string
	index           *int
	enum_values     *string
	help_text       *string
	work_order_type map[int]struct{}
}

// SetCreateTime sets the create_time field.
func (clidc *CheckListItemDefinitionCreate) SetCreateTime(t time.Time) *CheckListItemDefinitionCreate {
	clidc.create_time = &t
	return clidc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (clidc *CheckListItemDefinitionCreate) SetNillableCreateTime(t *time.Time) *CheckListItemDefinitionCreate {
	if t != nil {
		clidc.SetCreateTime(*t)
	}
	return clidc
}

// SetUpdateTime sets the update_time field.
func (clidc *CheckListItemDefinitionCreate) SetUpdateTime(t time.Time) *CheckListItemDefinitionCreate {
	clidc.update_time = &t
	return clidc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (clidc *CheckListItemDefinitionCreate) SetNillableUpdateTime(t *time.Time) *CheckListItemDefinitionCreate {
	if t != nil {
		clidc.SetUpdateTime(*t)
	}
	return clidc
}

// SetTitle sets the title field.
func (clidc *CheckListItemDefinitionCreate) SetTitle(s string) *CheckListItemDefinitionCreate {
	clidc.title = &s
	return clidc
}

// SetType sets the type field.
func (clidc *CheckListItemDefinitionCreate) SetType(s string) *CheckListItemDefinitionCreate {
	clidc._type = &s
	return clidc
}

// SetIndex sets the index field.
func (clidc *CheckListItemDefinitionCreate) SetIndex(i int) *CheckListItemDefinitionCreate {
	clidc.index = &i
	return clidc
}

// SetNillableIndex sets the index field if the given value is not nil.
func (clidc *CheckListItemDefinitionCreate) SetNillableIndex(i *int) *CheckListItemDefinitionCreate {
	if i != nil {
		clidc.SetIndex(*i)
	}
	return clidc
}

// SetEnumValues sets the enum_values field.
func (clidc *CheckListItemDefinitionCreate) SetEnumValues(s string) *CheckListItemDefinitionCreate {
	clidc.enum_values = &s
	return clidc
}

// SetNillableEnumValues sets the enum_values field if the given value is not nil.
func (clidc *CheckListItemDefinitionCreate) SetNillableEnumValues(s *string) *CheckListItemDefinitionCreate {
	if s != nil {
		clidc.SetEnumValues(*s)
	}
	return clidc
}

// SetHelpText sets the help_text field.
func (clidc *CheckListItemDefinitionCreate) SetHelpText(s string) *CheckListItemDefinitionCreate {
	clidc.help_text = &s
	return clidc
}

// SetNillableHelpText sets the help_text field if the given value is not nil.
func (clidc *CheckListItemDefinitionCreate) SetNillableHelpText(s *string) *CheckListItemDefinitionCreate {
	if s != nil {
		clidc.SetHelpText(*s)
	}
	return clidc
}

// SetWorkOrderTypeID sets the work_order_type edge to WorkOrderType by id.
func (clidc *CheckListItemDefinitionCreate) SetWorkOrderTypeID(id int) *CheckListItemDefinitionCreate {
	if clidc.work_order_type == nil {
		clidc.work_order_type = make(map[int]struct{})
	}
	clidc.work_order_type[id] = struct{}{}
	return clidc
}

// SetNillableWorkOrderTypeID sets the work_order_type edge to WorkOrderType by id if the given value is not nil.
func (clidc *CheckListItemDefinitionCreate) SetNillableWorkOrderTypeID(id *int) *CheckListItemDefinitionCreate {
	if id != nil {
		clidc = clidc.SetWorkOrderTypeID(*id)
	}
	return clidc
}

// SetWorkOrderType sets the work_order_type edge to WorkOrderType.
func (clidc *CheckListItemDefinitionCreate) SetWorkOrderType(w *WorkOrderType) *CheckListItemDefinitionCreate {
	return clidc.SetWorkOrderTypeID(w.ID)
}

// Save creates the CheckListItemDefinition in the database.
func (clidc *CheckListItemDefinitionCreate) Save(ctx context.Context) (*CheckListItemDefinition, error) {
	if clidc.create_time == nil {
		v := checklistitemdefinition.DefaultCreateTime()
		clidc.create_time = &v
	}
	if clidc.update_time == nil {
		v := checklistitemdefinition.DefaultUpdateTime()
		clidc.update_time = &v
	}
	if clidc.title == nil {
		return nil, errors.New("ent: missing required field \"title\"")
	}
	if clidc._type == nil {
		return nil, errors.New("ent: missing required field \"type\"")
	}
	if len(clidc.work_order_type) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"work_order_type\"")
	}
	return clidc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (clidc *CheckListItemDefinitionCreate) SaveX(ctx context.Context) *CheckListItemDefinition {
	v, err := clidc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (clidc *CheckListItemDefinitionCreate) sqlSave(ctx context.Context) (*CheckListItemDefinition, error) {
	var (
		clid  = &CheckListItemDefinition{config: clidc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: checklistitemdefinition.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: checklistitemdefinition.FieldID,
			},
		}
	)
	if value := clidc.create_time; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: checklistitemdefinition.FieldCreateTime,
		})
		clid.CreateTime = *value
	}
	if value := clidc.update_time; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: checklistitemdefinition.FieldUpdateTime,
		})
		clid.UpdateTime = *value
	}
	if value := clidc.title; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: checklistitemdefinition.FieldTitle,
		})
		clid.Title = *value
	}
	if value := clidc._type; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: checklistitemdefinition.FieldType,
		})
		clid.Type = *value
	}
	if value := clidc.index; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: checklistitemdefinition.FieldIndex,
		})
		clid.Index = *value
	}
	if value := clidc.enum_values; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: checklistitemdefinition.FieldEnumValues,
		})
		clid.EnumValues = value
	}
	if value := clidc.help_text; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: checklistitemdefinition.FieldHelpText,
		})
		clid.HelpText = value
	}
	if nodes := clidc.work_order_type; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   checklistitemdefinition.WorkOrderTypeTable,
			Columns: []string{checklistitemdefinition.WorkOrderTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workordertype.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, clidc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	clid.ID = int(id)
	return clid, nil
}
