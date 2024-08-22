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
		field.String("key").Unique(),
		field.Bool("is_active").Default(true),
		field.Time("revoked_at").Optional().Nillable(),
		field.Time("last_used").Default(time.Now),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("expires_at").Optional().Nillable(),
		field.String("description").Optional().Nillable(),
	}
}

// Edges of the ApiKey.
func (ApiKey) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("device", Device.Type).
			Ref("api_key").
			Unique(),
	}
}
