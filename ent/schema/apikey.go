package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
		field.Bool("is_active").Default(true),
		field.Time("last_used").Default(time.Now),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the ApiKeys.
func (ApiKey) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("device", Device.Type).
			Ref("api_keys").
			Unique(),
	}
}
