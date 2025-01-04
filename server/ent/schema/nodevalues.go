package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NodeValues holds the schema definition for the NodeValues entity.
type NodeValues struct {
	ent.Schema
}

// Fields of the NodeValues.
func (NodeValues) Fields() []ent.Field {
	return []ent.Field{
		field.String("value"),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the NodeValues.
func (NodeValues) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("node_api_key", NodeApiKey.Type).
			Ref("values").
			Unique(),
	}
}
