// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/entc/integration/gremlin/ent/filetype"
	"entgo.io/ent/entc/integration/gremlin/ent/user"
)

// File is the model entity for the File schema.
type File struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// SetID holds the value of the "set_id" field.
	SetID int `json:"set_id,omitempty"`
	// Size holds the value of the "size" field.
	Size int `json:"size,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// User holds the value of the "user" field.
	User *string `json:"user,omitempty"`
	// Group holds the value of the "group" field.
	Group string `json:"group,omitempty"`
	// Op holds the value of the "op" field.
	Op bool `json:"op,omitempty"`
	// FieldID holds the value of the "field_id" field.
	FieldID int `json:"field_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileQuery when eager-loading is set.
	Edges FileEdges `json:"file_edges"`
}

// FileEdges holds the relations/edges for other nodes in the graph.
type FileEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Type holds the value of the type edge.
	Type *FileType `json:"type,omitempty"`
	// Field holds the value of the field edge.
	Field []*FieldType `json:"field,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FileEdges) OwnerOrErr() (*User, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// TypeOrErr returns the Type value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FileEdges) TypeOrErr() (*FileType, error) {
	if e.Type != nil {
		return e.Type, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: filetype.Label}
	}
	return nil, &NotLoadedError{edge: "type"}
}

// FieldOrErr returns the Field value or an error if the edge
// was not loaded in eager-loading.
func (e FileEdges) FieldOrErr() ([]*FieldType, error) {
	if e.loadedTypes[2] {
		return e.Field, nil
	}
	return nil, &NotLoadedError{edge: "field"}
}

// FromResponse scans the gremlin response data into File.
func (f *File) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanf struct {
		ID      string  `json:"id,omitempty"`
		SetID   int     `json:"set_id,omitempty"`
		Size    int     `json:"fsize,omitempty"`
		Name    string  `json:"name,omitempty"`
		User    *string `json:"user,omitempty"`
		Group   string  `json:"group,omitempty"`
		Op      bool    `json:"op,omitempty"`
		FieldID int     `json:"field_id,omitempty"`
	}
	if err := vmap.Decode(&scanf); err != nil {
		return err
	}
	f.ID = scanf.ID
	f.SetID = scanf.SetID
	f.Size = scanf.Size
	f.Name = scanf.Name
	f.User = scanf.User
	f.Group = scanf.Group
	f.Op = scanf.Op
	f.FieldID = scanf.FieldID
	return nil
}

// QueryOwner queries the "owner" edge of the File entity.
func (f *File) QueryOwner() *UserQuery {
	return NewFileClient(f.config).QueryOwner(f)
}

// QueryType queries the "type" edge of the File entity.
func (f *File) QueryType() *FileTypeQuery {
	return NewFileClient(f.config).QueryType(f)
}

// QueryField queries the "field" edge of the File entity.
func (f *File) QueryField() *FieldTypeQuery {
	return NewFileClient(f.config).QueryField(f)
}

// Update returns a builder for updating this File.
// Note that you need to call File.Unwrap() before calling this method if this File
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *File) Update() *FileUpdateOne {
	return NewFileClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the File entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *File) Unwrap() *File {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: File is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *File) String() string {
	var builder strings.Builder
	builder.WriteString("File(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("set_id=")
	builder.WriteString(fmt.Sprintf("%v", f.SetID))
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", f.Size))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(f.Name)
	builder.WriteString(", ")
	if v := f.User; v != nil {
		builder.WriteString("user=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("group=")
	builder.WriteString(f.Group)
	builder.WriteString(", ")
	builder.WriteString("op=")
	builder.WriteString(fmt.Sprintf("%v", f.Op))
	builder.WriteString(", ")
	builder.WriteString("field_id=")
	builder.WriteString(fmt.Sprintf("%v", f.FieldID))
	builder.WriteByte(')')
	return builder.String()
}

// Files is a parsable slice of File.
type Files []*File

// FromResponse scans the gremlin response data into Files.
func (f *Files) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanf []struct {
		ID      string  `json:"id,omitempty"`
		SetID   int     `json:"set_id,omitempty"`
		Size    int     `json:"fsize,omitempty"`
		Name    string  `json:"name,omitempty"`
		User    *string `json:"user,omitempty"`
		Group   string  `json:"group,omitempty"`
		Op      bool    `json:"op,omitempty"`
		FieldID int     `json:"field_id,omitempty"`
	}
	if err := vmap.Decode(&scanf); err != nil {
		return err
	}
	for _, v := range scanf {
		node := &File{ID: v.ID}
		node.SetID = v.SetID
		node.Size = v.Size
		node.Name = v.Name
		node.User = v.User
		node.Group = v.Group
		node.Op = v.Op
		node.FieldID = v.FieldID
		*f = append(*f, node)
	}
	return nil
}
