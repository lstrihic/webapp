package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").
			NotEmpty().
			Unique(),
		field.String("username").
			NotEmpty().
			Unique(),
		field.String("password").
			NotEmpty(),
		field.String("token_key").
			Optional(),
		field.Bool("is_banned").
			Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sessions", Session.Type),
	}
}
