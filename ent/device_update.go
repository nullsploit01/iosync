// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"iosync/ent/apikey"
	"iosync/ent/device"
	"iosync/ent/predicate"
	"iosync/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeviceUpdate is the builder for updating Device entities.
type DeviceUpdate struct {
	config
	hooks    []Hook
	mutation *DeviceMutation
}

// Where appends a list predicates to the DeviceUpdate builder.
func (du *DeviceUpdate) Where(ps ...predicate.Device) *DeviceUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetName sets the "name" field.
func (du *DeviceUpdate) SetName(s string) *DeviceUpdate {
	du.mutation.SetName(s)
	return du
}

// SetNillableName sets the "name" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableName(s *string) *DeviceUpdate {
	if s != nil {
		du.SetName(*s)
	}
	return du
}

// SetIsActive sets the "is_active" field.
func (du *DeviceUpdate) SetIsActive(b bool) *DeviceUpdate {
	du.mutation.SetIsActive(b)
	return du
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableIsActive(b *bool) *DeviceUpdate {
	if b != nil {
		du.SetIsActive(*b)
	}
	return du
}

// SetCreatedAt sets the "created_at" field.
func (du *DeviceUpdate) SetCreatedAt(t time.Time) *DeviceUpdate {
	du.mutation.SetCreatedAt(t)
	return du
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableCreatedAt(t *time.Time) *DeviceUpdate {
	if t != nil {
		du.SetCreatedAt(*t)
	}
	return du
}

// SetUpdatedAt sets the "updated_at" field.
func (du *DeviceUpdate) SetUpdatedAt(t time.Time) *DeviceUpdate {
	du.mutation.SetUpdatedAt(t)
	return du
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableUpdatedAt(t *time.Time) *DeviceUpdate {
	if t != nil {
		du.SetUpdatedAt(*t)
	}
	return du
}

// SetUserID sets the "user" edge to the User entity by ID.
func (du *DeviceUpdate) SetUserID(id int) *DeviceUpdate {
	du.mutation.SetUserID(id)
	return du
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (du *DeviceUpdate) SetNillableUserID(id *int) *DeviceUpdate {
	if id != nil {
		du = du.SetUserID(*id)
	}
	return du
}

// SetUser sets the "user" edge to the User entity.
func (du *DeviceUpdate) SetUser(u *User) *DeviceUpdate {
	return du.SetUserID(u.ID)
}

// SetAPIKeyID sets the "api_key" edge to the ApiKey entity by ID.
func (du *DeviceUpdate) SetAPIKeyID(id int) *DeviceUpdate {
	du.mutation.SetAPIKeyID(id)
	return du
}

// SetNillableAPIKeyID sets the "api_key" edge to the ApiKey entity by ID if the given value is not nil.
func (du *DeviceUpdate) SetNillableAPIKeyID(id *int) *DeviceUpdate {
	if id != nil {
		du = du.SetAPIKeyID(*id)
	}
	return du
}

// SetAPIKey sets the "api_key" edge to the ApiKey entity.
func (du *DeviceUpdate) SetAPIKey(a *ApiKey) *DeviceUpdate {
	return du.SetAPIKeyID(a.ID)
}

// Mutation returns the DeviceMutation object of the builder.
func (du *DeviceUpdate) Mutation() *DeviceMutation {
	return du.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (du *DeviceUpdate) ClearUser() *DeviceUpdate {
	du.mutation.ClearUser()
	return du
}

// ClearAPIKey clears the "api_key" edge to the ApiKey entity.
func (du *DeviceUpdate) ClearAPIKey() *DeviceUpdate {
	du.mutation.ClearAPIKey()
	return du
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DeviceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DeviceUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DeviceUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DeviceUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DeviceUpdate) check() error {
	if v, ok := du.mutation.Name(); ok {
		if err := device.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Device.name": %w`, err)}
		}
	}
	return nil
}

func (du *DeviceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := du.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(device.Table, device.Columns, sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.SetField(device.FieldName, field.TypeString, value)
	}
	if value, ok := du.mutation.IsActive(); ok {
		_spec.SetField(device.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := du.mutation.CreatedAt(); ok {
		_spec.SetField(device.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.SetField(device.FieldUpdatedAt, field.TypeTime, value)
	}
	if du.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   device.UserTable,
			Columns: []string{device.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   device.UserTable,
			Columns: []string{device.UserColumn},
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
	if du.mutation.APIKeyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   device.APIKeyTable,
			Columns: []string{device.APIKeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.APIKeyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   device.APIKeyTable,
			Columns: []string{device.APIKeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{device.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DeviceUpdateOne is the builder for updating a single Device entity.
type DeviceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeviceMutation
}

// SetName sets the "name" field.
func (duo *DeviceUpdateOne) SetName(s string) *DeviceUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableName(s *string) *DeviceUpdateOne {
	if s != nil {
		duo.SetName(*s)
	}
	return duo
}

// SetIsActive sets the "is_active" field.
func (duo *DeviceUpdateOne) SetIsActive(b bool) *DeviceUpdateOne {
	duo.mutation.SetIsActive(b)
	return duo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableIsActive(b *bool) *DeviceUpdateOne {
	if b != nil {
		duo.SetIsActive(*b)
	}
	return duo
}

// SetCreatedAt sets the "created_at" field.
func (duo *DeviceUpdateOne) SetCreatedAt(t time.Time) *DeviceUpdateOne {
	duo.mutation.SetCreatedAt(t)
	return duo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableCreatedAt(t *time.Time) *DeviceUpdateOne {
	if t != nil {
		duo.SetCreatedAt(*t)
	}
	return duo
}

// SetUpdatedAt sets the "updated_at" field.
func (duo *DeviceUpdateOne) SetUpdatedAt(t time.Time) *DeviceUpdateOne {
	duo.mutation.SetUpdatedAt(t)
	return duo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableUpdatedAt(t *time.Time) *DeviceUpdateOne {
	if t != nil {
		duo.SetUpdatedAt(*t)
	}
	return duo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (duo *DeviceUpdateOne) SetUserID(id int) *DeviceUpdateOne {
	duo.mutation.SetUserID(id)
	return duo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableUserID(id *int) *DeviceUpdateOne {
	if id != nil {
		duo = duo.SetUserID(*id)
	}
	return duo
}

// SetUser sets the "user" edge to the User entity.
func (duo *DeviceUpdateOne) SetUser(u *User) *DeviceUpdateOne {
	return duo.SetUserID(u.ID)
}

// SetAPIKeyID sets the "api_key" edge to the ApiKey entity by ID.
func (duo *DeviceUpdateOne) SetAPIKeyID(id int) *DeviceUpdateOne {
	duo.mutation.SetAPIKeyID(id)
	return duo
}

// SetNillableAPIKeyID sets the "api_key" edge to the ApiKey entity by ID if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableAPIKeyID(id *int) *DeviceUpdateOne {
	if id != nil {
		duo = duo.SetAPIKeyID(*id)
	}
	return duo
}

// SetAPIKey sets the "api_key" edge to the ApiKey entity.
func (duo *DeviceUpdateOne) SetAPIKey(a *ApiKey) *DeviceUpdateOne {
	return duo.SetAPIKeyID(a.ID)
}

// Mutation returns the DeviceMutation object of the builder.
func (duo *DeviceUpdateOne) Mutation() *DeviceMutation {
	return duo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (duo *DeviceUpdateOne) ClearUser() *DeviceUpdateOne {
	duo.mutation.ClearUser()
	return duo
}

// ClearAPIKey clears the "api_key" edge to the ApiKey entity.
func (duo *DeviceUpdateOne) ClearAPIKey() *DeviceUpdateOne {
	duo.mutation.ClearAPIKey()
	return duo
}

// Where appends a list predicates to the DeviceUpdate builder.
func (duo *DeviceUpdateOne) Where(ps ...predicate.Device) *DeviceUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DeviceUpdateOne) Select(field string, fields ...string) *DeviceUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Device entity.
func (duo *DeviceUpdateOne) Save(ctx context.Context) (*Device, error) {
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DeviceUpdateOne) SaveX(ctx context.Context) *Device {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DeviceUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DeviceUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DeviceUpdateOne) check() error {
	if v, ok := duo.mutation.Name(); ok {
		if err := device.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Device.name": %w`, err)}
		}
	}
	return nil
}

func (duo *DeviceUpdateOne) sqlSave(ctx context.Context) (_node *Device, err error) {
	if err := duo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(device.Table, device.Columns, sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Device.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, device.FieldID)
		for _, f := range fields {
			if !device.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != device.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.SetField(device.FieldName, field.TypeString, value)
	}
	if value, ok := duo.mutation.IsActive(); ok {
		_spec.SetField(device.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := duo.mutation.CreatedAt(); ok {
		_spec.SetField(device.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.SetField(device.FieldUpdatedAt, field.TypeTime, value)
	}
	if duo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   device.UserTable,
			Columns: []string{device.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   device.UserTable,
			Columns: []string{device.UserColumn},
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
	if duo.mutation.APIKeyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   device.APIKeyTable,
			Columns: []string{device.APIKeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.APIKeyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   device.APIKeyTable,
			Columns: []string{device.APIKeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Device{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{device.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
