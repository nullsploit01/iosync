package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NodeApiKey holds the schema definition for the NodeApiKey entity.
type NodeApiKey struct {
	ent.Schema
}

// Fields of the NodeApiKey.
func (NodeApiKey) Fields() []ent.Field {
	return []ent.Field{
		field.String("api_key"),
		field.String("description"),
		field.Bool("is_revoked").Default(false),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the NodeApiKey.
func (NodeApiKey) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("node", Node.Type).
			Ref("api_keys").
			Unique(),
	}
}
