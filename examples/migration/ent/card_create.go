// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/migration/ent/card"
	"entgo.io/ent/examples/migration/ent/user"
	"entgo.io/ent/schema/field"
)

// CardCreate is the builder for creating a Card entity.
type CardCreate struct {
	config
	mutation *CardMutation
	hooks    []Hook
}

// SetOwnerID sets the "owner_id" field.
func (cc *CardCreate) SetOwnerID(i int) *CardCreate {
	cc.mutation.SetOwnerID(i)
	return cc
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (cc *CardCreate) SetNillableOwnerID(i *int) *CardCreate {
	if i != nil {
		cc.SetOwnerID(*i)
	}
	return cc
}

// SetOwner sets the "owner" edge to the User entity.
func (cc *CardCreate) SetOwner(u *User) *CardCreate {
	return cc.SetOwnerID(u.ID)
}

// Mutation returns the CardMutation object of the builder.
func (cc *CardCreate) Mutation() *CardMutation {
	return cc.mutation
}

// Save creates the Card in the database.
func (cc *CardCreate) Save(ctx context.Context) (*Card, error) {
	cc.defaults()
	return withHooks[*Card, CardMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CardCreate) SaveX(ctx context.Context) *Card {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CardCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CardCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CardCreate) defaults() {
	if _, ok := cc.mutation.OwnerID(); !ok {
		v := card.DefaultOwnerID
		cc.mutation.SetOwnerID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CardCreate) check() error {
	if _, ok := cc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`ent: missing required field "Card.owner_id"`)}
	}
	if _, ok := cc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Card.owner"`)}
	}
	return nil
}

func (cc *CardCreate) sqlSave(ctx context.Context) (*Card, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CardCreate) createSpec() (*Card, *sqlgraph.CreateSpec) {
	var (
		_node = &Card{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: card.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: card.FieldID,
			},
		}
	)
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CardCreateBulk is the builder for creating many Card entities in bulk.
type CardCreateBulk struct {
	config
	builders []*CardCreate
}

// Save creates the Card entities in the database.
func (ccb *CardCreateBulk) Save(ctx context.Context) ([]*Card, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Card, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CardMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CardCreateBulk) SaveX(ctx context.Context) []*Card {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CardCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CardCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
