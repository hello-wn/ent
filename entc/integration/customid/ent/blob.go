// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/customid/ent/blob"
	"github.com/google/uuid"
)

// Blob is the model entity for the Blob schema.
type Blob struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID uuid.UUID `json:"uuid,omitempty"`
	// Count holds the value of the "count" field.
	Count int `json:"count,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BlobQuery when eager-loading is set.
	Edges       BlobEdges `json:"edges"`
	blob_parent *uuid.UUID
}

// BlobEdges holds the relations/edges for other nodes in the graph.
type BlobEdges struct {
	// Parent holds the value of the parent edge.
	Parent *Blob `json:"parent,omitempty"`
	// Links holds the value of the links edge.
	Links []*Blob `json:"links,omitempty"`
	// BlobLinks holds the value of the blob_links edge.
	BlobLinks []*BlobLink `json:"blob_links,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BlobEdges) ParentOrErr() (*Blob, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: blob.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// LinksOrErr returns the Links value or an error if the edge
// was not loaded in eager-loading.
func (e BlobEdges) LinksOrErr() ([]*Blob, error) {
	if e.loadedTypes[1] {
		return e.Links, nil
	}
	return nil, &NotLoadedError{edge: "links"}
}

// BlobLinksOrErr returns the BlobLinks value or an error if the edge
// was not loaded in eager-loading.
func (e BlobEdges) BlobLinksOrErr() ([]*BlobLink, error) {
	if e.loadedTypes[2] {
		return e.BlobLinks, nil
	}
	return nil, &NotLoadedError{edge: "blob_links"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Blob) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case blob.FieldCount:
			values[i] = new(sql.NullInt64)
		case blob.FieldID, blob.FieldUUID:
			values[i] = new(uuid.UUID)
		case blob.ForeignKeys[0]: // blob_parent
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Blob", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Blob fields.
func (b *Blob) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case blob.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				b.ID = *value
			}
		case blob.FieldUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value != nil {
				b.UUID = *value
			}
		case blob.FieldCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field count", values[i])
			} else if value.Valid {
				b.Count = int(value.Int64)
			}
		case blob.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field blob_parent", values[i])
			} else if value.Valid {
				b.blob_parent = new(uuid.UUID)
				*b.blob_parent = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the Blob entity.
func (b *Blob) QueryParent() *BlobQuery {
	return (&BlobClient{config: b.config}).QueryParent(b)
}

// QueryLinks queries the "links" edge of the Blob entity.
func (b *Blob) QueryLinks() *BlobQuery {
	return (&BlobClient{config: b.config}).QueryLinks(b)
}

// QueryBlobLinks queries the "blob_links" edge of the Blob entity.
func (b *Blob) QueryBlobLinks() *BlobLinkQuery {
	return (&BlobClient{config: b.config}).QueryBlobLinks(b)
}

// Update returns a builder for updating this Blob.
// Note that you need to call Blob.Unwrap() before calling this method if this Blob
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Blob) Update() *BlobUpdateOne {
	return (&BlobClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the Blob entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Blob) Unwrap() *Blob {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Blob is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Blob) String() string {
	var builder strings.Builder
	builder.WriteString("Blob(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("uuid=")
	builder.WriteString(fmt.Sprintf("%v", b.UUID))
	builder.WriteString(", ")
	builder.WriteString("count=")
	builder.WriteString(fmt.Sprintf("%v", b.Count))
	builder.WriteByte(')')
	return builder.String()
}

// Blobs is a parsable slice of Blob.
type Blobs []*Blob

func (b Blobs) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
