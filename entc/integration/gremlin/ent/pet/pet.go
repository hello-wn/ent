// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package pet

const (
	// Label holds the string label denoting the pet type in the database.
	Label = "pet"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldTrained holds the string denoting the trained field in the database.
	FieldTrained = "trained"
	// EdgeTeam holds the string denoting the team edge name in mutations.
	EdgeTeam = "team"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// TeamInverseLabel holds the string label denoting the team inverse edge type in the database.
	TeamInverseLabel = "user_team"
	// OwnerInverseLabel holds the string label denoting the owner inverse edge type in the database.
	OwnerInverseLabel = "user_pets"
)

var (
	// DefaultAge holds the default value on creation for the "age" field.
	DefaultAge float64
	// DefaultTrained holds the default value on creation for the "trained" field.
	DefaultTrained bool
)

// comment from another template.
