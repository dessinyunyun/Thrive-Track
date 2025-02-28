package repository

import (
	"context"
	"go-gin/app/tools"
	"go-gin/app/user"
	"go-gin/database/ent"
	entUser "go-gin/database/ent/user"
	"time"

	googleUUID "github.com/google/uuid"
)

type UserRepository struct {
	db *ent.Client
}

func NewUserRepository(db *ent.Client) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAllUser(ctx context.Context, pagination *tools.Pagination, filter *user.FilterUser) ([]*user.UserResponse, *tools.Pagination, error) {
	userQuery := r.db.User.Query().
		Where(
			entUser.EmailContainsFold(filter.Email),
		)

	User, err := userQuery.Offset(pagination.Offset).Limit(pagination.Limit).All(ctx)
	if err != nil {
		return nil, nil, err
	}

	// Pagination
	count, _ := userQuery.Count(ctx)
	pagination.Count = int(count)
	pagination = tools.Paging(pagination)

	var users []*user.UserResponse

	for _, v := range User {
		users = append(users, &user.UserResponse{
			ID:       v.ID,
			Name:     v.Name,
			Username: v.Username,
			Email:    v.Email,
		})
	}

	return users, pagination, nil
}

func (r *UserRepository) GetDetailUser(ctx context.Context, ID googleUUID.UUID) (*user.UserResponse, error) {
	userQuery := r.db.User.Query().
		Where(
			entUser.IDEQ(ID),
		)

	if count, _ := userQuery.Count(ctx); count == 0 {
		return nil, nil
	}

	exec, _ := userQuery.First(ctx)

	return &user.UserResponse{
		ID:       exec.ID,
		Name:     exec.Name,
		Username: exec.Username,
		Email:    exec.Email,
	}, nil
}

func (r *UserRepository) CheckUserIdentifier(ctx context.Context, identifier string) (*user.UserResponseSensitiveCase, error) {
	userQuery := r.db.User.Query().
		Where(
			entUser.Or(
				entUser.UsernameEQ(identifier),
				entUser.EmailEQ(identifier),
			),
		)

	if count, _ := userQuery.Count(ctx); count == 0 {
		return nil, nil
	}

	exec, _ := userQuery.First(ctx)

	return &user.UserResponseSensitiveCase{
		ID:       exec.ID,
		Name:     exec.Name,
		Username: exec.Username,
		Email:    exec.Email,
		Password: exec.Password,
	}, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, form *user.UserForm) error {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return err
	}

	_, err = tx.User.Create().
		SetName(form.Name).
		SetEmail(form.Email).
		SetUsername(form.Username).
		SetPassword(form.Password).
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

func (r *UserRepository) UpdateUser(ctx context.Context, form *user.UserForm) error {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return err
	}

	err = tx.User.Update().
		Where(entUser.IDEQ(form.ID)).
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

func (r *UserRepository) DeleteUser(ctx context.Context, ID googleUUID.UUID) error {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return err
	}

	queryUser, err := tx.User.Query().Where(entUser.IDEQ(ID)).First(ctx)
	if err != nil {
		return err
	}

	deleted_at := time.Now()

	err = tx.User.Update().
		Where(
			entUser.IDEQ(queryUser.ID),
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
