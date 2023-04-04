// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entgo.io/ent/entc/integration/cascadelete/ent/post"
	"entgo.io/ent/entc/integration/cascadelete/ent/schema"
	"entgo.io/ent/entc/integration/cascadelete/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescText is the schema descriptor for text field.
	postDescText := postFields[0].Descriptor()
	// post.DefaultText holds the default value on creation for the text field.
	post.DefaultText = postDescText.Default.(string)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
}
