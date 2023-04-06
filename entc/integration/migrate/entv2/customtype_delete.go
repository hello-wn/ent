// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv2

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/migrate/entv2/customtype"
	"entgo.io/ent/entc/integration/migrate/entv2/predicate"
	"entgo.io/ent/schema/field"
)

// CustomTypeDelete is the builder for deleting a CustomType entity.
type CustomTypeDelete struct {
	config
	hooks    []Hook
	mutation *CustomTypeMutation
}

// Where appends a list predicates to the CustomTypeDelete builder.
func (ctd *CustomTypeDelete) Where(ps ...predicate.CustomType) *CustomTypeDelete {
	ctd.mutation.Where(ps...)
	return ctd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ctd *CustomTypeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, CustomTypeMutation](ctx, ctd.sqlExec, ctd.mutation, ctd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ctd *CustomTypeDelete) ExecX(ctx context.Context) int {
	n, err := ctd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ctd *CustomTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: customtype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customtype.FieldID,
			},
		},
	}
	if ps := ctd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ctd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ctd.mutation.done = true
	return affected, err
}

// CustomTypeDeleteOne is the builder for deleting a single CustomType entity.
type CustomTypeDeleteOne struct {
	ctd *CustomTypeDelete
}

// Where appends a list predicates to the CustomTypeDelete builder.
func (ctdo *CustomTypeDeleteOne) Where(ps ...predicate.CustomType) *CustomTypeDeleteOne {
	ctdo.ctd.mutation.Where(ps...)
	return ctdo
}

// Exec executes the deletion query.
func (ctdo *CustomTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := ctdo.ctd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{customtype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ctdo *CustomTypeDeleteOne) ExecX(ctx context.Context) {
	if err := ctdo.Exec(ctx); err != nil {
		panic(err)
	}
}
