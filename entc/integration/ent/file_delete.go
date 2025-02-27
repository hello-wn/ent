// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/file"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/schema/field"
)

// FileDelete is the builder for deleting a File entity.
type FileDelete struct {
	config
	hooks    []Hook
	mutation *FileMutation
}

// Where appends a list predicates to the FileDelete builder.
func (fd *FileDelete) Where(ps ...predicate.File) *FileDelete {
	fd.mutation.Where(ps...)
	return fd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fd *FileDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, FileMutation](ctx, fd.sqlExec, fd.mutation, fd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fd *FileDelete) ExecX(ctx context.Context) int {
	n, err := fd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fd *FileDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: file.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: file.FieldID,
			},
		},
	}
	if ps := fd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fd.mutation.done = true
	return affected, err
}

// FileDeleteOne is the builder for deleting a single File entity.
type FileDeleteOne struct {
	fd *FileDelete
}

// Where appends a list predicates to the FileDelete builder.
func (fdo *FileDeleteOne) Where(ps ...predicate.File) *FileDeleteOne {
	fdo.fd.mutation.Where(ps...)
	return fdo
}

// Exec executes the deletion query.
func (fdo *FileDeleteOne) Exec(ctx context.Context) error {
	n, err := fdo.fd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{file.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fdo *FileDeleteOne) ExecX(ctx context.Context) {
	if err := fdo.Exec(ctx); err != nil {
		panic(err)
	}
}
