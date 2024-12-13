// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nullsploit01/iosync/ent/node"
	"github.com/nullsploit01/iosync/ent/predicate"
)

// NodeUpdate is the builder for updating Node entities.
type NodeUpdate struct {
	config
	hooks    []Hook
	mutation *NodeMutation
}

// Where appends a list predicates to the NodeUpdate builder.
func (nu *NodeUpdate) Where(ps ...predicate.Node) *NodeUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetName sets the "name" field.
func (nu *NodeUpdate) SetName(s string) *NodeUpdate {
	nu.mutation.SetName(s)
	return nu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (nu *NodeUpdate) SetNillableName(s *string) *NodeUpdate {
	if s != nil {
		nu.SetName(*s)
	}
	return nu
}

// SetDescription sets the "description" field.
func (nu *NodeUpdate) SetDescription(s string) *NodeUpdate {
	nu.mutation.SetDescription(s)
	return nu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (nu *NodeUpdate) SetNillableDescription(s *string) *NodeUpdate {
	if s != nil {
		nu.SetDescription(*s)
	}
	return nu
}

// SetIsActive sets the "is_active" field.
func (nu *NodeUpdate) SetIsActive(b bool) *NodeUpdate {
	nu.mutation.SetIsActive(b)
	return nu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (nu *NodeUpdate) SetNillableIsActive(b *bool) *NodeUpdate {
	if b != nil {
		nu.SetIsActive(*b)
	}
	return nu
}

// SetUpdatedAt sets the "updated_at" field.
func (nu *NodeUpdate) SetUpdatedAt(t time.Time) *NodeUpdate {
	nu.mutation.SetUpdatedAt(t)
	return nu
}

// Mutation returns the NodeMutation object of the builder.
func (nu *NodeUpdate) Mutation() *NodeMutation {
	return nu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NodeUpdate) Save(ctx context.Context) (int, error) {
	nu.defaults()
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NodeUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NodeUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NodeUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nu *NodeUpdate) defaults() {
	if _, ok := nu.mutation.UpdatedAt(); !ok {
		v := node.UpdateDefaultUpdatedAt()
		nu.mutation.SetUpdatedAt(v)
	}
}

func (nu *NodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(node.Table, node.Columns, sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.Name(); ok {
		_spec.SetField(node.FieldName, field.TypeString, value)
	}
	if value, ok := nu.mutation.Description(); ok {
		_spec.SetField(node.FieldDescription, field.TypeString, value)
	}
	if value, ok := nu.mutation.IsActive(); ok {
		_spec.SetField(node.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := nu.mutation.UpdatedAt(); ok {
		_spec.SetField(node.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{node.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NodeUpdateOne is the builder for updating a single Node entity.
type NodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NodeMutation
}

// SetName sets the "name" field.
func (nuo *NodeUpdateOne) SetName(s string) *NodeUpdateOne {
	nuo.mutation.SetName(s)
	return nuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillableName(s *string) *NodeUpdateOne {
	if s != nil {
		nuo.SetName(*s)
	}
	return nuo
}

// SetDescription sets the "description" field.
func (nuo *NodeUpdateOne) SetDescription(s string) *NodeUpdateOne {
	nuo.mutation.SetDescription(s)
	return nuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillableDescription(s *string) *NodeUpdateOne {
	if s != nil {
		nuo.SetDescription(*s)
	}
	return nuo
}

// SetIsActive sets the "is_active" field.
func (nuo *NodeUpdateOne) SetIsActive(b bool) *NodeUpdateOne {
	nuo.mutation.SetIsActive(b)
	return nuo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillableIsActive(b *bool) *NodeUpdateOne {
	if b != nil {
		nuo.SetIsActive(*b)
	}
	return nuo
}

// SetUpdatedAt sets the "updated_at" field.
func (nuo *NodeUpdateOne) SetUpdatedAt(t time.Time) *NodeUpdateOne {
	nuo.mutation.SetUpdatedAt(t)
	return nuo
}

// Mutation returns the NodeMutation object of the builder.
func (nuo *NodeUpdateOne) Mutation() *NodeMutation {
	return nuo.mutation
}

// Where appends a list predicates to the NodeUpdate builder.
func (nuo *NodeUpdateOne) Where(ps ...predicate.Node) *NodeUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NodeUpdateOne) Select(field string, fields ...string) *NodeUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Node entity.
func (nuo *NodeUpdateOne) Save(ctx context.Context) (*Node, error) {
	nuo.defaults()
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NodeUpdateOne) SaveX(ctx context.Context) *Node {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NodeUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NodeUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nuo *NodeUpdateOne) defaults() {
	if _, ok := nuo.mutation.UpdatedAt(); !ok {
		v := node.UpdateDefaultUpdatedAt()
		nuo.mutation.SetUpdatedAt(v)
	}
}

func (nuo *NodeUpdateOne) sqlSave(ctx context.Context) (_node *Node, err error) {
	_spec := sqlgraph.NewUpdateSpec(node.Table, node.Columns, sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Node.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, node.FieldID)
		for _, f := range fields {
			if !node.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != node.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.Name(); ok {
		_spec.SetField(node.FieldName, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Description(); ok {
		_spec.SetField(node.FieldDescription, field.TypeString, value)
	}
	if value, ok := nuo.mutation.IsActive(); ok {
		_spec.SetField(node.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := nuo.mutation.UpdatedAt(); ok {
		_spec.SetField(node.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Node{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{node.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}
