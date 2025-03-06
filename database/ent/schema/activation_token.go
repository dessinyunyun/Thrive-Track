package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Activation_token holds the schema definition for the Activation_tokens entity.
type Activation_token struct {
	ent.Schema
}

// Fields of the Activation_token.
func (Activation_token) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.UUID("user_id", uuid.UUID{}),
		field.String("token").MaxLen(255).NotEmpty().Unique(),
		field.Bool("isused").Default(false),
	}
}

// Edges of the Activation_token.
func (Activation_token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("activation_tokens").
			Unique().
			Required().
			Field("user_id"),
	}
}

// Add created_at, updated_at, and deleted_at
func (Activation_token) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		DeleteMixin{},
	}
}
