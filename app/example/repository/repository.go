package repository

import (
	"context"
	"go-gin/app/example"
	"go-gin/app/tools"
	"go-gin/database/ent"
	entExample "go-gin/database/ent/example"
	"time"

	"github.com/google/uuid"
)

type ExampleRepository struct {
	db *ent.Client
}

func NewExampleRepository(db *ent.Client) *ExampleRepository {
	return &ExampleRepository{
		db: db,
	}
}

func (r *ExampleRepository) GetAllExample(ctx context.Context, pagination *tools.Pagination, filter *example.FilterExample) ([]*example.ExampleResponse, *tools.Pagination, error) {
	exampleQuery := r.db.Example.Query().
		Where(
			entExample.DeletedAtIsNil(),
		)

	Example, err := exampleQuery.Offset(pagination.Offset).Limit(pagination.Limit).All(ctx)
	if err != nil {
		return nil, nil, err
	}

	// Pagination
	count, _ := exampleQuery.Count(ctx)
	pagination.Count = int(count)
	pagination = tools.Paging(pagination)

	var examples []*example.ExampleResponse

	for _, v := range Example {
		examples = append(examples, &example.ExampleResponse{
			ID:       v.ID,
			Name:     v.Name,
			Username: v.Username,
			Email:    v.Email,
		})
	}

	return examples, pagination, nil
}

func (r *ExampleRepository) GetDetailExample(ctx context.Context, ID uuid.UUID) (*example.ExampleResponse, error) {
	exampleQuery := r.db.Example.Query().
		Where(
			entExample.IDEQ(ID),
		)

	if count, _ := exampleQuery.Count(ctx); count == 0 {
		return nil, nil
	}

	exec, _ := exampleQuery.First(ctx)

	return &example.ExampleResponse{
		ID:       exec.ID,
		Name:     exec.Name,
		Username: exec.Username,
		Email:    exec.Email,
	}, nil
}

func (r *ExampleRepository) CreateExample(ctx context.Context, form *example.ExampleForm) error {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Example.Create().
		SetName(form.Name).
		SetEmail(form.Email).
		SetUsername(form.Username).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *ExampleRepository) UpdateExample(ctx context.Context, form *example.ExampleForm) error {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return err
	}

	err = tx.Example.Update().
		Where(entExample.IDEQ(form.ID)).
		SetName(form.Name).
		SetEmail(form.Email).
		SetUsername(form.Username).
		Exec(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *ExampleRepository) DeleteExample(ctx context.Context, ID uuid.UUID) error {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return err
	}

	queryExample, err := tx.Example.Query().Where(entExample.IDEQ(ID)).First(ctx)
	if err != nil {
		return err
	}

	deleted_at := time.Now()

	err = tx.Example.Update().
		Where(
			entExample.IDEQ(queryExample.ID),
		).
		SetNillableDeletedAt(&deleted_at).
		Exec(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
