package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NodeUpdate holds the schema definition for the NodeUpdate entity.
type NodeUpdate struct {
	ent.Schema
}

// Fields of the NodeUpdate.
func (NodeUpdate) Fields() []ent.Field {
	return []ent.Field{
		field.String("value"),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the NodeUpdate.
func (NodeUpdate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("node", Node.Type).
			Ref("node_updates").
			Unique(),
	}
}
