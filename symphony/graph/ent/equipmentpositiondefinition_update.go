// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/symphony/graph/ent/equipmentposition"
	"github.com/facebookincubator/symphony/graph/ent/equipmentpositiondefinition"
	"github.com/facebookincubator/symphony/graph/ent/equipmenttype"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
)

// EquipmentPositionDefinitionUpdate is the builder for updating EquipmentPositionDefinition entities.
type EquipmentPositionDefinitionUpdate struct {
	config

	update_time           *time.Time
	name                  *string
	index                 *int
	addindex              *int
	clearindex            bool
	visibility_label      *string
	clearvisibility_label bool
	positions             map[int]struct{}
	equipment_type        map[int]struct{}
	removedPositions      map[int]struct{}
	clearedEquipmentType  bool
	predicates            []predicate.EquipmentPositionDefinition
}

// Where adds a new predicate for the builder.
func (epdu *EquipmentPositionDefinitionUpdate) Where(ps ...predicate.EquipmentPositionDefinition) *EquipmentPositionDefinitionUpdate {
	epdu.predicates = append(epdu.predicates, ps...)
	return epdu
}

// SetName sets the name field.
func (epdu *EquipmentPositionDefinitionUpdate) SetName(s string) *EquipmentPositionDefinitionUpdate {
	epdu.name = &s
	return epdu
}

// SetIndex sets the index field.
func (epdu *EquipmentPositionDefinitionUpdate) SetIndex(i int) *EquipmentPositionDefinitionUpdate {
	epdu.index = &i
	epdu.addindex = nil
	return epdu
}

// SetNillableIndex sets the index field if the given value is not nil.
func (epdu *EquipmentPositionDefinitionUpdate) SetNillableIndex(i *int) *EquipmentPositionDefinitionUpdate {
	if i != nil {
		epdu.SetIndex(*i)
	}
	return epdu
}

// AddIndex adds i to index.
func (epdu *EquipmentPositionDefinitionUpdate) AddIndex(i int) *EquipmentPositionDefinitionUpdate {
	if epdu.addindex == nil {
		epdu.addindex = &i
	} else {
		*epdu.addindex += i
	}
	return epdu
}

// ClearIndex clears the value of index.
func (epdu *EquipmentPositionDefinitionUpdate) ClearIndex() *EquipmentPositionDefinitionUpdate {
	epdu.index = nil
	epdu.clearindex = true
	return epdu
}

// SetVisibilityLabel sets the visibility_label field.
func (epdu *EquipmentPositionDefinitionUpdate) SetVisibilityLabel(s string) *EquipmentPositionDefinitionUpdate {
	epdu.visibility_label = &s
	return epdu
}

// SetNillableVisibilityLabel sets the visibility_label field if the given value is not nil.
func (epdu *EquipmentPositionDefinitionUpdate) SetNillableVisibilityLabel(s *string) *EquipmentPositionDefinitionUpdate {
	if s != nil {
		epdu.SetVisibilityLabel(*s)
	}
	return epdu
}

// ClearVisibilityLabel clears the value of visibility_label.
func (epdu *EquipmentPositionDefinitionUpdate) ClearVisibilityLabel() *EquipmentPositionDefinitionUpdate {
	epdu.visibility_label = nil
	epdu.clearvisibility_label = true
	return epdu
}

// AddPositionIDs adds the positions edge to EquipmentPosition by ids.
func (epdu *EquipmentPositionDefinitionUpdate) AddPositionIDs(ids ...int) *EquipmentPositionDefinitionUpdate {
	if epdu.positions == nil {
		epdu.positions = make(map[int]struct{})
	}
	for i := range ids {
		epdu.positions[ids[i]] = struct{}{}
	}
	return epdu
}

// AddPositions adds the positions edges to EquipmentPosition.
func (epdu *EquipmentPositionDefinitionUpdate) AddPositions(e ...*EquipmentPosition) *EquipmentPositionDefinitionUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return epdu.AddPositionIDs(ids...)
}

// SetEquipmentTypeID sets the equipment_type edge to EquipmentType by id.
func (epdu *EquipmentPositionDefinitionUpdate) SetEquipmentTypeID(id int) *EquipmentPositionDefinitionUpdate {
	if epdu.equipment_type == nil {
		epdu.equipment_type = make(map[int]struct{})
	}
	epdu.equipment_type[id] = struct{}{}
	return epdu
}

