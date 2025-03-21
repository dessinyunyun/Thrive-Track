package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TOKEN holds the schema definition for the TOKEN entity.
type Token struct {
	ent.Schema
}

// Fields of the TOKEN.
func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.UUID("user_id", uuid.UUID{}),
		field.String("access_token").MaxLen(520).NotEmpty(),
		field.String("refresh_token").MaxLen(520).NotEmpty(),
		field.Time("access_token_expires_at"),
		field.Time("refresh_token_expires_at"),
		field.Bool("revoked"),
	}
}

// Edges of the TOKEN.
func (Token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("tokens").
			Unique().
			Required().
			Field("user_id"),
	}
}
