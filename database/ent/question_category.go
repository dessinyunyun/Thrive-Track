// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-gin/database/ent/question_category"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Question_Category is the model entity for the Question_Category schema.
type Question_Category struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Order holds the value of the "order" field.
	Order int `json:"order,omitempty"`
	// Language holds the value of the "language" field.
	Language string `json:"language,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Question_CategoryQuery when eager-loading is set.
	Edges        Question_CategoryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// Question_CategoryEdges holds the relations/edges for other nodes in the graph.
type Question_CategoryEdges struct {
	// Questions holds the value of the questions edge.
	Questions []*Question `json:"questions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// QuestionsOrErr returns the Questions value or an error if the edge
// was not loaded in eager-loading.
func (e Question_CategoryEdges) QuestionsOrErr() ([]*Question, error) {
	if e.loadedTypes[0] {
		return e.Questions, nil
	}
	return nil, &NotLoadedError{edge: "questions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Question_Category) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case question_category.FieldID, question_category.FieldOrder:
			values[i] = new(sql.NullInt64)
		case question_category.FieldName, question_category.FieldLanguage:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Question_Category fields.
func (qc *Question_Category) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case question_category.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			qc.ID = int(value.Int64)
		case question_category.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				qc.Name = value.String
			}
		case question_category.FieldOrder:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order", values[i])
			} else if value.Valid {
				qc.Order = int(value.Int64)
			}
		case question_category.FieldLanguage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field language", values[i])
			} else if value.Valid {
				qc.Language = value.String
			}
		default:
			qc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Question_Category.
// This includes values selected through modifiers, order, etc.
func (qc *Question_Category) Value(name string) (ent.Value, error) {
	return qc.selectValues.Get(name)
}

// QueryQuestions queries the "questions" edge of the Question_Category entity.
func (qc *Question_Category) QueryQuestions() *QuestionQuery {
	return NewQuestionCategoryClient(qc.config).QueryQuestions(qc)
}

// Update returns a builder for updating this Question_Category.
// Note that you need to call Question_Category.Unwrap() before calling this method if this Question_Category
// was returned from a transaction, and the transaction was committed or rolled back.
func (qc *Question_Category) Update() *QuestionCategoryUpdateOne {
	return NewQuestionCategoryClient(qc.config).UpdateOne(qc)
}

// Unwrap unwraps the Question_Category entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (qc *Question_Category) Unwrap() *Question_Category {
	_tx, ok := qc.config.driver.(*txDriver)
	if !ok {
		panic("ent: Question_Category is not a transactional entity")
	}
	qc.config.driver = _tx.drv
	return qc
}

// String implements the fmt.Stringer.
func (qc *Question_Category) String() string {
	var builder strings.Builder
	builder.WriteString("Question_Category(")
	builder.WriteString(fmt.Sprintf("id=%v, ", qc.ID))
	builder.WriteString("name=")
	builder.WriteString(qc.Name)
	builder.WriteString(", ")
	builder.WriteString("order=")
	builder.WriteString(fmt.Sprintf("%v", qc.Order))
	builder.WriteString(", ")
	builder.WriteString("language=")
	builder.WriteString(qc.Language)
	builder.WriteByte(')')
	return builder.String()
}

// Question_Categories is a parsable slice of Question_Category.
type Question_Categories []*Question_Category
