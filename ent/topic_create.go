// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"iosync/ent/apikey"
	"iosync/ent/device"
	"iosync/ent/topic"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TopicCreate is the builder for creating a Topic entity.
type TopicCreate struct {
	config
	mutation *TopicMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tc *TopicCreate) SetName(s string) *TopicCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetIsActive sets the "is_active" field.
func (tc *TopicCreate) SetIsActive(b bool) *TopicCreate {
	tc.mutation.SetIsActive(b)
	return tc
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (tc *TopicCreate) SetNillableIsActive(b *bool) *TopicCreate {
	if b != nil {
		tc.SetIsActive(*b)
	}
	return tc
}

// SetQos sets the "qos" field.
func (tc *TopicCreate) SetQos(i int) *TopicCreate {
	tc.mutation.SetQos(i)
	return tc
}

// SetNillableQos sets the "qos" field if the given value is not nil.
func (tc *TopicCreate) SetNillableQos(i *int) *TopicCreate {
	if i != nil {
		tc.SetQos(*i)
	}
	return tc
}

// SetRetain sets the "retain" field.
func (tc *TopicCreate) SetRetain(b bool) *TopicCreate {
	tc.mutation.SetRetain(b)
	return tc
}

// SetNillableRetain sets the "retain" field if the given value is not nil.
func (tc *TopicCreate) SetNillableRetain(b *bool) *TopicCreate {
	if b != nil {
		tc.SetRetain(*b)
	}
	return tc
}

// SetLastUsed sets the "last_used" field.
func (tc *TopicCreate) SetLastUsed(t time.Time) *TopicCreate {
	tc.mutation.SetLastUsed(t)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TopicCreate) SetCreatedAt(t time.Time) *TopicCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TopicCreate) SetNillableCreatedAt(t *time.Time) *TopicCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TopicCreate) SetUpdatedAt(t time.Time) *TopicCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TopicCreate) SetNillableUpdatedAt(t *time.Time) *TopicCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetDeviceID sets the "device" edge to the Device entity by ID.
func (tc *TopicCreate) SetDeviceID(id int) *TopicCreate {
	tc.mutation.SetDeviceID(id)
	return tc
}

// SetNillableDeviceID sets the "device" edge to the Device entity by ID if the given value is not nil.
func (tc *TopicCreate) SetNillableDeviceID(id *int) *TopicCreate {
	if id != nil {
		tc = tc.SetDeviceID(*id)
	}
	return tc
}

// SetDevice sets the "device" edge to the Device entity.
func (tc *TopicCreate) SetDevice(d *Device) *TopicCreate {
	return tc.SetDeviceID(d.ID)
}

// SetAPIKeyID sets the "api_key" edge to the ApiKey entity by ID.
func (tc *TopicCreate) SetAPIKeyID(id int) *TopicCreate {
	tc.mutation.SetAPIKeyID(id)
	return tc
}

// SetNillableAPIKeyID sets the "api_key" edge to the ApiKey entity by ID if the given value is not nil.
func (tc *TopicCreate) SetNillableAPIKeyID(id *int) *TopicCreate {
	if id != nil {
		tc = tc.SetAPIKeyID(*id)
	}
	return tc
}

// SetAPIKey sets the "api_key" edge to the ApiKey entity.
func (tc *TopicCreate) SetAPIKey(a *ApiKey) *TopicCreate {
	return tc.SetAPIKeyID(a.ID)
}

// Mutation returns the TopicMutation object of the builder.
func (tc *TopicCreate) Mutation() *TopicMutation {
	return tc.mutation
}

// Save creates the Topic in the database.
func (tc *TopicCreate) Save(ctx context.Context) (*Topic, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TopicCreate) SaveX(ctx context.Context) *Topic {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TopicCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TopicCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TopicCreate) defaults() {
	if _, ok := tc.mutation.IsActive(); !ok {
		v := topic.DefaultIsActive
		tc.mutation.SetIsActive(v)
	}
	if _, ok := tc.mutation.Qos(); !ok {
		v := topic.DefaultQos
		tc.mutation.SetQos(v)
	}
	if _, ok := tc.mutation.Retain(); !ok {
		v := topic.DefaultRetain
		tc.mutation.SetRetain(v)
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := topic.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := topic.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TopicCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Topic.name"`)}
	}
	if v, ok := tc.mutation.Name(); ok {
		if err := topic.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Topic.name": %w`, err)}
		}
	}
	if _, ok := tc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "Topic.is_active"`)}
	}
	if _, ok := tc.mutation.Qos(); !ok {
		return &ValidationError{Name: "qos", err: errors.New(`ent: missing required field "Topic.qos"`)}
	}
	if _, ok := tc.mutation.Retain(); !ok {
		return &ValidationError{Name: "retain", err: errors.New(`ent: missing required field "Topic.retain"`)}
	}
	if _, ok := tc.mutation.LastUsed(); !ok {
		return &ValidationError{Name: "last_used", err: errors.New(`ent: missing required field "Topic.last_used"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Topic.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Topic.updated_at"`)}
	}
	return nil
}

func (tc *TopicCreate) sqlSave(ctx context.Context) (*Topic, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TopicCreate) createSpec() (*Topic, *sqlgraph.CreateSpec) {
	var (
		_node = &Topic{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(topic.Table, sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(topic.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.IsActive(); ok {
		_spec.SetField(topic.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := tc.mutation.Qos(); ok {
		_spec.SetField(topic.FieldQos, field.TypeInt, value)
		_node.Qos = value
	}
	if value, ok := tc.mutation.Retain(); ok {
		_spec.SetField(topic.FieldRetain, field.TypeBool, value)
		_node.Retain = value
	}
	if value, ok := tc.mutation.LastUsed(); ok {
		_spec.SetField(topic.FieldLastUsed, field.TypeTime, value)
		_node.LastUsed = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(topic.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(topic.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := tc.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.DeviceTable,
			Columns: []string{topic.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.device_topics = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.APIKeyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.APIKeyTable,
			Columns: []string{topic.APIKeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.api_key_topics = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TopicCreateBulk is the builder for creating many Topic entities in bulk.
type TopicCreateBulk struct {
	config
	err      error
	builders []*TopicCreate
}

// Save creates the Topic entities in the database.
func (tcb *TopicCreateBulk) Save(ctx context.Context) ([]*Topic, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Topic, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TopicMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TopicCreateBulk) SaveX(ctx context.Context) []*Topic {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TopicCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TopicCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
