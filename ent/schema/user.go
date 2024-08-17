package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique(),
		field.String("password"),
		field.String("name"),
		field.Bool("is_active").Default(true),
		field.Time("last_login").Default(time.Now),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("devices", Device.Type),
		edge.To("sessions", Session.Type),
	}
}

// Index of the User.

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username"),
	}
}
