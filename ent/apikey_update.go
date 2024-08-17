// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"iosync/ent/apikey"
	"iosync/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ApiKeyUpdate is the builder for updating ApiKey entities.
type ApiKeyUpdate struct {
	config
	hooks    []Hook
	mutation *ApiKeyMutation
}

// Where appends a list predicates to the ApiKeyUpdate builder.
func (aku *ApiKeyUpdate) Where(ps ...predicate.ApiKey) *ApiKeyUpdate {
	aku.mutation.Where(ps...)
	return aku
}

// SetKey sets the "key" field.
func (aku *ApiKeyUpdate) SetKey(s string) *ApiKeyUpdate {
	aku.mutation.SetKey(s)
	return aku
}

// SetNillableKey sets the "key" field if the given value is not nil.
func (aku *ApiKeyUpdate) SetNillableKey(s *string) *ApiKeyUpdate {
	if s != nil {
		aku.SetKey(*s)
	}
	return aku
}

// SetDeviceID sets the "device_id" field.
func (aku *ApiKeyUpdate) SetDeviceID(i int) *ApiKeyUpdate {
	aku.mutation.ResetDeviceID()
	aku.mutation.SetDeviceID(i)
	return aku
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (aku *ApiKeyUpdate) SetNillableDeviceID(i *int) *ApiKeyUpdate {
	if i != nil {
		aku.SetDeviceID(*i)
	}
	return aku
}

// AddDeviceID adds i to the "device_id" field.
func (aku *ApiKeyUpdate) AddDeviceID(i int) *ApiKeyUpdate {
	aku.mutation.AddDeviceID(i)
	return aku
}

// SetIsActive sets the "is_active" field.
func (aku *ApiKeyUpdate) SetIsActive(b bool) *ApiKeyUpdate {
	aku.mutation.SetIsActive(b)
	return aku
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (aku *ApiKeyUpdate) SetNillableIsActive(b *bool) *ApiKeyUpdate {
	if b != nil {
		aku.SetIsActive(*b)
	}
	return aku
}

// SetLastUsed sets the "last_used" field.
func (aku *ApiKeyUpdate) SetLastUsed(t time.Time) *ApiKeyUpdate {
	aku.mutation.SetLastUsed(t)
	return aku
}

// SetNillableLastUsed sets the "last_used" field if the given value is not nil.
func (aku *ApiKeyUpdate) SetNillableLastUsed(t *time.Time) *ApiKeyUpdate {
	if t != nil {
		aku.SetLastUsed(*t)
	}
	return aku
}

// SetCreatedAt sets the "created_at" field.
func (aku *ApiKeyUpdate) SetCreatedAt(t time.Time) *ApiKeyUpdate {
	aku.mutation.SetCreatedAt(t)
	return aku
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aku *ApiKeyUpdate) SetNillableCreatedAt(t *time.Time) *ApiKeyUpdate {
	if t != nil {
		aku.SetCreatedAt(*t)
	}
	return aku
}

// SetUpdatedAt sets the "updated_at" field.
func (aku *ApiKeyUpdate) SetUpdatedAt(t time.Time) *ApiKeyUpdate {
	aku.mutation.SetUpdatedAt(t)
	return aku
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (aku *ApiKeyUpdate) SetNillableUpdatedAt(t *time.Time) *ApiKeyUpdate {
	if t != nil {
		aku.SetUpdatedAt(*t)
	}
	return aku
}

// Mutation returns the ApiKeyMutation object of the builder.
func (aku *ApiKeyUpdate) Mutation() *ApiKeyMutation {
	return aku.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aku *ApiKeyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, aku.sqlSave, aku.mutation, aku.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aku *ApiKeyUpdate) SaveX(ctx context.Context) int {
	affected, err := aku.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aku *ApiKeyUpdate) Exec(ctx context.Context) error {
	_, err := aku.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aku *ApiKeyUpdate) ExecX(ctx context.Context) {
	if err := aku.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aku *ApiKeyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(apikey.Table, apikey.Columns, sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeInt))
	if ps := aku.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aku.mutation.Key(); ok {
		_spec.SetField(apikey.FieldKey, field.TypeString, value)
	}
	if value, ok := aku.mutation.DeviceID(); ok {
		_spec.SetField(apikey.FieldDeviceID, field.TypeInt, value)
	}
	if value, ok := aku.mutation.AddedDeviceID(); ok {
		_spec.AddField(apikey.FieldDeviceID, field.TypeInt, value)
	}
	if value, ok := aku.mutation.IsActive(); ok {
		_spec.SetField(apikey.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := aku.mutation.LastUsed(); ok {
		_spec.SetField(apikey.FieldLastUsed, field.TypeTime, value)
	}
	if value, ok := aku.mutation.CreatedAt(); ok {
		_spec.SetField(apikey.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := aku.mutation.UpdatedAt(); ok {
		_spec.SetField(apikey.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aku.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apikey.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aku.mutation.done = true
	return n, nil
}

// ApiKeyUpdateOne is the builder for updating a single ApiKey entity.
type ApiKeyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ApiKeyMutation
}

// SetKey sets the "key" field.
func (akuo *ApiKeyUpdateOne) SetKey(s string) *ApiKeyUpdateOne {
	akuo.mutation.SetKey(s)
	return akuo
}

// SetNillableKey sets the "key" field if the given value is not nil.
func (akuo *ApiKeyUpdateOne) SetNillableKey(s *string) *ApiKeyUpdateOne {
	if s != nil {
		akuo.SetKey(*s)
	}
	return akuo
}

// SetDeviceID sets the "device_id" field.
func (akuo *ApiKeyUpdateOne) SetDeviceID(i int) *ApiKeyUpdateOne {
	akuo.mutation.ResetDeviceID()
	akuo.mutation.SetDeviceID(i)
	return akuo
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (akuo *ApiKeyUpdateOne) SetNillableDeviceID(i *int) *ApiKeyUpdateOne {
	if i != nil {
		akuo.SetDeviceID(*i)
	}
	return akuo
}

// AddDeviceID adds i to the "device_id" field.
func (akuo *ApiKeyUpdateOne) AddDeviceID(i int) *ApiKeyUpdateOne {
	akuo.mutation.AddDeviceID(i)
	return akuo
}

// SetIsActive sets the "is_active" field.
func (akuo *ApiKeyUpdateOne) SetIsActive(b bool) *ApiKeyUpdateOne {
	akuo.mutation.SetIsActive(b)
	return akuo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (akuo *ApiKeyUpdateOne) SetNillableIsActive(b *bool) *ApiKeyUpdateOne {
	if b != nil {
		akuo.SetIsActive(*b)
	}
	return akuo
}

// SetLastUsed sets the "last_used" field.
func (akuo *ApiKeyUpdateOne) SetLastUsed(t time.Time) *ApiKeyUpdateOne {
	akuo.mutation.SetLastUsed(t)
	return akuo
}

// SetNillableLastUsed sets the "last_used" field if the given value is not nil.
func (akuo *ApiKeyUpdateOne) SetNillableLastUsed(t *time.Time) *ApiKeyUpdateOne {
	if t != nil {
		akuo.SetLastUsed(*t)
	}
	return akuo
}

// SetCreatedAt sets the "created_at" field.
func (akuo *ApiKeyUpdateOne) SetCreatedAt(t time.Time) *ApiKeyUpdateOne {
	akuo.mutation.SetCreatedAt(t)
	return akuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (akuo *ApiKeyUpdateOne) SetNillableCreatedAt(t *time.Time) *ApiKeyUpdateOne {
	if t != nil {
		akuo.SetCreatedAt(*t)
	}
	return akuo
}

// SetUpdatedAt sets the "updated_at" field.
func (akuo *ApiKeyUpdateOne) SetUpdatedAt(t time.Time) *ApiKeyUpdateOne {
	akuo.mutation.SetUpdatedAt(t)
	return akuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (akuo *ApiKeyUpdateOne) SetNillableUpdatedAt(t *time.Time) *ApiKeyUpdateOne {
	if t != nil {
		akuo.SetUpdatedAt(*t)
	}
	return akuo
}

// Mutation returns the ApiKeyMutation object of the builder.
func (akuo *ApiKeyUpdateOne) Mutation() *ApiKeyMutation {
	return akuo.mutation
}

// Where appends a list predicates to the ApiKeyUpdate builder.
func (akuo *ApiKeyUpdateOne) Where(ps ...predicate.ApiKey) *ApiKeyUpdateOne {
	akuo.mutation.Where(ps...)
	return akuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (akuo *ApiKeyUpdateOne) Select(field string, fields ...string) *ApiKeyUpdateOne {
	akuo.fields = append([]string{field}, fields...)
	return akuo
}

// Save executes the query and returns the updated ApiKey entity.
func (akuo *ApiKeyUpdateOne) Save(ctx context.Context) (*ApiKey, error) {
	return withHooks(ctx, akuo.sqlSave, akuo.mutation, akuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (akuo *ApiKeyUpdateOne) SaveX(ctx context.Context) *ApiKey {
	node, err := akuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (akuo *ApiKeyUpdateOne) Exec(ctx context.Context) error {
	_, err := akuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (akuo *ApiKeyUpdateOne) ExecX(ctx context.Context) {
	if err := akuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (akuo *ApiKeyUpdateOne) sqlSave(ctx context.Context) (_node *ApiKey, err error) {
	_spec := sqlgraph.NewUpdateSpec(apikey.Table, apikey.Columns, sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeInt))
	id, ok := akuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ApiKey.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := akuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apikey.FieldID)
		for _, f := range fields {
			if !apikey.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != apikey.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := akuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := akuo.mutation.Key(); ok {
		_spec.SetField(apikey.FieldKey, field.TypeString, value)
	}
	if value, ok := akuo.mutation.DeviceID(); ok {
		_spec.SetField(apikey.FieldDeviceID, field.TypeInt, value)
	}
	if value, ok := akuo.mutation.AddedDeviceID(); ok {
		_spec.AddField(apikey.FieldDeviceID, field.TypeInt, value)
	}
	if value, ok := akuo.mutation.IsActive(); ok {
		_spec.SetField(apikey.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := akuo.mutation.LastUsed(); ok {
		_spec.SetField(apikey.FieldLastUsed, field.TypeTime, value)
	}
	if value, ok := akuo.mutation.CreatedAt(); ok {
		_spec.SetField(apikey.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := akuo.mutation.UpdatedAt(); ok {
		_spec.SetField(apikey.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &ApiKey{config: akuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, akuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apikey.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	akuo.mutation.done = true
	return _node, nil
}
