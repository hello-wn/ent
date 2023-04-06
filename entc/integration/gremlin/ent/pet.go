// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/entc/integration/gremlin/ent/user"
	"github.com/google/uuid"
)

// Pet is the model entity for the Pet schema.
type Pet struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Age holds the value of the "age" field.
	Age float64 `json:"age,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID uuid.UUID `json:"uuid,omitempty"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// Trained holds the value of the "trained" field.
	Trained bool `json:"trained,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PetQuery when eager-loading is set.
	Edges PetEdges `json:"edges"`
}

// PetEdges holds the relations/edges for other nodes in the graph.
type PetEdges struct {
	// Team holds the value of the team edge.
	Team *User `json:"team,omitempty"`
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TeamOrErr returns the Team value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PetEdges) TeamOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Team == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Team, nil
	}
	return nil, &NotLoadedError{edge: "team"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PetEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// FromResponse scans the gremlin response data into Pet.
func (pe *Pet) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanpe struct {
		ID       string    `json:"id,omitempty"`
		Age      float64   `json:"age,omitempty"`
		Name     string    `json:"name,omitempty"`
		UUID     uuid.UUID `json:"uuid,omitempty"`
		Nickname string    `json:"nickname,omitempty"`
		Trained  bool      `json:"trained,omitempty"`
	}
	if err := vmap.Decode(&scanpe); err != nil {
		return err
	}
	pe.ID = scanpe.ID
	pe.Age = scanpe.Age
	pe.Name = scanpe.Name
	pe.UUID = scanpe.UUID
	pe.Nickname = scanpe.Nickname
	pe.Trained = scanpe.Trained
	return nil
}

// QueryTeam queries the "team" edge of the Pet entity.
func (pe *Pet) QueryTeam() *UserQuery {
	return NewPetClient(pe.config).QueryTeam(pe)
}

// QueryOwner queries the "owner" edge of the Pet entity.
func (pe *Pet) QueryOwner() *UserQuery {
	return NewPetClient(pe.config).QueryOwner(pe)
}

// Update returns a builder for updating this Pet.
// Note that you need to call Pet.Unwrap() before calling this method if this Pet
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Pet) Update() *PetUpdateOne {
	return NewPetClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the Pet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Pet) Unwrap() *Pet {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pet is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Pet) String() string {
	var builder strings.Builder
	builder.WriteString("Pet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", pe.Age))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pe.Name)
	builder.WriteString(", ")
	builder.WriteString("uuid=")
	builder.WriteString(fmt.Sprintf("%v", pe.UUID))
	builder.WriteString(", ")
	builder.WriteString("nickname=")
	builder.WriteString(pe.Nickname)
	builder.WriteString(", ")
	builder.WriteString("trained=")
	builder.WriteString(fmt.Sprintf("%v", pe.Trained))
	builder.WriteByte(')')
	return builder.String()
}

// Pets is a parsable slice of Pet.
type Pets []*Pet

// FromResponse scans the gremlin response data into Pets.
func (pe *Pets) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanpe []struct {
		ID       string    `json:"id,omitempty"`
		Age      float64   `json:"age,omitempty"`
		Name     string    `json:"name,omitempty"`
		UUID     uuid.UUID `json:"uuid,omitempty"`
		Nickname string    `json:"nickname,omitempty"`
		Trained  bool      `json:"trained,omitempty"`
	}
	if err := vmap.Decode(&scanpe); err != nil {
		return err
	}
	for _, v := range scanpe {
		node := &Pet{ID: v.ID}
		node.Age = v.Age
		node.Name = v.Name
		node.UUID = v.UUID
		node.Nickname = v.Nickname
		node.Trained = v.Trained
		*pe = append(*pe, node)
	}
	return nil
}

func (pe Pets) config(cfg config) {
	for _i := range pe {
		pe[_i].config = cfg
	}
}
