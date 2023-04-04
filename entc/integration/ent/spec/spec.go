// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package spec

const (
	// Label holds the string label denoting the spec type in the database.
	Label = "spec"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeCard holds the string denoting the card edge name in mutations.
	EdgeCard = "card"
	// Table holds the table name of the spec in the database.
	Table = "specs"
	// CardTable is the table that holds the card relation/edge. The primary key declared below.
	CardTable = "spec_card"
	// CardInverseTable is the table name for the Card entity.
	// It exists in this package in order to avoid circular dependency with the "card" package.
	CardInverseTable = "cards"
)

// Columns holds all SQL columns for spec fields.
var Columns = []string{
	FieldID,
}

var (
	// CardPrimaryKey and CardColumn2 are the table columns denoting the
	// primary key for the card relation (M2M).
	CardPrimaryKey = []string{"spec_id", "card_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// comment from another template.
