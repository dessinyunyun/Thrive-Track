// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-gin/database/ent/category_questions"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CategoryQuestionsCreate is the builder for creating a Category_Questions entity.
type CategoryQuestionsCreate struct {
	config
	mutation *CategoryQuestionsMutation
	hooks    []Hook
}

// Mutation returns the CategoryQuestionsMutation object of the builder.
func (cqc *CategoryQuestionsCreate) Mutation() *CategoryQuestionsMutation {
	return cqc.mutation
}

// Save creates the Category_Questions in the database.
func (cqc *CategoryQuestionsCreate) Save(ctx context.Context) (*Category_Questions, error) {
	return withHooks(ctx, cqc.sqlSave, cqc.mutation, cqc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cqc *CategoryQuestionsCreate) SaveX(ctx context.Context) *Category_Questions {
	v, err := cqc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cqc *CategoryQuestionsCreate) Exec(ctx context.Context) error {
	_, err := cqc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cqc *CategoryQuestionsCreate) ExecX(ctx context.Context) {
	if err := cqc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cqc *CategoryQuestionsCreate) check() error {
	return nil
}

func (cqc *CategoryQuestionsCreate) sqlSave(ctx context.Context) (*Category_Questions, error) {
	if err := cqc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cqc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cqc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cqc.mutation.id = &_node.ID
	cqc.mutation.done = true
	return _node, nil
}

func (cqc *CategoryQuestionsCreate) createSpec() (*Category_Questions, *sqlgraph.CreateSpec) {
	var (
		_node = &Category_Questions{config: cqc.config}
		_spec = sqlgraph.NewCreateSpec(category_questions.Table, sqlgraph.NewFieldSpec(category_questions.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// CategoryQuestionsCreateBulk is the builder for creating many Category_Questions entities in bulk.
type CategoryQuestionsCreateBulk struct {
	config
	err      error
	builders []*CategoryQuestionsCreate
}

// Save creates the Category_Questions entities in the database.
func (cqcb *CategoryQuestionsCreateBulk) Save(ctx context.Context) ([]*Category_Questions, error) {
	if cqcb.err != nil {
		return nil, cqcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cqcb.builders))
	nodes := make([]*Category_Questions, len(cqcb.builders))
	mutators := make([]Mutator, len(cqcb.builders))
	for i := range cqcb.builders {
		func(i int, root context.Context) {
			builder := cqcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CategoryQuestionsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cqcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cqcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, cqcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cqcb *CategoryQuestionsCreateBulk) SaveX(ctx context.Context) []*Category_Questions {
	v, err := cqcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cqcb *CategoryQuestionsCreateBulk) Exec(ctx context.Context) error {
	_, err := cqcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cqcb *CategoryQuestionsCreateBulk) ExecX(ctx context.Context) {
	if err := cqcb.Exec(ctx); err != nil {
		panic(err)
	}
}