// SetNillableEquipmentTypeID sets the equipment_type edge to EquipmentType by id if the given value is not nil.
func (epdu *EquipmentPositionDefinitionUpdate) SetNillableEquipmentTypeID(id *int) *EquipmentPositionDefinitionUpdate {
	if id != nil {
		epdu = epdu.SetEquipmentTypeID(*id)
	}
	return epdu
}

// SetEquipmentType sets the equipment_type edge to EquipmentType.
func (epdu *EquipmentPositionDefinitionUpdate) SetEquipmentType(e *EquipmentType) *EquipmentPositionDefinitionUpdate {
	return epdu.SetEquipmentTypeID(e.ID)
}

// RemovePositionIDs removes the positions edge to EquipmentPosition by ids.
func (epdu *EquipmentPositionDefinitionUpdate) RemovePositionIDs(ids ...int) *EquipmentPositionDefinitionUpdate {
	if epdu.removedPositions == nil {
		epdu.removedPositions = make(map[int]struct{})
	}
	for i := range ids {
		epdu.removedPositions[ids[i]] = struct{}{}
	}
	return epdu
}

// RemovePositions removes positions edges to EquipmentPosition.
func (epdu *EquipmentPositionDefinitionUpdate) RemovePositions(e ...*EquipmentPosition) *EquipmentPositionDefinitionUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return epdu.RemovePositionIDs(ids...)
}

