// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/edgeschema/ent/group"
	"entgo.io/ent/entc/integration/edgeschema/ent/grouptag"
	"entgo.io/ent/entc/integration/edgeschema/ent/tag"
)

// GroupTag is the model entity for the GroupTag schema.
type GroupTag struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TagID holds the value of the "tag_id" field.
	TagID int `json:"tag_id,omitempty"`
	// GroupID holds the value of the "group_id" field.
	GroupID int `json:"group_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupTagQuery when eager-loading is set.
	Edges GroupTagEdges `json:"edges"`
}

// GroupTagEdges holds the relations/edges for other nodes in the graph.
type GroupTagEdges struct {
	// Tag holds the value of the tag edge.
	Tag *Tag `json:"tag,omitempty"`
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TagOrErr returns the Tag value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupTagEdges) TagOrErr() (*Tag, error) {
	if e.loadedTypes[0] {
		if e.Tag == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tag.Label}
		}
		return e.Tag, nil
	}
	return nil, &NotLoadedError{edge: "tag"}
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupTagEdges) GroupOrErr() (*Group, error) {
	if e.loadedTypes[1] {
		if e.Group == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroupTag) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case grouptag.FieldID, grouptag.FieldTagID, grouptag.FieldGroupID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GroupTag", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupTag fields.
func (gt *GroupTag) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case grouptag.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gt.ID = int(value.Int64)
		case grouptag.FieldTagID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tag_id", values[i])
			} else if value.Valid {
				gt.TagID = int(value.Int64)
			}
		case grouptag.FieldGroupID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value.Valid {
				gt.GroupID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryTag queries the "tag" edge of the GroupTag entity.
func (gt *GroupTag) QueryTag() *TagQuery {
	return NewGroupTagClient(gt.config).QueryTag(gt)
}

// QueryGroup queries the "group" edge of the GroupTag entity.
func (gt *GroupTag) QueryGroup() *GroupQuery {
	return NewGroupTagClient(gt.config).QueryGroup(gt)
}

// Update returns a builder for updating this GroupTag.
// Note that you need to call GroupTag.Unwrap() before calling this method if this GroupTag
// was returned from a transaction, and the transaction was committed or rolled back.
func (gt *GroupTag) Update() *GroupTagUpdateOne {
	return NewGroupTagClient(gt.config).UpdateOne(gt)
}

// Unwrap unwraps the GroupTag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gt *GroupTag) Unwrap() *GroupTag {
	_tx, ok := gt.config.driver.(*txDriver)
	if !ok {
		panic("ent: GroupTag is not a transactional entity")
	}
	gt.config.driver = _tx.drv
	return gt
}

// String implements the fmt.Stringer.
func (gt *GroupTag) String() string {
	var builder strings.Builder
	builder.WriteString("GroupTag(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gt.ID))
	builder.WriteString("tag_id=")
	builder.WriteString(fmt.Sprintf("%v", gt.TagID))
	builder.WriteString(", ")
	builder.WriteString("group_id=")
	builder.WriteString(fmt.Sprintf("%v", gt.GroupID))
	builder.WriteByte(')')
	return builder.String()
}

// GroupTags is a parsable slice of GroupTag.
type GroupTags []*GroupTag

func (gt GroupTags) config(cfg config) {
	for _i := range gt {
		gt[_i].config = cfg
	}
}
