package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Example holds the schema definition for the Example entity.
type Example struct {
	ent.Schema
}

// Fields of the Example.
func (Example) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.String("name").NotEmpty(),
		field.String("username").MaxLen(100).NotEmpty().Unique(),
		field.String("email").MaxLen(150).Optional().Nillable().Unique(),
	}
}

// Edges of the Example.
func (Example) Edges() []ent.Edge {
	return nil
}

// Add created_at, updated_at, and deleted_at
func (Example) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		DeleteMixin{},
	}
}
