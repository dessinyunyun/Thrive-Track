package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.String("name").NotEmpty(),
		field.String("username").MaxLen(100).NotEmpty().Unique(),
		field.String("email").MaxLen(150).Optional().Nillable().Unique(),
		field.String("password").
			Sensitive(). // Menandai bahwa ini adalah field sensitif
			NotEmpty().
			MaxLen(255),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("form_responses", Form_Response.Type),
	}
}

// Add created_at, updated_at, and deleted_at
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		DeleteMixin{},
	}
}
