package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// History_Answer holds the schema definition for the History_Answer entity.
type History_Answer struct {
	ent.Schema
}

// Fields of the History_Answer.
func (History_Answer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("form_response_id").Positive(), // Foreign key ke Form_Response
		field.Int("question_id").Positive(),      // Foreign key ke Questions
		field.Int("answer").Range(0, 4),          // Jawaban antara 0-4
	}
}

// Edges of the History_Answer.
func (History_Answer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("form_response", Form_Response.Type).
			Ref("history_answers").
			Field("form_response_id").
			Unique().   // Gunakan form_response_id sebagai foreign key
			Required(), // Satu History_Answer hanya ke satu Form_Response
		edge.From("question", Questions.Type).
			Ref("history_answers").
			Field("question_id").
			Unique().
			Required(),
	}
}

// Add created_at, updated_at, and deleted_at
func (History_Answer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		DeleteMixin{},
	}
}
