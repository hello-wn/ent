// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/ent/card"
	"entgo.io/ent/entc/integration/ent/pet"
	"entgo.io/ent/entc/integration/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `graphql:"-" json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OptionalInt holds the value of the "optional_int" field.
	OptionalInt int `json:"optional_int,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"first_name" graphql:"first_name"`
	// Last holds the value of the "last" field.
	Last string `json:"last,omitempty" graphql:"last_name"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Password holds the value of the "password" field.
	Password string `graphql:"-" json:"-"`
	// Role holds the value of the "role" field.
	Role user.Role `json:"role,omitempty"`
	// Employment holds the value of the "employment" field.
	Employment user.Employment `json:"employment,omitempty"`
	// SSOCert holds the value of the "SSOCert" field.
	SSOCert string `json:"SSOCert,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges         UserEdges `json:"edges"`
	group_blocked *int
	user_spouse   *int
	user_parent   *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Cards associated with this user. O2O edge
	Card *Card `json:"card,omitempty"`
	// Pets holds the value of the pets edge.
	Pets []*Pet `json:"pets,omitempty"`
	// Files holds the value of the files edge.
	Files []*File `json:"files,omitempty"`
	// Groups holds the value of the groups edge.
	Groups []*Group `json:"groups,omitempty"`
	// Friends holds the value of the friends edge.
	Friends []*User `json:"friends,omitempty"`
	// Followers holds the value of the followers edge.
	Followers []*User `json:"followers,omitempty"`
	// Following holds the value of the following edge.
	Following []*User `json:"following,omitempty"`
	// Team holds the value of the team edge.
	Team *Pet `json:"team,omitempty"`
	// Spouse holds the value of the spouse edge.
	Spouse *User `json:"spouse,omitempty"`
	// Children holds the value of the children edge.
	Children []*User `json:"children,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *User `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes    [11]bool
	namedPets      map[string][]*Pet
	namedFiles     map[string][]*File
	namedGroups    map[string][]*Group
	namedFriends   map[string][]*User
	namedFollowers map[string][]*User
	namedFollowing map[string][]*User
	namedChildren  map[string][]*User
}

// CardOrErr returns the Card value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) CardOrErr() (*Card, error) {
	if e.loadedTypes[0] {
		if e.Card == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: card.Label}
		}
		return e.Card, nil
	}
	return nil, &NotLoadedError{edge: "card"}
}

// PetsOrErr returns the Pets value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PetsOrErr() ([]*Pet, error) {
	if e.loadedTypes[1] {
		return e.Pets, nil
	}
	return nil, &NotLoadedError{edge: "pets"}
}

// FilesOrErr returns the Files value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FilesOrErr() ([]*File, error) {
	if e.loadedTypes[2] {
		return e.Files, nil
	}
	return nil, &NotLoadedError{edge: "files"}
}

// GroupsOrErr returns the Groups value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GroupsOrErr() ([]*Group, error) {
	if e.loadedTypes[3] {
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// FriendsOrErr returns the Friends value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FriendsOrErr() ([]*User, error) {
	if e.loadedTypes[4] {
		return e.Friends, nil
	}
	return nil, &NotLoadedError{edge: "friends"}
}

// FollowersOrErr returns the Followers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FollowersOrErr() ([]*User, error) {
	if e.loadedTypes[5] {
		return e.Followers, nil
	}
	return nil, &NotLoadedError{edge: "followers"}
}

// FollowingOrErr returns the Following value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FollowingOrErr() ([]*User, error) {
	if e.loadedTypes[6] {
		return e.Following, nil
	}
	return nil, &NotLoadedError{edge: "following"}
}

// TeamOrErr returns the Team value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) TeamOrErr() (*Pet, error) {
	if e.loadedTypes[7] {
		if e.Team == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: pet.Label}
		}
		return e.Team, nil
	}
	return nil, &NotLoadedError{edge: "team"}
}

