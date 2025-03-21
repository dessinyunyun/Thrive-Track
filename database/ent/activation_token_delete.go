// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"go-gin/database/ent/activation_token"
	"go-gin/database/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ActivationTokenDelete is the builder for deleting a Activation_token entity.
type ActivationTokenDelete struct {
	config
	hooks    []Hook
	mutation *ActivationTokenMutation
}

// Where appends a list predicates to the ActivationTokenDelete builder.
func (atd *ActivationTokenDelete) Where(ps ...predicate.Activation_token) *ActivationTokenDelete {
	atd.mutation.Where(ps...)
	return atd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (atd *ActivationTokenDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, atd.sqlExec, atd.mutation, atd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (atd *ActivationTokenDelete) ExecX(ctx context.Context) int {
	n, err := atd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (atd *ActivationTokenDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(activation_token.Table, sqlgraph.NewFieldSpec(activation_token.FieldID, field.TypeUUID))
	if ps := atd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, atd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	atd.mutation.done = true
	return affected, err
}

// ActivationTokenDeleteOne is the builder for deleting a single Activation_token entity.
type ActivationTokenDeleteOne struct {
	atd *ActivationTokenDelete
}

// Where appends a list predicates to the ActivationTokenDelete builder.
func (atdo *ActivationTokenDeleteOne) Where(ps ...predicate.Activation_token) *ActivationTokenDeleteOne {
	atdo.atd.mutation.Where(ps...)
	return atdo
}

// Exec executes the deletion query.
func (atdo *ActivationTokenDeleteOne) Exec(ctx context.Context) error {
	n, err := atdo.atd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{activation_token.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (atdo *ActivationTokenDeleteOne) ExecX(ctx context.Context) {
	if err := atdo.Exec(ctx); err != nil {
		panic(err)
	}
}
