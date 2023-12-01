// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"testEntGo/ent/people"
	"testEntGo/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PeopleUpdate is the builder for updating People entities.
type PeopleUpdate struct {
	config
	hooks    []Hook
	mutation *PeopleMutation
}

// Where appends a list predicates to the PeopleUpdate builder.
func (pu *PeopleUpdate) Where(ps ...predicate.People) *PeopleUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PeopleUpdate) SetName(s string) *PeopleUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PeopleUpdate) SetNillableName(s *string) *PeopleUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetLastName sets the "last_name" field.
func (pu *PeopleUpdate) SetLastName(s string) *PeopleUpdate {
	pu.mutation.SetLastName(s)
	return pu
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (pu *PeopleUpdate) SetNillableLastName(s *string) *PeopleUpdate {
	if s != nil {
		pu.SetLastName(*s)
	}
	return pu
}

// SetAge sets the "age" field.
func (pu *PeopleUpdate) SetAge(i int) *PeopleUpdate {
	pu.mutation.ResetAge()
	pu.mutation.SetAge(i)
	return pu
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (pu *PeopleUpdate) SetNillableAge(i *int) *PeopleUpdate {
	if i != nil {
		pu.SetAge(*i)
	}
	return pu
}

// AddAge adds i to the "age" field.
func (pu *PeopleUpdate) AddAge(i int) *PeopleUpdate {
	pu.mutation.AddAge(i)
	return pu
}

// Mutation returns the PeopleMutation object of the builder.
func (pu *PeopleUpdate) Mutation() *PeopleMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PeopleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PeopleUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PeopleUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PeopleUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PeopleUpdate) check() error {
	if v, ok := pu.mutation.Age(); ok {
		if err := people.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "People.age": %w`, err)}
		}
	}
	return nil
}

func (pu *PeopleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(people.Table, people.Columns, sqlgraph.NewFieldSpec(people.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(people.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.LastName(); ok {
		_spec.SetField(people.FieldLastName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Age(); ok {
		_spec.SetField(people.FieldAge, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedAge(); ok {
		_spec.AddField(people.FieldAge, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{people.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PeopleUpdateOne is the builder for updating a single People entity.
type PeopleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PeopleMutation
}

// SetName sets the "name" field.
func (puo *PeopleUpdateOne) SetName(s string) *PeopleUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PeopleUpdateOne) SetNillableName(s *string) *PeopleUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetLastName sets the "last_name" field.
func (puo *PeopleUpdateOne) SetLastName(s string) *PeopleUpdateOne {
	puo.mutation.SetLastName(s)
	return puo
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (puo *PeopleUpdateOne) SetNillableLastName(s *string) *PeopleUpdateOne {
	if s != nil {
		puo.SetLastName(*s)
	}
	return puo
}

// SetAge sets the "age" field.
func (puo *PeopleUpdateOne) SetAge(i int) *PeopleUpdateOne {
	puo.mutation.ResetAge()
	puo.mutation.SetAge(i)
	return puo
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (puo *PeopleUpdateOne) SetNillableAge(i *int) *PeopleUpdateOne {
	if i != nil {
		puo.SetAge(*i)
	}
	return puo
}

// AddAge adds i to the "age" field.
func (puo *PeopleUpdateOne) AddAge(i int) *PeopleUpdateOne {
	puo.mutation.AddAge(i)
	return puo
}

// Mutation returns the PeopleMutation object of the builder.
func (puo *PeopleUpdateOne) Mutation() *PeopleMutation {
	return puo.mutation
}

// Where appends a list predicates to the PeopleUpdate builder.
func (puo *PeopleUpdateOne) Where(ps ...predicate.People) *PeopleUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PeopleUpdateOne) Select(field string, fields ...string) *PeopleUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated People entity.
func (puo *PeopleUpdateOne) Save(ctx context.Context) (*People, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PeopleUpdateOne) SaveX(ctx context.Context) *People {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PeopleUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PeopleUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PeopleUpdateOne) check() error {
	if v, ok := puo.mutation.Age(); ok {
		if err := people.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "People.age": %w`, err)}
		}
	}
	return nil
}

func (puo *PeopleUpdateOne) sqlSave(ctx context.Context) (_node *People, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(people.Table, people.Columns, sqlgraph.NewFieldSpec(people.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "People.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, people.FieldID)
		for _, f := range fields {
			if !people.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != people.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(people.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.LastName(); ok {
		_spec.SetField(people.FieldLastName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Age(); ok {
		_spec.SetField(people.FieldAge, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedAge(); ok {
		_spec.AddField(people.FieldAge, field.TypeInt, value)
	}
	_node = &People{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{people.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}