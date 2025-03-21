package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Question_Category holds the schema definition for the Question_Category entity.
type Question_Category struct {
	ent.Schema
}

// Fields of the Question_Category.
func (Question_Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Int("order"),
		field.String("language").NotEmpty(),
	}
}

// Edges of the Question_Category.
func (Question_Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("questions", Question.Type),
	}
}
