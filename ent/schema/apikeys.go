package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ApiKeys holds the schema definition for the ApiKeys entity.
type ApiKeys struct {
	ent.Schema
}

// Fields of the ApiKeys.
func (ApiKeys) Fields() []ent.Field {
	return []ent.Field{
		field.String("key"),
		field.Int("device_id"),
		field.Time("last_used"),
		field.Bool("is_active").Default(true),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the ApiKeys.
func (ApiKeys) Edges() []ent.Edge {
	return nil
}
