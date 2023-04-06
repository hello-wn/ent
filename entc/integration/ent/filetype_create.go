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
	"entgo.io/ent/entc/integration/ent/file"
	"entgo.io/ent/entc/integration/ent/filetype"
	"entgo.io/ent/schema/field"
)

// FileTypeCreate is the builder for creating a FileType entity.
type FileTypeCreate struct {
	config
	mutation *FileTypeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (ftc *FileTypeCreate) SetName(s string) *FileTypeCreate {
	ftc.mutation.SetName(s)
	return ftc
}

// SetType sets the "type" field.
func (ftc *FileTypeCreate) SetType(f filetype.Type) *FileTypeCreate {
	ftc.mutation.SetType(f)
	return ftc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ftc *FileTypeCreate) SetNillableType(f *filetype.Type) *FileTypeCreate {
	if f != nil {
		ftc.SetType(*f)
	}
	return ftc
}

// SetState sets the "state" field.
func (ftc *FileTypeCreate) SetState(f filetype.State) *FileTypeCreate {
	ftc.mutation.SetState(f)
	return ftc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (ftc *FileTypeCreate) SetNillableState(f *filetype.State) *FileTypeCreate {
	if f != nil {
		ftc.SetState(*f)
	}
	return ftc
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (ftc *FileTypeCreate) AddFileIDs(ids ...int) *FileTypeCreate {
	ftc.mutation.AddFileIDs(ids...)
	return ftc
}

// AddFiles adds the "files" edges to the File entity.
func (ftc *FileTypeCreate) AddFiles(f ...*File) *FileTypeCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftc.AddFileIDs(ids...)
}

// Mutation returns the FileTypeMutation object of the builder.
func (ftc *FileTypeCreate) Mutation() *FileTypeMutation {
	return ftc.mutation
}

// Save creates the FileType in the database.
func (ftc *FileTypeCreate) Save(ctx context.Context) (*FileType, error) {
	ftc.defaults()
	return withHooks[*FileType, FileTypeMutation](ctx, ftc.sqlSave, ftc.mutation, ftc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ftc *FileTypeCreate) SaveX(ctx context.Context) *FileType {
	v, err := ftc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ftc *FileTypeCreate) Exec(ctx context.Context) error {
	_, err := ftc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftc *FileTypeCreate) ExecX(ctx context.Context) {
	if err := ftc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ftc *FileTypeCreate) defaults() {
	if _, ok := ftc.mutation.GetType(); !ok {
		v := filetype.DefaultType
		ftc.mutation.SetType(v)
	}
	if _, ok := ftc.mutation.State(); !ok {
		v := filetype.DefaultState
		ftc.mutation.SetState(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftc *FileTypeCreate) check() error {
	if _, ok := ftc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "FileType.name"`)}
	}
	if _, ok := ftc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "FileType.type"`)}
	}
	if v, ok := ftc.mutation.GetType(); ok {
		if err := filetype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "FileType.type": %w`, err)}
		}
	}
	if _, ok := ftc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "FileType.state"`)}
	}
	if v, ok := ftc.mutation.State(); ok {
		if err := filetype.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "FileType.state": %w`, err)}
		}
	}
	return nil
}

func (ftc *FileTypeCreate) sqlSave(ctx context.Context) (*FileType, error) {
	if err := ftc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ftc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ftc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ftc.mutation.id = &_node.ID
	ftc.mutation.done = true
	return _node, nil
}

func (ftc *FileTypeCreate) createSpec() (*FileType, *sqlgraph.CreateSpec) {
	var (
		_node = &FileType{config: ftc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: filetype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: filetype.FieldID,
			},
		}
	)
	_spec.OnConflict = ftc.conflict
	if value, ok := ftc.mutation.Name(); ok {
		_spec.SetField(filetype.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ftc.mutation.GetType(); ok {
		_spec.SetField(filetype.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := ftc.mutation.State(); ok {
		_spec.SetField(filetype.FieldState, field.TypeEnum, value)
		_node.State = value
	}
	if nodes := ftc.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filetype.FilesTable,
			Columns: []string{filetype.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FileType.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FileTypeUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (ftc *FileTypeCreate) OnConflict(opts ...sql.ConflictOption) *FileTypeUpsertOne {
	ftc.conflict = opts
	return &FileTypeUpsertOne{
		create: ftc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FileType.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ftc *FileTypeCreate) OnConflictColumns(columns ...string) *FileTypeUpsertOne {
	ftc.conflict = append(ftc.conflict, sql.ConflictColumns(columns...))
	return &FileTypeUpsertOne{
		create: ftc,
	}
}

type (
	// FileTypeUpsertOne is the builder for "upsert"-ing
	//  one FileType node.
	FileTypeUpsertOne struct {
		create *FileTypeCreate
	}

	// FileTypeUpsert is the "OnConflict" setter.
	FileTypeUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *FileTypeUpsert) SetName(v string) *FileTypeUpsert {
	u.Set(filetype.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FileTypeUpsert) UpdateName() *FileTypeUpsert {
	u.SetExcluded(filetype.FieldName)
	return u
}

// SetType sets the "type" field.
func (u *FileTypeUpsert) SetType(v filetype.Type) *FileTypeUpsert {
	u.Set(filetype.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *FileTypeUpsert) UpdateType() *FileTypeUpsert {
	u.SetExcluded(filetype.FieldType)
	return u
}

// SetState sets the "state" field.
func (u *FileTypeUpsert) SetState(v filetype.State) *FileTypeUpsert {
	u.Set(filetype.FieldState, v)
	return u
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *FileTypeUpsert) UpdateState() *FileTypeUpsert {
	u.SetExcluded(filetype.FieldState)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.FileType.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *FileTypeUpsertOne) UpdateNewValues() *FileTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FileType.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FileTypeUpsertOne) Ignore() *FileTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FileTypeUpsertOne) DoNothing() *FileTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FileTypeCreate.OnConflict
// documentation for more info.
func (u *FileTypeUpsertOne) Update(set func(*FileTypeUpsert)) *FileTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FileTypeUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *FileTypeUpsertOne) SetName(v string) *FileTypeUpsertOne {
	return u.Update(func(s *FileTypeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FileTypeUpsertOne) UpdateName() *FileTypeUpsertOne {
	return u.Update(func(s *FileTypeUpsert) {
		s.UpdateName()
	})
}

// SetType sets the "type" field.
func (u *FileTypeUpsertOne) SetType(v filetype.Type) *FileTypeUpsertOne {
	return u.Update(func(s *FileTypeUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *FileTypeUpsertOne) UpdateType() *FileTypeUpsertOne {
	return u.Update(func(s *FileTypeUpsert) {
		s.UpdateType()
	})
}

// SetState sets the "state" field.
func (u *FileTypeUpsertOne) SetState(v filetype.State) *FileTypeUpsertOne {
	return u.Update(func(s *FileTypeUpsert) {
		s.SetState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *FileTypeUpsertOne) UpdateState() *FileTypeUpsertOne {
	return u.Update(func(s *FileTypeUpsert) {
		s.UpdateState()
	})
}

// Exec executes the query.
func (u *FileTypeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FileTypeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FileTypeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FileTypeUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FileTypeUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FileTypeCreateBulk is the builder for creating many FileType entities in bulk.
type FileTypeCreateBulk struct {
	config
	builders []*FileTypeCreate
	conflict []sql.ConflictOption
}

// Save creates the FileType entities in the database.
func (ftcb *FileTypeCreateBulk) Save(ctx context.Context) ([]*FileType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ftcb.builders))
	nodes := make([]*FileType, len(ftcb.builders))
	mutators := make([]Mutator, len(ftcb.builders))
	for i := range ftcb.builders {
		func(i int, root context.Context) {
			builder := ftcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileTypeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ftcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ftcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ftcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ftcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ftcb *FileTypeCreateBulk) SaveX(ctx context.Context) []*FileType {
	v, err := ftcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ftcb *FileTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := ftcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftcb *FileTypeCreateBulk) ExecX(ctx context.Context) {
	if err := ftcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FileType.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FileTypeUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (ftcb *FileTypeCreateBulk) OnConflict(opts ...sql.ConflictOption) *FileTypeUpsertBulk {
	ftcb.conflict = opts
	return &FileTypeUpsertBulk{
		create: ftcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FileType.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ftcb *FileTypeCreateBulk) OnConflictColumns(columns ...string) *FileTypeUpsertBulk {
	ftcb.conflict = append(ftcb.conflict, sql.ConflictColumns(columns...))
	return &FileTypeUpsertBulk{
		create: ftcb,
	}
}

// FileTypeUpsertBulk is the builder for "upsert"-ing
// a bulk of FileType nodes.
type FileTypeUpsertBulk struct {
	create *FileTypeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.FileType.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *FileTypeUpsertBulk) UpdateNewValues() *FileTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FileType.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FileTypeUpsertBulk) Ignore() *FileTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FileTypeUpsertBulk) DoNothing() *FileTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FileTypeCreateBulk.OnConflict
// documentation for more info.
func (u *FileTypeUpsertBulk) Update(set func(*FileTypeUpsert)) *FileTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FileTypeUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *FileTypeUpsertBulk) SetName(v string) *FileTypeUpsertBulk {
	return u.Update(func(s *FileTypeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FileTypeUpsertBulk) UpdateName() *FileTypeUpsertBulk {
	return u.Update(func(s *FileTypeUpsert) {
		s.UpdateName()
	})
}

// SetType sets the "type" field.
func (u *FileTypeUpsertBulk) SetType(v filetype.Type) *FileTypeUpsertBulk {
	return u.Update(func(s *FileTypeUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *FileTypeUpsertBulk) UpdateType() *FileTypeUpsertBulk {
	return u.Update(func(s *FileTypeUpsert) {
		s.UpdateType()
	})
}

// SetState sets the "state" field.
func (u *FileTypeUpsertBulk) SetState(v filetype.State) *FileTypeUpsertBulk {
	return u.Update(func(s *FileTypeUpsert) {
		s.SetState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *FileTypeUpsertBulk) UpdateState() *FileTypeUpsertBulk {
	return u.Update(func(s *FileTypeUpsert) {
		s.UpdateState()
	})
}

// Exec executes the query.
func (u *FileTypeUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FileTypeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FileTypeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FileTypeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
