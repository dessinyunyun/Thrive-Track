package repository

import (
	"context"
	"go-gin/app/question"
	"go-gin/database/ent"
	entQuestion "go-gin/database/ent/questions"
)

type Repository struct {
	db *ent.Client
}

func NewQuestionRepository(db *ent.Client) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetDetail(ctx context.Context, order int, language string) (*question.Response, error) {

	query := r.db.Questions.Query().
		Where(
			entQuestion.OrderEQ(order),
			entQuestion.LanguageEQ(language),
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
	return &question.Response{
		ID:          exec.ID,
		Text:        exec.Text,        // Ganti sesuai dengan nama kolom di entitas Question
		Language:    exec.Language,    // Ganti sesuai dengan nama kolom di entitas Question
		Order:       exec.Order,       // Ganti sesuai dengan nama kolom di entitas Question
		Description: exec.Description, // Ganti sesuai dengan nama kolom di entitas Question
	}, nil
}

func (r *Repository) Create(ctx context.Context, form *question.Form) error {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Questions.Create().
		SetText(form.Text).
		SetOrder(form.Order).
		SetLanguage(form.Language).
		SetDescription(form.Description).
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

// func (r *Repository) GetAllUser(ctx context.Context, pagination *tools.Pagination, filter *question.Filter) ([]*question.Response, *tools.Pagination, error) {
// 	// userQuery := r.db.User.Query().
// 	// 	Where(
// 	// 		entQuestion.Order(filter.Email),
// 	// 	)

// 	// User, err := userQuery.Offset(pagination.Offset).Limit(pagination.Limit).All(ctx)
// 	// if err != nil {
// 	// 	return nil, nil, err
// 	// }

// 	// Pagination
// 	count, _ := userQuery.Count(ctx)
// 	pagination.Count = int(count)
// 	pagination = tools.Paging(pagination)

// 	var users []*user.UserResponse

// 	for _, v := range User {
// 		users = append(users, &question.Response{
// 			ID:       v.ID,
// 			Text:     v.Name,
// 			Language: v.Username,
// 			Order:    v.Email,
// 		})
// 	}

// 	return users, pagination, nil
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
