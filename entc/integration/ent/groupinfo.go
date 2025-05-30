// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/ent/groupinfo"
)

// GroupInfo is the model entity for the GroupInfo schema.
type GroupInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc string `json:"desc,omitempty"`
	// MaxUsers holds the value of the "max_users" field.
	MaxUsers int `json:"max_users,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupInfoQuery when eager-loading is set.
	Edges        GroupInfoEdges `json:"edges"`
	selectValues sql.SelectValues
}

// GroupInfoEdges holds the relations/edges for other nodes in the graph.
type GroupInfoEdges struct {
	// Groups holds the value of the groups edge.
	Groups []*Group `json:"groups,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	namedGroups map[string][]*Group
}

// GroupsOrErr returns the Groups value or an error if the edge
// was not loaded in eager-loading.
func (e GroupInfoEdges) GroupsOrErr() ([]*Group, error) {
	if e.loadedTypes[0] {
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroupInfo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case groupinfo.FieldID, groupinfo.FieldMaxUsers:
			values[i] = new(sql.NullInt64)
		case groupinfo.FieldDesc:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupInfo fields.
func (_m *GroupInfo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case groupinfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			_m.ID = int(value.Int64)
		case groupinfo.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				_m.Desc = value.String
			}
		case groupinfo.FieldMaxUsers:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field max_users", values[i])
			} else if value.Valid {
				_m.MaxUsers = int(value.Int64)
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the GroupInfo.
// This includes values selected through modifiers, order, etc.
func (_m *GroupInfo) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// QueryGroups queries the "groups" edge of the GroupInfo entity.
func (_m *GroupInfo) QueryGroups() *GroupQuery {
	return NewGroupInfoClient(_m.config).QueryGroups(_m)
}

// Update returns a builder for updating this GroupInfo.
// Note that you need to call GroupInfo.Unwrap() before calling this method if this GroupInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *GroupInfo) Update() *GroupInfoUpdateOne {
	return NewGroupInfoClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the GroupInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *GroupInfo) Unwrap() *GroupInfo {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: GroupInfo is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *GroupInfo) String() string {
	var builder strings.Builder
	builder.WriteString("GroupInfo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("desc=")
	builder.WriteString(_m.Desc)
	builder.WriteString(", ")
	builder.WriteString("max_users=")
	builder.WriteString(fmt.Sprintf("%v", _m.MaxUsers))
	builder.WriteByte(')')
	return builder.String()
}

// NamedGroups returns the Groups named value or an error if the edge was not
// loaded in eager-loading with this name.
func (_m *GroupInfo) NamedGroups(name string) ([]*Group, error) {
	if _m.Edges.namedGroups == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := _m.Edges.namedGroups[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (_m *GroupInfo) appendNamedGroups(name string, edges ...*Group) {
	if _m.Edges.namedGroups == nil {
		_m.Edges.namedGroups = make(map[string][]*Group)
	}
	if len(edges) == 0 {
		_m.Edges.namedGroups[name] = []*Group{}
	} else {
		_m.Edges.namedGroups[name] = append(_m.Edges.namedGroups[name], edges...)
	}
}

// GroupInfos is a parsable slice of GroupInfo.
type GroupInfos []*GroupInfo
