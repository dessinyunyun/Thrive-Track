package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Form_Response holds the schema definition for the Form_Response entity.
type Form_Response struct {
	ent.Schema
}

// Fields of the Form_Response.
func (Form_Response) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}),             // Foreign key ke User
		field.Uint8("total_score").Max(100).Default(0), // Total skor
		field.String("depression_level").MaxLen(100).NotEmpty(),
	}
}

// Edges of the Form_Response.
func (Form_Response) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("history_answers", History_Answer.Type), // Menyimpan foreign key di History_Answer
		edge.From("user", User.Type).
			Ref("form_responses").
			Unique().
			Required().
			Field("user_id"),
	}
}

// Add created_at, updated_at, and deleted_at
func (Form_Response) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		DeleteMixin{},
	}
}
