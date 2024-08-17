package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ApiKey holds the schema definition for the ApiKey entity.
type ApiKey struct {
	ent.Schema
}

// Fields of the ApiKeys.
func (ApiKey) Fields() []ent.Field {
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
func (ApiKey) Edges() []ent.Edge {
	return nil
}
