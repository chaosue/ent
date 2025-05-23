// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/privacy/ent/predicate"
	"entgo.io/ent/entc/integration/privacy/ent/task"
	"entgo.io/ent/entc/integration/privacy/ent/team"
	"entgo.io/ent/entc/integration/privacy/ent/user"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (_u *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetTitle sets the "title" field.
func (_u *TaskUpdate) SetTitle(v string) *TaskUpdate {
	_u.mutation.SetTitle(v)
	return _u
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (_u *TaskUpdate) SetNillableTitle(v *string) *TaskUpdate {
	if v != nil {
		_u.SetTitle(*v)
	}
	return _u
}

// SetDescription sets the "description" field.
func (_u *TaskUpdate) SetDescription(v string) *TaskUpdate {
	_u.mutation.SetDescription(v)
	return _u
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (_u *TaskUpdate) SetNillableDescription(v *string) *TaskUpdate {
	if v != nil {
		_u.SetDescription(*v)
	}
	return _u
}

// ClearDescription clears the value of the "description" field.
func (_u *TaskUpdate) ClearDescription() *TaskUpdate {
	_u.mutation.ClearDescription()
	return _u
}

// SetStatus sets the "status" field.
func (_u *TaskUpdate) SetStatus(v task.Status) *TaskUpdate {
	_u.mutation.SetStatus(v)
	return _u
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (_u *TaskUpdate) SetNillableStatus(v *task.Status) *TaskUpdate {
	if v != nil {
		_u.SetStatus(*v)
	}
	return _u
}

// SetUUID sets the "uuid" field.
func (_u *TaskUpdate) SetUUID(v uuid.UUID) *TaskUpdate {
	_u.mutation.SetUUID(v)
	return _u
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (_u *TaskUpdate) SetNillableUUID(v *uuid.UUID) *TaskUpdate {
	if v != nil {
		_u.SetUUID(*v)
	}
	return _u
}

// ClearUUID clears the value of the "uuid" field.
func (_u *TaskUpdate) ClearUUID() *TaskUpdate {
	_u.mutation.ClearUUID()
	return _u
}

// AddTeamIDs adds the "teams" edge to the Team entity by IDs.
func (_u *TaskUpdate) AddTeamIDs(ids ...int) *TaskUpdate {
	_u.mutation.AddTeamIDs(ids...)
	return _u
}

// AddTeams adds the "teams" edges to the Team entity.
func (_u *TaskUpdate) AddTeams(v ...*Team) *TaskUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.AddTeamIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (_u *TaskUpdate) SetOwnerID(id int) *TaskUpdate {
	_u.mutation.SetOwnerID(id)
	return _u
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (_u *TaskUpdate) SetNillableOwnerID(id *int) *TaskUpdate {
	if id != nil {
		_u = _u.SetOwnerID(*id)
	}
	return _u
}

// SetOwner sets the "owner" edge to the User entity.
func (_u *TaskUpdate) SetOwner(v *User) *TaskUpdate {
	return _u.SetOwnerID(v.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (_u *TaskUpdate) Mutation() *TaskMutation {
	return _u.mutation
}

// ClearTeams clears all "teams" edges to the Team entity.
func (_u *TaskUpdate) ClearTeams() *TaskUpdate {
	_u.mutation.ClearTeams()
	return _u
}

// RemoveTeamIDs removes the "teams" edge to Team entities by IDs.
func (_u *TaskUpdate) RemoveTeamIDs(ids ...int) *TaskUpdate {
	_u.mutation.RemoveTeamIDs(ids...)
	return _u
}

// RemoveTeams removes "teams" edges to Team entities.
func (_u *TaskUpdate) RemoveTeams(v ...*Team) *TaskUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.RemoveTeamIDs(ids...)
}

// ClearOwner clears the "owner" edge to the User entity.
func (_u *TaskUpdate) ClearOwner() *TaskUpdate {
	_u.mutation.ClearOwner()
	return _u
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *TaskUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *TaskUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *TaskUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *TaskUpdate) check() error {
	if v, ok := _u.mutation.Title(); ok {
		if err := task.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Task.title": %w`, err)}
		}
	}
	if v, ok := _u.mutation.Status(); ok {
		if err := task.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Task.status": %w`, err)}
		}
	}
	return nil
}

func (_u *TaskUpdate) sqlSave(ctx context.Context) (_node int, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Title(); ok {
		_spec.SetField(task.FieldTitle, field.TypeString, value)
	}
	if value, ok := _u.mutation.Description(); ok {
		_spec.SetField(task.FieldDescription, field.TypeString, value)
	}
	if _u.mutation.DescriptionCleared() {
		_spec.ClearField(task.FieldDescription, field.TypeString)
	}
	if value, ok := _u.mutation.Status(); ok {
		_spec.SetField(task.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := _u.mutation.UUID(); ok {
		_spec.SetField(task.FieldUUID, field.TypeUUID, value)
	}
	if _u.mutation.UUIDCleared() {
		_spec.ClearField(task.FieldUUID, field.TypeUUID)
	}
	if _u.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   task.TeamsTable,
			Columns: task.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedTeamsIDs(); len(nodes) > 0 && !_u.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   task.TeamsTable,
			Columns: task.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.TeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   task.TeamsTable,
			Columns: task.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.OwnerTable,
			Columns: []string{task.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.OwnerTable,
			Columns: []string{task.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _node, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return _node, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskMutation
}

// SetTitle sets the "title" field.
func (_u *TaskUpdateOne) SetTitle(v string) *TaskUpdateOne {
	_u.mutation.SetTitle(v)
	return _u
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (_u *TaskUpdateOne) SetNillableTitle(v *string) *TaskUpdateOne {
	if v != nil {
		_u.SetTitle(*v)
	}
	return _u
}

// SetDescription sets the "description" field.
func (_u *TaskUpdateOne) SetDescription(v string) *TaskUpdateOne {
	_u.mutation.SetDescription(v)
	return _u
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (_u *TaskUpdateOne) SetNillableDescription(v *string) *TaskUpdateOne {
	if v != nil {
		_u.SetDescription(*v)
	}
	return _u
}

// ClearDescription clears the value of the "description" field.
func (_u *TaskUpdateOne) ClearDescription() *TaskUpdateOne {
	_u.mutation.ClearDescription()
	return _u
}

// SetStatus sets the "status" field.
func (_u *TaskUpdateOne) SetStatus(v task.Status) *TaskUpdateOne {
	_u.mutation.SetStatus(v)
	return _u
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (_u *TaskUpdateOne) SetNillableStatus(v *task.Status) *TaskUpdateOne {
	if v != nil {
		_u.SetStatus(*v)
	}
	return _u
}

// SetUUID sets the "uuid" field.
func (_u *TaskUpdateOne) SetUUID(v uuid.UUID) *TaskUpdateOne {
	_u.mutation.SetUUID(v)
	return _u
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (_u *TaskUpdateOne) SetNillableUUID(v *uuid.UUID) *TaskUpdateOne {
	if v != nil {
		_u.SetUUID(*v)
	}
	return _u
}

// ClearUUID clears the value of the "uuid" field.
func (_u *TaskUpdateOne) ClearUUID() *TaskUpdateOne {
	_u.mutation.ClearUUID()
	return _u
}

// AddTeamIDs adds the "teams" edge to the Team entity by IDs.
func (_u *TaskUpdateOne) AddTeamIDs(ids ...int) *TaskUpdateOne {
	_u.mutation.AddTeamIDs(ids...)
	return _u
}

// AddTeams adds the "teams" edges to the Team entity.
func (_u *TaskUpdateOne) AddTeams(v ...*Team) *TaskUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.AddTeamIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (_u *TaskUpdateOne) SetOwnerID(id int) *TaskUpdateOne {
	_u.mutation.SetOwnerID(id)
	return _u
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (_u *TaskUpdateOne) SetNillableOwnerID(id *int) *TaskUpdateOne {
	if id != nil {
		_u = _u.SetOwnerID(*id)
	}
	return _u
}

// SetOwner sets the "owner" edge to the User entity.
func (_u *TaskUpdateOne) SetOwner(v *User) *TaskUpdateOne {
	return _u.SetOwnerID(v.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (_u *TaskUpdateOne) Mutation() *TaskMutation {
	return _u.mutation
}

// ClearTeams clears all "teams" edges to the Team entity.
func (_u *TaskUpdateOne) ClearTeams() *TaskUpdateOne {
	_u.mutation.ClearTeams()
	return _u
}

// RemoveTeamIDs removes the "teams" edge to Team entities by IDs.
func (_u *TaskUpdateOne) RemoveTeamIDs(ids ...int) *TaskUpdateOne {
	_u.mutation.RemoveTeamIDs(ids...)
	return _u
}

// RemoveTeams removes "teams" edges to Team entities.
func (_u *TaskUpdateOne) RemoveTeams(v ...*Team) *TaskUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.RemoveTeamIDs(ids...)
}

// ClearOwner clears the "owner" edge to the User entity.
func (_u *TaskUpdateOne) ClearOwner() *TaskUpdateOne {
	_u.mutation.ClearOwner()
	return _u
}

// Where appends a list predicates to the TaskUpdate builder.
func (_u *TaskUpdateOne) Where(ps ...predicate.Task) *TaskUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *TaskUpdateOne) Select(field string, fields ...string) *TaskUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated Task entity.
func (_u *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *TaskUpdateOne) check() error {
	if v, ok := _u.mutation.Title(); ok {
		if err := task.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Task.title": %w`, err)}
		}
	}
	if v, ok := _u.mutation.Status(); ok {
		if err := task.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Task.status": %w`, err)}
		}
	}
	return nil
}

func (_u *TaskUpdateOne) sqlSave(ctx context.Context) (_node *Task, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Task.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, task.FieldID)
		for _, f := range fields {
			if !task.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != task.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Title(); ok {
		_spec.SetField(task.FieldTitle, field.TypeString, value)
	}
	if value, ok := _u.mutation.Description(); ok {
		_spec.SetField(task.FieldDescription, field.TypeString, value)
	}
	if _u.mutation.DescriptionCleared() {
		_spec.ClearField(task.FieldDescription, field.TypeString)
	}
	if value, ok := _u.mutation.Status(); ok {
		_spec.SetField(task.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := _u.mutation.UUID(); ok {
		_spec.SetField(task.FieldUUID, field.TypeUUID, value)
	}
	if _u.mutation.UUIDCleared() {
		_spec.ClearField(task.FieldUUID, field.TypeUUID)
	}
	if _u.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   task.TeamsTable,
			Columns: task.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedTeamsIDs(); len(nodes) > 0 && !_u.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   task.TeamsTable,
			Columns: task.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.TeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   task.TeamsTable,
			Columns: task.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.OwnerTable,
			Columns: []string{task.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.OwnerTable,
			Columns: []string{task.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Task{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
