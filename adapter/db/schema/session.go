package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").
			NotEmpty().
			Unique(),
		field.String("ip").
			Optional(),
		field.Bool("is_valid").
			Default(true),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Int("user_id"),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Required().
			Ref("sessions").
			Field("user_id").
			Unique(),
	}
}
