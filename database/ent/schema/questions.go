package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Questions holds the schema definition for the Questions entity.
type Questions struct {
	ent.Schema
}

// Fields of the Questions.
func (Questions) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").NotEmpty(),     // Teks pertanyaan
		field.String("language").NotEmpty(), // bahasa
		field.String("description"),         // bahasa
		field.Int("order"),                  // Urutan pertanyaan
	}
}

// Edges of the Questions.
func (Questions) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("history_answers", History_Answer.Type), // Relasi ke HistoryAnswer
	}
}

// Add created_at, updated_at, and deleted_at
func (Questions) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		DeleteMixin{},
	}
}