// SpouseOrErr returns the Spouse value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) SpouseOrErr() (*User, error) {
	if e.loadedTypes[8] {
		if e.Spouse == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Spouse, nil
	}
	return nil, &NotLoadedError{edge: "spouse"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ChildrenOrErr() ([]*User, error) {
	if e.loadedTypes[9] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ParentOrErr() (*User, error) {
	if e.loadedTypes[10] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldOptionalInt, user.FieldAge:
			values[i] = new(sql.NullInt64)
		case user.FieldName, user.FieldLast, user.FieldNickname, user.FieldAddress, user.FieldPhone, user.FieldPassword, user.FieldRole, user.FieldEmployment, user.FieldSSOCert:
			values[i] = new(sql.NullString)
		case user.ForeignKeys[0]: // group_blocked
			values[i] = new(sql.NullInt64)
		case user.ForeignKeys[1]: // user_spouse
			values[i] = new(sql.NullInt64)
		case user.ForeignKeys[2]: // user_parent
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldOptionalInt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field optional_int", values[i])
			} else if value.Valid {
				u.OptionalInt = int(value.Int64)
			}
		case user.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				u.Age = int(value.Int64)
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldLast:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last", values[i])
			} else if value.Valid {
				u.Last = value.String
			}
		case user.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				u.Nickname = value.String
			}
		case user.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				u.Address = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				u.Phone = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				u.Role = user.Role(value.String)
			}
		case user.FieldEmployment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field employment", values[i])
			} else if value.Valid {
				u.Employment = user.Employment(value.String)
			}
		case user.FieldSSOCert:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field SSOCert", values[i])
			} else if value.Valid {
				u.SSOCert = value.String
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field group_blocked", value)
			} else if value.Valid {
				u.group_blocked = new(int)
				*u.group_blocked = int(value.Int64)
			}
		case user.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_spouse", value)
			} else if value.Valid {
				u.user_spouse = new(int)
				*u.user_spouse = int(value.Int64)
			}
		case user.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_parent", value)
			} else if value.Valid {
				u.user_parent = new(int)
				*u.user_parent = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryCard queries the "card" edge of the User entity.
func (u *User) QueryCard() *CardQuery {
	return NewUserClient(u.config).QueryCard(u)
}

// QueryPets queries the "pets" edge of the User entity.
func (u *User) QueryPets() *PetQuery {
	return NewUserClient(u.config).QueryPets(u)
}

// QueryFiles queries the "files" edge of the User entity.
func (u *User) QueryFiles() *FileQuery {
	return NewUserClient(u.config).QueryFiles(u)
}

// QueryGroups queries the "groups" edge of the User entity.
func (u *User) QueryGroups() *GroupQuery {
	return NewUserClient(u.config).QueryGroups(u)
}

// QueryFriends queries the "friends" edge of the User entity.
func (u *User) QueryFriends() *UserQuery {
	return NewUserClient(u.config).QueryFriends(u)
}

// QueryFollowers queries the "followers" edge of the User entity.
func (u *User) QueryFollowers() *UserQuery {
	return NewUserClient(u.config).QueryFollowers(u)
}

// QueryFollowing queries the "following" edge of the User entity.
func (u *User) QueryFollowing() *UserQuery {
	return NewUserClient(u.config).QueryFollowing(u)
}

// QueryTeam queries the "team" edge of the User entity.
func (u *User) QueryTeam() *PetQuery {
	return NewUserClient(u.config).QueryTeam(u)
}

// QuerySpouse queries the "spouse" edge of the User entity.
func (u *User) QuerySpouse() *UserQuery {
	return NewUserClient(u.config).QuerySpouse(u)
}

// QueryChildren queries the "children" edge of the User entity.
func (u *User) QueryChildren() *UserQuery {
	return NewUserClient(u.config).QueryChildren(u)
}

// QueryParent queries the "parent" edge of the User entity.
func (u *User) QueryParent() *UserQuery {
	return NewUserClient(u.config).QueryParent(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("optional_int=")
	builder.WriteString(fmt.Sprintf("%v", u.OptionalInt))
	builder.WriteString(", ")
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", u.Age))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("last=")
	builder.WriteString(u.Last)
	builder.WriteString(", ")
	builder.WriteString("nickname=")
	builder.WriteString(u.Nickname)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(u.Address)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(u.Phone)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", u.Role))
	builder.WriteString(", ")
	builder.WriteString("employment=")
	builder.WriteString(fmt.Sprintf("%v", u.Employment))
	builder.WriteString(", ")
	builder.WriteString("SSOCert=")
	builder.WriteString(u.SSOCert)
	builder.WriteByte(')')
	return builder.String()
}

// NamedPets returns the Pets named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedPets(name string) ([]*Pet, error) {
	if u.Edges.namedPets == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedPets[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedPets(name string, edges ...*Pet) {
	if u.Edges.namedPets == nil {
		u.Edges.namedPets = make(map[string][]*Pet)
	}
	if len(edges) == 0 {
		u.Edges.namedPets[name] = []*Pet{}
	} else {
		u.Edges.namedPets[name] = append(u.Edges.namedPets[name], edges...)
	}
}

// NamedFiles returns the Files named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedFiles(name string) ([]*File, error) {
	if u.Edges.namedFiles == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedFiles[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedFiles(name string, edges ...*File) {
	if u.Edges.namedFiles == nil {
		u.Edges.namedFiles = make(map[string][]*File)
	}
	if len(edges) == 0 {
		u.Edges.namedFiles[name] = []*File{}
	} else {
		u.Edges.namedFiles[name] = append(u.Edges.namedFiles[name], edges...)
	}
}

// NamedGroups returns the Groups named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedGroups(name string) ([]*Group, error) {
	if u.Edges.namedGroups == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedGroups[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedGroups(name string, edges ...*Group) {
	if u.Edges.namedGroups == nil {
		u.Edges.namedGroups = make(map[string][]*Group)
	}
	if len(edges) == 0 {
		u.Edges.namedGroups[name] = []*Group{}
	} else {
		u.Edges.namedGroups[name] = append(u.Edges.namedGroups[name], edges...)
	}
}

// NamedFriends returns the Friends named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedFriends(name string) ([]*User, error) {
	if u.Edges.namedFriends == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedFriends[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedFriends(name string, edges ...*User) {
	if u.Edges.namedFriends == nil {
		u.Edges.namedFriends = make(map[string][]*User)
	}
	if len(edges) == 0 {
		u.Edges.namedFriends[name] = []*User{}
	} else {
		u.Edges.namedFriends[name] = append(u.Edges.namedFriends[name], edges...)
	}
}

// NamedFollowers returns the Followers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedFollowers(name string) ([]*User, error) {
	if u.Edges.namedFollowers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedFollowers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedFollowers(name string, edges ...*User) {
	if u.Edges.namedFollowers == nil {
		u.Edges.namedFollowers = make(map[string][]*User)
	}
	if len(edges) == 0 {
		u.Edges.namedFollowers[name] = []*User{}
	} else {
		u.Edges.namedFollowers[name] = append(u.Edges.namedFollowers[name], edges...)
	}
}

// NamedFollowing returns the Following named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedFollowing(name string) ([]*User, error) {
	if u.Edges.namedFollowing == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedFollowing[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedFollowing(name string, edges ...*User) {
	if u.Edges.namedFollowing == nil {
		u.Edges.namedFollowing = make(map[string][]*User)
	}
	if len(edges) == 0 {
		u.Edges.namedFollowing[name] = []*User{}
	} else {
		u.Edges.namedFollowing[name] = append(u.Edges.namedFollowing[name], edges...)
	}
}

// NamedChildren returns the Children named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedChildren(name string) ([]*User, error) {
	if u.Edges.namedChildren == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedChildren[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedChildren(name string, edges ...*User) {
	if u.Edges.namedChildren == nil {
		u.Edges.namedChildren = make(map[string][]*User)
	}
	if len(edges) == 0 {
		u.Edges.namedChildren[name] = []*User{}
	} else {
		u.Edges.namedChildren[name] = append(u.Edges.namedChildren[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
