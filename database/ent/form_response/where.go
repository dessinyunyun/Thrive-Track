// Code generated by ent, DO NOT EDIT.

package form_response

import (
	"go-gin/database/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldEQ(FieldDeletedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldEQ(FieldUserID, v))
}

// TotalScore applies equality check predicate on the "total_score" field. It's identical to TotalScoreEQ.
func TotalScore(v uint8) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldEQ(FieldTotalScore, v))
}

// DepressionLevel applies equality check predicate on the "depression_level" field. It's identical to DepressionLevelEQ.
func DepressionLevel(v string) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldEQ(FieldDepressionLevel, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Form_Response {
	return predicate.Form_Response(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Form_Response {
	return predicate.Form_Response(sql.FieldNotNull(FieldDeletedAt))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldNotIn(FieldUserID, vs...))
}

// TotalScoreEQ applies the EQ predicate on the "total_score" field.
func TotalScoreEQ(v uint8) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldEQ(FieldTotalScore, v))
}

// TotalScoreNEQ applies the NEQ predicate on the "total_score" field.
func TotalScoreNEQ(v uint8) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldNEQ(FieldTotalScore, v))
}

// TotalScoreIn applies the In predicate on the "total_score" field.
func TotalScoreIn(vs ...uint8) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldIn(FieldTotalScore, vs...))
}

// TotalScoreNotIn applies the NotIn predicate on the "total_score" field.
func TotalScoreNotIn(vs ...uint8) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldNotIn(FieldTotalScore, vs...))
}

// TotalScoreGT applies the GT predicate on the "total_score" field.
func TotalScoreGT(v uint8) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldGT(FieldTotalScore, v))
}

// TotalScoreGTE applies the GTE predicate on the "total_score" field.
func TotalScoreGTE(v uint8) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldGTE(FieldTotalScore, v))
}

// TotalScoreLT applies the LT predicate on the "total_score" field.
func TotalScoreLT(v uint8) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldLT(FieldTotalScore, v))
}

// TotalScoreLTE applies the LTE predicate on the "total_score" field.
func TotalScoreLTE(v uint8) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldLTE(FieldTotalScore, v))
}

// DepressionLevelEQ applies the EQ predicate on the "depression_level" field.
func DepressionLevelEQ(v string) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldEQ(FieldDepressionLevel, v))
}

// DepressionLevelNEQ applies the NEQ predicate on the "depression_level" field.
func DepressionLevelNEQ(v string) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldNEQ(FieldDepressionLevel, v))
}

// DepressionLevelIn applies the In predicate on the "depression_level" field.
func DepressionLevelIn(vs ...string) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldIn(FieldDepressionLevel, vs...))
}

// DepressionLevelNotIn applies the NotIn predicate on the "depression_level" field.
func DepressionLevelNotIn(vs ...string) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldNotIn(FieldDepressionLevel, vs...))
}

// DepressionLevelGT applies the GT predicate on the "depression_level" field.
func DepressionLevelGT(v string) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldGT(FieldDepressionLevel, v))
}

// DepressionLevelGTE applies the GTE predicate on the "depression_level" field.
func DepressionLevelGTE(v string) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldGTE(FieldDepressionLevel, v))
}

// DepressionLevelLT applies the LT predicate on the "depression_level" field.
func DepressionLevelLT(v string) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldLT(FieldDepressionLevel, v))
}

// DepressionLevelLTE applies the LTE predicate on the "depression_level" field.
func DepressionLevelLTE(v string) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldLTE(FieldDepressionLevel, v))
}

// DepressionLevelContains applies the Contains predicate on the "depression_level" field.
func DepressionLevelContains(v string) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldContains(FieldDepressionLevel, v))
}

// DepressionLevelHasPrefix applies the HasPrefix predicate on the "depression_level" field.
func DepressionLevelHasPrefix(v string) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldHasPrefix(FieldDepressionLevel, v))
}

// DepressionLevelHasSuffix applies the HasSuffix predicate on the "depression_level" field.
func DepressionLevelHasSuffix(v string) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldHasSuffix(FieldDepressionLevel, v))
}

// DepressionLevelEqualFold applies the EqualFold predicate on the "depression_level" field.
func DepressionLevelEqualFold(v string) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldEqualFold(FieldDepressionLevel, v))
}

// DepressionLevelContainsFold applies the ContainsFold predicate on the "depression_level" field.
func DepressionLevelContainsFold(v string) predicate.Form_Response {
	return predicate.Form_Response(sql.FieldContainsFold(FieldDepressionLevel, v))
}

// HasHistoryAnswers applies the HasEdge predicate on the "history_answers" edge.
func HasHistoryAnswers() predicate.Form_Response {
	return predicate.Form_Response(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, HistoryAnswersTable, HistoryAnswersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHistoryAnswersWith applies the HasEdge predicate on the "history_answers" edge with a given conditions (other predicates).
func HasHistoryAnswersWith(preds ...predicate.History_Answer) predicate.Form_Response {
	return predicate.Form_Response(func(s *sql.Selector) {
		step := newHistoryAnswersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Form_Response {
	return predicate.Form_Response(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Form_Response {
	return predicate.Form_Response(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Form_Response) predicate.Form_Response {
	return predicate.Form_Response(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Form_Response) predicate.Form_Response {
	return predicate.Form_Response(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Form_Response) predicate.Form_Response {
	return predicate.Form_Response(sql.NotPredicates(p))
}