// ClearEquipmentType clears the equipment_type edge to EquipmentType.
func (epdu *EquipmentPositionDefinitionUpdate) ClearEquipmentType() *EquipmentPositionDefinitionUpdate {
	epdu.clearedEquipmentType = true
	return epdu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (epdu *EquipmentPositionDefinitionUpdate) Save(ctx context.Context) (int, error) {
	if epdu.update_time == nil {
		v := equipmentpositiondefinition.UpdateDefaultUpdateTime()
		epdu.update_time = &v
	}
	if len(epdu.equipment_type) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"equipment_type\"")
	}
	return epdu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (epdu *EquipmentPositionDefinitionUpdate) SaveX(ctx context.Context) int {
	affected, err := epdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (epdu *EquipmentPositionDefinitionUpdate) Exec(ctx context.Context) error {
	_, err := epdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (epdu *EquipmentPositionDefinitionUpdate) ExecX(ctx context.Context) {
	if err := epdu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (epdu *EquipmentPositionDefinitionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   equipmentpositiondefinition.Table,
			Columns: equipmentpositiondefinition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: equipmentpositiondefinition.FieldID,
			},
		},
	}
	if ps := epdu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := epdu.update_time; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: equipmentpositiondefinition.FieldUpdateTime,
		})
	}
	if value := epdu.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: equipmentpositiondefinition.FieldName,
		})
	}
	if value := epdu.index; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: equipmentpositiondefinition.FieldIndex,
		})
	}
	if value := epdu.addindex; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: equipmentpositiondefinition.FieldIndex,
		})
	}
	if epdu.clearindex {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: equipmentpositiondefinition.FieldIndex,
		})
	}
	if value := epdu.visibility_label; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: equipmentpositiondefinition.FieldVisibilityLabel,
		})
	}
	if epdu.clearvisibility_label {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: equipmentpositiondefinition.FieldVisibilityLabel,
		})
	}
	if nodes := epdu.removedPositions; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   equipmentpositiondefinition.PositionsTable,
			Columns: []string{equipmentpositiondefinition.PositionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipmentposition.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := epdu.positions; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   equipmentpositiondefinition.PositionsTable,
			Columns: []string{equipmentpositiondefinition.PositionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipmentposition.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if epdu.clearedEquipmentType {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   equipmentpositiondefinition.EquipmentTypeTable,
			Columns: []string{equipmentpositiondefinition.EquipmentTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipmenttype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := epdu.equipment_type; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   equipmentpositiondefinition.EquipmentTypeTable,
			Columns: []string{equipmentpositiondefinition.EquipmentTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipmenttype.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, epdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{equipmentpositiondefinition.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// EquipmentPositionDefinitionUpdateOne is the builder for updating a single EquipmentPositionDefinition entity.
type EquipmentPositionDefinitionUpdateOne struct {
	config
	id int

	update_time           *time.Time
	name                  *string
	index                 *int
	addindex              *int
	clearindex            bool
	visibility_label      *string
	clearvisibility_label bool
	positions             map[int]struct{}
	equipment_type        map[int]struct{}
	removedPositions      map[int]struct{}
	clearedEquipmentType  bool
}

// SetName sets the name field.
func (epduo *EquipmentPositionDefinitionUpdateOne) SetName(s string) *EquipmentPositionDefinitionUpdateOne {
	epduo.name = &s
	return epduo
}

// SetIndex sets the index field.
func (epduo *EquipmentPositionDefinitionUpdateOne) SetIndex(i int) *EquipmentPositionDefinitionUpdateOne {
	epduo.index = &i
	epduo.addindex = nil
	return epduo
}

// SetNillableIndex sets the index field if the given value is not nil.
func (epduo *EquipmentPositionDefinitionUpdateOne) SetNillableIndex(i *int) *EquipmentPositionDefinitionUpdateOne {
	if i != nil {
		epduo.SetIndex(*i)
	}
	return epduo
}

// AddIndex adds i to index.
func (epduo *EquipmentPositionDefinitionUpdateOne) AddIndex(i int) *EquipmentPositionDefinitionUpdateOne {
	if epduo.addindex == nil {
		epduo.addindex = &i
	} else {
		*epduo.addindex += i
	}
	return epduo
}

// ClearIndex clears the value of index.
func (epduo *EquipmentPositionDefinitionUpdateOne) ClearIndex() *EquipmentPositionDefinitionUpdateOne {
	epduo.index = nil
	epduo.clearindex = true
	return epduo
}

// SetVisibilityLabel sets the visibility_label field.
func (epduo *EquipmentPositionDefinitionUpdateOne) SetVisibilityLabel(s string) *EquipmentPositionDefinitionUpdateOne {
	epduo.visibility_label = &s
	return epduo
}

// SetNillableVisibilityLabel sets the visibility_label field if the given value is not nil.
func (epduo *EquipmentPositionDefinitionUpdateOne) SetNillableVisibilityLabel(s *string) *EquipmentPositionDefinitionUpdateOne {
	if s != nil {
		epduo.SetVisibilityLabel(*s)
	}
	return epduo
}

// ClearVisibilityLabel clears the value of visibility_label.
func (epduo *EquipmentPositionDefinitionUpdateOne) ClearVisibilityLabel() *EquipmentPositionDefinitionUpdateOne {
	epduo.visibility_label = nil
	epduo.clearvisibility_label = true
	return epduo
}

// AddPositionIDs adds the positions edge to EquipmentPosition by ids.
func (epduo *EquipmentPositionDefinitionUpdateOne) AddPositionIDs(ids ...int) *EquipmentPositionDefinitionUpdateOne {
	if epduo.positions == nil {
		epduo.positions = make(map[int]struct{})
	}
	for i := range ids {
		epduo.positions[ids[i]] = struct{}{}
	}
	return epduo
}

// AddPositions adds the positions edges to EquipmentPosition.
func (epduo *EquipmentPositionDefinitionUpdateOne) AddPositions(e ...*EquipmentPosition) *EquipmentPositionDefinitionUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return epduo.AddPositionIDs(ids...)
}

// SetEquipmentTypeID sets the equipment_type edge to EquipmentType by id.
func (epduo *EquipmentPositionDefinitionUpdateOne) SetEquipmentTypeID(id int) *EquipmentPositionDefinitionUpdateOne {
	if epduo.equipment_type == nil {
		epduo.equipment_type = make(map[int]struct{})
	}
	epduo.equipment_type[id] = struct{}{}
	return epduo
}

// SetNillableEquipmentTypeID sets the equipment_type edge to EquipmentType by id if the given value is not nil.
func (epduo *EquipmentPositionDefinitionUpdateOne) SetNillableEquipmentTypeID(id *int) *EquipmentPositionDefinitionUpdateOne {
	if id != nil {
		epduo = epduo.SetEquipmentTypeID(*id)
	}
	return epduo
}

// SetEquipmentType sets the equipment_type edge to EquipmentType.
func (epduo *EquipmentPositionDefinitionUpdateOne) SetEquipmentType(e *EquipmentType) *EquipmentPositionDefinitionUpdateOne {
	return epduo.SetEquipmentTypeID(e.ID)
}

// RemovePositionIDs removes the positions edge to EquipmentPosition by ids.
func (epduo *EquipmentPositionDefinitionUpdateOne) RemovePositionIDs(ids ...int) *EquipmentPositionDefinitionUpdateOne {
	if epduo.removedPositions == nil {
		epduo.removedPositions = make(map[int]struct{})
	}
	for i := range ids {
		epduo.removedPositions[ids[i]] = struct{}{}
	}
	return epduo
}

// RemovePositions removes positions edges to EquipmentPosition.
func (epduo *EquipmentPositionDefinitionUpdateOne) RemovePositions(e ...*EquipmentPosition) *EquipmentPositionDefinitionUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return epduo.RemovePositionIDs(ids...)
}

// ClearEquipmentType clears the equipment_type edge to EquipmentType.
func (epduo *EquipmentPositionDefinitionUpdateOne) ClearEquipmentType() *EquipmentPositionDefinitionUpdateOne {
	epduo.clearedEquipmentType = true
	return epduo
}

// Save executes the query and returns the updated entity.
func (epduo *EquipmentPositionDefinitionUpdateOne) Save(ctx context.Context) (*EquipmentPositionDefinition, error) {
	if epduo.update_time == nil {
		v := equipmentpositiondefinition.UpdateDefaultUpdateTime()
		epduo.update_time = &v
	}
	if len(epduo.equipment_type) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"equipment_type\"")
	}
	return epduo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (epduo *EquipmentPositionDefinitionUpdateOne) SaveX(ctx context.Context) *EquipmentPositionDefinition {
	epd, err := epduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return epd
}

// Exec executes the query on the entity.
func (epduo *EquipmentPositionDefinitionUpdateOne) Exec(ctx context.Context) error {
	_, err := epduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (epduo *EquipmentPositionDefinitionUpdateOne) ExecX(ctx context.Context) {
	if err := epduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (epduo *EquipmentPositionDefinitionUpdateOne) sqlSave(ctx context.Context) (epd *EquipmentPositionDefinition, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   equipmentpositiondefinition.Table,
			Columns: equipmentpositiondefinition.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  epduo.id,
				Type:   field.TypeInt,
				Column: equipmentpositiondefinition.FieldID,
			},
		},
	}
	if value := epduo.update_time; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: equipmentpositiondefinition.FieldUpdateTime,
		})
	}
	if value := epduo.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: equipmentpositiondefinition.FieldName,
		})
	}
	if value := epduo.index; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: equipmentpositiondefinition.FieldIndex,
		})
	}
	if value := epduo.addindex; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: equipmentpositiondefinition.FieldIndex,
		})
	}
	if epduo.clearindex {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: equipmentpositiondefinition.FieldIndex,
		})
	}
	if value := epduo.visibility_label; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: equipmentpositiondefinition.FieldVisibilityLabel,
		})
	}
	if epduo.clearvisibility_label {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: equipmentpositiondefinition.FieldVisibilityLabel,
		})
	}
	if nodes := epduo.removedPositions; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   equipmentpositiondefinition.PositionsTable,
			Columns: []string{equipmentpositiondefinition.PositionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipmentposition.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := epduo.positions; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   equipmentpositiondefinition.PositionsTable,
			Columns: []string{equipmentpositiondefinition.PositionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipmentposition.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if epduo.clearedEquipmentType {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   equipmentpositiondefinition.EquipmentTypeTable,
			Columns: []string{equipmentpositiondefinition.EquipmentTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipmenttype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := epduo.equipment_type; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   equipmentpositiondefinition.EquipmentTypeTable,
			Columns: []string{equipmentpositiondefinition.EquipmentTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipmenttype.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	epd = &EquipmentPositionDefinition{config: epduo.config}
	_spec.Assign = epd.assignValues
	_spec.ScanValues = epd.scanValues()
	if err = sqlgraph.UpdateNode(ctx, epduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{equipmentpositiondefinition.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return epd, nil
}
