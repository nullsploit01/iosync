package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.String("session_id").Unique(),
		field.String("username").NotEmpty(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
		field.Time("expires_at"),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return nil
}
