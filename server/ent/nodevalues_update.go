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
	"github.com/nullsploit01/iosync/ent/nodeapikey"
	"github.com/nullsploit01/iosync/ent/nodevalues"
	"github.com/nullsploit01/iosync/ent/predicate"
)

// NodeValuesUpdate is the builder for updating NodeValues entities.
type NodeValuesUpdate struct {
	config
	hooks    []Hook
	mutation *NodeValuesMutation
}

// Where appends a list predicates to the NodeValuesUpdate builder.
func (nvu *NodeValuesUpdate) Where(ps ...predicate.NodeValues) *NodeValuesUpdate {
	nvu.mutation.Where(ps...)
	return nvu
}

// SetValue sets the "value" field.
func (nvu *NodeValuesUpdate) SetValue(s string) *NodeValuesUpdate {
	nvu.mutation.SetValue(s)
	return nvu
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (nvu *NodeValuesUpdate) SetNillableValue(s *string) *NodeValuesUpdate {
	if s != nil {
		nvu.SetValue(*s)
	}
	return nvu
}

// SetUpdatedAt sets the "updated_at" field.
func (nvu *NodeValuesUpdate) SetUpdatedAt(t time.Time) *NodeValuesUpdate {
	nvu.mutation.SetUpdatedAt(t)
	return nvu
}

// SetNodeAPIKeyID sets the "node_api_key" edge to the NodeApiKey entity by ID.
func (nvu *NodeValuesUpdate) SetNodeAPIKeyID(id int) *NodeValuesUpdate {
	nvu.mutation.SetNodeAPIKeyID(id)
	return nvu
}

// SetNillableNodeAPIKeyID sets the "node_api_key" edge to the NodeApiKey entity by ID if the given value is not nil.
func (nvu *NodeValuesUpdate) SetNillableNodeAPIKeyID(id *int) *NodeValuesUpdate {
	if id != nil {
		nvu = nvu.SetNodeAPIKeyID(*id)
	}
	return nvu
}

// SetNodeAPIKey sets the "node_api_key" edge to the NodeApiKey entity.
func (nvu *NodeValuesUpdate) SetNodeAPIKey(n *NodeApiKey) *NodeValuesUpdate {
	return nvu.SetNodeAPIKeyID(n.ID)
}

// Mutation returns the NodeValuesMutation object of the builder.
func (nvu *NodeValuesUpdate) Mutation() *NodeValuesMutation {
	return nvu.mutation
}

// ClearNodeAPIKey clears the "node_api_key" edge to the NodeApiKey entity.
func (nvu *NodeValuesUpdate) ClearNodeAPIKey() *NodeValuesUpdate {
	nvu.mutation.ClearNodeAPIKey()
	return nvu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nvu *NodeValuesUpdate) Save(ctx context.Context) (int, error) {
	nvu.defaults()
	return withHooks(ctx, nvu.sqlSave, nvu.mutation, nvu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nvu *NodeValuesUpdate) SaveX(ctx context.Context) int {
	affected, err := nvu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nvu *NodeValuesUpdate) Exec(ctx context.Context) error {
	_, err := nvu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nvu *NodeValuesUpdate) ExecX(ctx context.Context) {
	if err := nvu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nvu *NodeValuesUpdate) defaults() {
	if _, ok := nvu.mutation.UpdatedAt(); !ok {
		v := nodevalues.UpdateDefaultUpdatedAt()
		nvu.mutation.SetUpdatedAt(v)
	}
}

func (nvu *NodeValuesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(nodevalues.Table, nodevalues.Columns, sqlgraph.NewFieldSpec(nodevalues.FieldID, field.TypeInt))
	if ps := nvu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nvu.mutation.Value(); ok {
		_spec.SetField(nodevalues.FieldValue, field.TypeString, value)
	}
	if value, ok := nvu.mutation.UpdatedAt(); ok {
		_spec.SetField(nodevalues.FieldUpdatedAt, field.TypeTime, value)
	}
	if nvu.mutation.NodeAPIKeyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   nodevalues.NodeAPIKeyTable,
			Columns: []string{nodevalues.NodeAPIKeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(nodeapikey.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nvu.mutation.NodeAPIKeyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   nodevalues.NodeAPIKeyTable,
			Columns: []string{nodevalues.NodeAPIKeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(nodeapikey.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nvu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nodevalues.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nvu.mutation.done = true
	return n, nil
}

// NodeValuesUpdateOne is the builder for updating a single NodeValues entity.
type NodeValuesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NodeValuesMutation
}

// SetValue sets the "value" field.
func (nvuo *NodeValuesUpdateOne) SetValue(s string) *NodeValuesUpdateOne {
	nvuo.mutation.SetValue(s)
	return nvuo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (nvuo *NodeValuesUpdateOne) SetNillableValue(s *string) *NodeValuesUpdateOne {
	if s != nil {
		nvuo.SetValue(*s)
	}
	return nvuo
}

// SetUpdatedAt sets the "updated_at" field.
func (nvuo *NodeValuesUpdateOne) SetUpdatedAt(t time.Time) *NodeValuesUpdateOne {
	nvuo.mutation.SetUpdatedAt(t)
	return nvuo
}

// SetNodeAPIKeyID sets the "node_api_key" edge to the NodeApiKey entity by ID.
func (nvuo *NodeValuesUpdateOne) SetNodeAPIKeyID(id int) *NodeValuesUpdateOne {
	nvuo.mutation.SetNodeAPIKeyID(id)
	return nvuo
}

// SetNillableNodeAPIKeyID sets the "node_api_key" edge to the NodeApiKey entity by ID if the given value is not nil.
func (nvuo *NodeValuesUpdateOne) SetNillableNodeAPIKeyID(id *int) *NodeValuesUpdateOne {
	if id != nil {
		nvuo = nvuo.SetNodeAPIKeyID(*id)
	}
	return nvuo
}

// SetNodeAPIKey sets the "node_api_key" edge to the NodeApiKey entity.
func (nvuo *NodeValuesUpdateOne) SetNodeAPIKey(n *NodeApiKey) *NodeValuesUpdateOne {
	return nvuo.SetNodeAPIKeyID(n.ID)
}

// Mutation returns the NodeValuesMutation object of the builder.
func (nvuo *NodeValuesUpdateOne) Mutation() *NodeValuesMutation {
	return nvuo.mutation
}

// ClearNodeAPIKey clears the "node_api_key" edge to the NodeApiKey entity.
func (nvuo *NodeValuesUpdateOne) ClearNodeAPIKey() *NodeValuesUpdateOne {
	nvuo.mutation.ClearNodeAPIKey()
	return nvuo
}

// Where appends a list predicates to the NodeValuesUpdate builder.
func (nvuo *NodeValuesUpdateOne) Where(ps ...predicate.NodeValues) *NodeValuesUpdateOne {
	nvuo.mutation.Where(ps...)
	return nvuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nvuo *NodeValuesUpdateOne) Select(field string, fields ...string) *NodeValuesUpdateOne {
	nvuo.fields = append([]string{field}, fields...)
	return nvuo
}

// Save executes the query and returns the updated NodeValues entity.
func (nvuo *NodeValuesUpdateOne) Save(ctx context.Context) (*NodeValues, error) {
	nvuo.defaults()
	return withHooks(ctx, nvuo.sqlSave, nvuo.mutation, nvuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nvuo *NodeValuesUpdateOne) SaveX(ctx context.Context) *NodeValues {
	node, err := nvuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nvuo *NodeValuesUpdateOne) Exec(ctx context.Context) error {
	_, err := nvuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nvuo *NodeValuesUpdateOne) ExecX(ctx context.Context) {
	if err := nvuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nvuo *NodeValuesUpdateOne) defaults() {
	if _, ok := nvuo.mutation.UpdatedAt(); !ok {
		v := nodevalues.UpdateDefaultUpdatedAt()
		nvuo.mutation.SetUpdatedAt(v)
	}
}

func (nvuo *NodeValuesUpdateOne) sqlSave(ctx context.Context) (_node *NodeValues, err error) {
	_spec := sqlgraph.NewUpdateSpec(nodevalues.Table, nodevalues.Columns, sqlgraph.NewFieldSpec(nodevalues.FieldID, field.TypeInt))
	id, ok := nvuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NodeValues.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nvuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nodevalues.FieldID)
		for _, f := range fields {
			if !nodevalues.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != nodevalues.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nvuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nvuo.mutation.Value(); ok {
		_spec.SetField(nodevalues.FieldValue, field.TypeString, value)
	}
	if value, ok := nvuo.mutation.UpdatedAt(); ok {
		_spec.SetField(nodevalues.FieldUpdatedAt, field.TypeTime, value)
	}
	if nvuo.mutation.NodeAPIKeyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   nodevalues.NodeAPIKeyTable,
			Columns: []string{nodevalues.NodeAPIKeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(nodeapikey.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nvuo.mutation.NodeAPIKeyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   nodevalues.NodeAPIKeyTable,
			Columns: []string{nodevalues.NodeAPIKeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(nodeapikey.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &NodeValues{config: nvuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nvuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nodevalues.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nvuo.mutation.done = true
	return _node, nil
}
