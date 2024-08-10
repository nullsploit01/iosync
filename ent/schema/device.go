package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Device holds the schema definition for the Device entity.
type Device struct {
	ent.Schema
}

// Fields of the Device.
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").NotEmpty(),
		field.String("name").NotEmpty(),
		field.Bool("is_active").Default(true),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return nil
}
