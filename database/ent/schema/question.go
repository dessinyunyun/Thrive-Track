package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Questions holds the schema definition for the Questions entity.
type Question struct {
	ent.Schema
}

// Fields of the Questions.
func (Question) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").NotEmpty(),     // Teks pertanyaan
		field.String("language").NotEmpty(), // Bahasa
		field.String("description"),         // Deskripsi
		field.String("example"),             // Contoh
		field.Int("order"),                  // Urutan pertanyaan
		field.Int("category_id"),            // ID kategori pertanyaan
	}
}

// Edges of the Questions.
func (Question) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("history_answers", History_Answer.Type), // Relasi ke History_Answer
		edge.From("category", Question_Category.Type).
			Ref("questions").
			Unique().
			Required().
			Field("category_id"),
	}
}

// Add created_at, updated_at, and deleted_at
func (Question) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		DeleteMixin{},
	}
}
