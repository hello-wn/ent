// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/link"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/entc/integration/customid/ent/schema"
	"entgo.io/ent/schema/field"
)

// LinkUpdate is the builder for updating Link entities.
type LinkUpdate struct {
	config
	hooks    []Hook
	mutation *LinkMutation
}

// Where appends a list predicates to the LinkUpdate builder.
func (lu *LinkUpdate) Where(ps ...predicate.Link) *LinkUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetLinkInformation sets the "link_information" field.
func (lu *LinkUpdate) SetLinkInformation(mi map[string]schema.LinkInformation) *LinkUpdate {
	lu.mutation.SetLinkInformation(mi)
	return lu
}

// Mutation returns the LinkMutation object of the builder.
func (lu *LinkUpdate) Mutation() *LinkMutation {
	return lu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LinkUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, LinkMutation](ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LinkUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LinkUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LinkUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LinkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   link.Table,
			Columns: link.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: link.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.LinkInformation(); ok {
		_spec.SetField(link.FieldLinkInformation, field.TypeJSON, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{link.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LinkUpdateOne is the builder for updating a single Link entity.
type LinkUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LinkMutation
}

// SetLinkInformation sets the "link_information" field.
func (luo *LinkUpdateOne) SetLinkInformation(mi map[string]schema.LinkInformation) *LinkUpdateOne {
	luo.mutation.SetLinkInformation(mi)
	return luo
}

// Mutation returns the LinkMutation object of the builder.
func (luo *LinkUpdateOne) Mutation() *LinkMutation {
	return luo.mutation
}

// Where appends a list predicates to the LinkUpdate builder.
func (luo *LinkUpdateOne) Where(ps ...predicate.Link) *LinkUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LinkUpdateOne) Select(field string, fields ...string) *LinkUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Link entity.
func (luo *LinkUpdateOne) Save(ctx context.Context) (*Link, error) {
	return withHooks[*Link, LinkMutation](ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LinkUpdateOne) SaveX(ctx context.Context) *Link {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LinkUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LinkUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LinkUpdateOne) sqlSave(ctx context.Context) (_node *Link, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   link.Table,
			Columns: link.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: link.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Link.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, link.FieldID)
		for _, f := range fields {
			if !link.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != link.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.LinkInformation(); ok {
		_spec.SetField(link.FieldLinkInformation, field.TypeJSON, value)
	}
	_node = &Link{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{link.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
