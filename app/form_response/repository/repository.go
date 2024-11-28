package repository

import (
	"context"
	"fmt"
	"go-gin/app/form_response"
	"go-gin/app/tools"
	"go-gin/database/ent"
	entFormResponse "go-gin/database/ent/form_response"

	"github.com/google/uuid"
)

type Repository struct {
	db *ent.Client
}

func NewFormResponseRepository(db *ent.Client) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAll(ctx context.Context, pagination *tools.Pagination, filter *form_response.Filter) ([]*form_response.Response, *tools.Pagination, error) {
	var userIdUUID *uuid.UUID
	if filter.UserId != "" {
		parsedUUID, err := uuid.Parse(filter.UserId)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid user_id format: %v", err)
		}
		userIdUUID = &parsedUUID
	}

	query := r.db.Form_Response.Query()

	if userIdUUID != nil {
		query = query.Where(entFormResponse.UserIDEQ(*userIdUUID))
	}

	formResponse, err := query.Offset(pagination.Offset).Limit(pagination.Limit).All(ctx)
	if err != nil {
		return nil, nil, err
	}

	count, _ := query.Count(ctx)
	pagination.Count = int(count)
	pagination = tools.Paging(pagination)

	var formResponses []*form_response.Response
	for _, v := range formResponse {
		formResponses = append(formResponses, &form_response.Response{
			ID:              v.ID,
			TotalScore:      v.TotalScore,
			DepressionLevel: v.DepressionLevel,
			UserId:          v.UserID,
		})
	}

	return formResponses, pagination, nil
}

func (r *Repository) GetDetail(ctx context.Context, id int) (*form_response.Response, error) {

	query := r.db.Form_Response.Query().
		Where(
			entFormResponse.IDEQ(id),
		)

	// Periksa apakah data ada
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, nil
	}

	// Ambil data pertama
	exec, err := query.First(ctx)
	if err != nil {
		return nil, err
	}

	// Return hasil dalam format yang sesuai
	return &form_response.Response{
		ID:              exec.ID,
		TotalScore:      exec.TotalScore,      // Ganti sesuai dengan nama kolom di entitas Question
		DepressionLevel: exec.DepressionLevel, // Ganti sesuai dengan nama kolom di entitas Question
		UserId:          exec.UserID,          // Ganti sesuai dengan nama kolom di entitas Question    // Ganti sesuai dengan nama kolom di entitas Question
	}, nil
}

func (r *Repository) Create(ctx context.Context, form *form_response.Form) error {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Form_Response.Create().
		SetDepressionLevel(form.DepressionLevel).
		SetTotalScore(form.TotalScore).
		SetUserID(form.UserId).
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

// func (r *UserRepository) UpdateUser(ctx context.Context, form *user.UserForm) error {
// 	tx, err := r.db.Tx(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	err = tx.User.Update().
// 		Where(entUser.IDEQ(form.ID)).
// 		SetName(form.Name).
// 		SetEmail(form.Email).
// 		SetUsername(form.Username).
// 		Exec(ctx)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (r *UserRepository) DeleteUser(ctx context.Context, ID uuid.UUID) error {
// 	tx, err := r.db.Tx(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	queryUser, err := tx.User.Query().Where(entUser.IDEQ(ID)).First(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	deleted_at := time.Now()

// 	err = tx.User.Update().
// 		Where(
// 			entUser.IDEQ(queryUser.ID),
// 		).
// 		SetNillableDeletedAt(&deleted_at).
// 		Exec(ctx)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (r *UserRepository) CheckUserIdentifier(ctx context.Context, identifier string) (*user.UserResponseSensitiveCase, error) {
// 	userQuery := r.db.User.Query().
// 		Where(
// 			entUser.Or(
// 				entUser.UsernameEQ(identifier),
// 				entUser.EmailEQ(identifier),
// 			),
// 		)

// 	if count, _ := userQuery.Count(ctx); count == 0 {
// 		return nil, nil
// 	}

// 	exec, _ := userQuery.First(ctx)

// 	return &user.UserResponseSensitiveCase{
// 		ID:       exec.ID,
// 		Name:     exec.Name,
// 		Username: exec.Username,
// 		Email:    exec.Email,
// 		Password: exec.Password,
// 	}, nil
// }

// func (r *UserRepository) Create(ctx context.Context, form *user.UserForm) error {
// 	tx, err := r.db.Tx(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = tx.User.Create().
// 		SetName(form.Name).
// 		SetEmail(form.Email).
// 		SetUsername(form.Username).
// 		SetPassword(form.Password).
// 		Save(ctx)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (r *UserRepository) UpdateUser(ctx context.Context, form *user.UserForm) error {
// 	tx, err := r.db.Tx(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	err = tx.User.Update().
// 		Where(entUser.IDEQ(form.ID)).
// 		SetName(form.Name).
// 		SetEmail(form.Email).
// 		SetUsername(form.Username).
// 		Exec(ctx)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (r *UserRepository) DeleteUser(ctx context.Context, ID uuid.UUID) error {
// 	tx, err := r.db.Tx(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	queryUser, err := tx.User.Query().Where(entUser.IDEQ(ID)).First(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	deleted_at := time.Now()

// 	err = tx.User.Update().
// 		Where(
// 			entUser.IDEQ(queryUser.ID),
// 		).
// 		SetNillableDeletedAt(&deleted_at).
// 		Exec(ctx)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return err
// 	}

// 	return nil
// }
