// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/entc/integration/gremlin/ent/card"
	"entgo.io/ent/entc/integration/gremlin/ent/pet"
	"entgo.io/ent/entc/integration/gremlin/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `graphql:"-" json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
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
	// FilesCount holds the value of the "files_count" field.
	FilesCount int `json:"files_count,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
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
	loadedTypes [11]bool
}

// CardOrErr returns the Card value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) CardOrErr() (*Card, error) {
	if e.Card != nil {
		return e.Card, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: card.Label}
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
	if e.Team != nil {
		return e.Team, nil
	} else if e.loadedTypes[7] {
		return nil, &NotFoundError{label: pet.Label}
	}
	return nil, &NotLoadedError{edge: "team"}
}

// SpouseOrErr returns the Spouse value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) SpouseOrErr() (*User, error) {
	if e.Spouse != nil {
		return e.Spouse, nil
	} else if e.loadedTypes[8] {
		return nil, &NotFoundError{label: user.Label}
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
	if e.Parent != nil {
		return e.Parent, nil
	} else if e.loadedTypes[10] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// FromResponse scans the gremlin response data into User.
func (_m *User) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scan_m struct {
		ID          string          `json:"id,omitempty"`
		OptionalInt int             `json:"optional_int,omitempty"`
		Age         int             `json:"age,omitempty"`
		Name        string          `json:"name,omitempty"`
		Last        string          `json:"last,omitempty"`
		Nickname    string          `json:"nickname,omitempty"`
		Address     string          `json:"address,omitempty"`
		Phone       string          `json:"phone,omitempty"`
		Password    string          `json:"password,omitempty"`
		Role        user.Role       `json:"role,omitempty"`
		Employment  user.Employment `json:"employment,omitempty"`
		SSOCert     string          `json:"sso_cert,omitempty"`
		FilesCount  int             `json:"files_count,omitempty"`
	}
	if err := vmap.Decode(&scan_m); err != nil {
		return err
	}
	_m.ID = scan_m.ID
	_m.OptionalInt = scan_m.OptionalInt
	_m.Age = scan_m.Age
	_m.Name = scan_m.Name
	_m.Last = scan_m.Last
	_m.Nickname = scan_m.Nickname
	_m.Address = scan_m.Address
	_m.Phone = scan_m.Phone
	_m.Password = scan_m.Password
	_m.Role = scan_m.Role
	_m.Employment = scan_m.Employment
	_m.SSOCert = scan_m.SSOCert
	_m.FilesCount = scan_m.FilesCount
	return nil
}

// QueryCard queries the "card" edge of the User entity.
func (_m *User) QueryCard() *CardQuery {
	return NewUserClient(_m.config).QueryCard(_m)
}

// QueryPets queries the "pets" edge of the User entity.
func (_m *User) QueryPets() *PetQuery {
	return NewUserClient(_m.config).QueryPets(_m)
}

// QueryFiles queries the "files" edge of the User entity.
func (_m *User) QueryFiles() *FileQuery {
	return NewUserClient(_m.config).QueryFiles(_m)
}

// QueryGroups queries the "groups" edge of the User entity.
func (_m *User) QueryGroups() *GroupQuery {
	return NewUserClient(_m.config).QueryGroups(_m)
}

// QueryFriends queries the "friends" edge of the User entity.
func (_m *User) QueryFriends() *UserQuery {
	return NewUserClient(_m.config).QueryFriends(_m)
}

// QueryFollowers queries the "followers" edge of the User entity.
func (_m *User) QueryFollowers() *UserQuery {
	return NewUserClient(_m.config).QueryFollowers(_m)
}

// QueryFollowing queries the "following" edge of the User entity.
func (_m *User) QueryFollowing() *UserQuery {
	return NewUserClient(_m.config).QueryFollowing(_m)
}

// QueryTeam queries the "team" edge of the User entity.
func (_m *User) QueryTeam() *PetQuery {
	return NewUserClient(_m.config).QueryTeam(_m)
}

// QuerySpouse queries the "spouse" edge of the User entity.
func (_m *User) QuerySpouse() *UserQuery {
	return NewUserClient(_m.config).QuerySpouse(_m)
}

// QueryChildren queries the "children" edge of the User entity.
func (_m *User) QueryChildren() *UserQuery {
	return NewUserClient(_m.config).QueryChildren(_m)
}

// QueryParent queries the "parent" edge of the User entity.
func (_m *User) QueryParent() *UserQuery {
	return NewUserClient(_m.config).QueryParent(_m)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *User) Update() *UserUpdateOne {
	return NewUserClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *User) Unwrap() *User {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("optional_int=")
	builder.WriteString(fmt.Sprintf("%v", _m.OptionalInt))
	builder.WriteString(", ")
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", _m.Age))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(_m.Name)
	builder.WriteString(", ")
	builder.WriteString("last=")
	builder.WriteString(_m.Last)
	builder.WriteString(", ")
	builder.WriteString("nickname=")
	builder.WriteString(_m.Nickname)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(_m.Address)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(_m.Phone)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", _m.Role))
	builder.WriteString(", ")
	builder.WriteString("employment=")
	builder.WriteString(fmt.Sprintf("%v", _m.Employment))
	builder.WriteString(", ")
	builder.WriteString("SSOCert=")
	builder.WriteString(_m.SSOCert)
	builder.WriteString(", ")
	builder.WriteString("files_count=")
	builder.WriteString(fmt.Sprintf("%v", _m.FilesCount))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

// FromResponse scans the gremlin response data into Users.
func (_m *Users) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scan_m []struct {
		ID          string          `json:"id,omitempty"`
		OptionalInt int             `json:"optional_int,omitempty"`
		Age         int             `json:"age,omitempty"`
		Name        string          `json:"name,omitempty"`
		Last        string          `json:"last,omitempty"`
		Nickname    string          `json:"nickname,omitempty"`
		Address     string          `json:"address,omitempty"`
		Phone       string          `json:"phone,omitempty"`
		Password    string          `json:"password,omitempty"`
		Role        user.Role       `json:"role,omitempty"`
		Employment  user.Employment `json:"employment,omitempty"`
		SSOCert     string          `json:"sso_cert,omitempty"`
		FilesCount  int             `json:"files_count,omitempty"`
	}
	if err := vmap.Decode(&scan_m); err != nil {
		return err
	}
	for _, v := range scan_m {
		node := &User{ID: v.ID}
		node.OptionalInt = v.OptionalInt
		node.Age = v.Age
		node.Name = v.Name
		node.Last = v.Last
		node.Nickname = v.Nickname
		node.Address = v.Address
		node.Phone = v.Phone
		node.Password = v.Password
		node.Role = v.Role
		node.Employment = v.Employment
		node.SSOCert = v.SSOCert
		node.FilesCount = v.FilesCount
		*_m = append(*_m, node)
	}
	return nil
}
