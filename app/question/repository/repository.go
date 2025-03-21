package repository

import (
	"context"
	"go-gin/app/question"
	"go-gin/app/tools"
	"go-gin/database/ent"
	entQuestion "go-gin/database/ent/question"
)

type Repository struct {
	db *ent.Client
}

func NewQuestionRepository(db *ent.Client) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAll(ctx context.Context, pagination *tools.Pagination, filter *question.Filter) ([]*question.Question, *tools.Pagination, error) {
	query := r.db.Question.Query().Order(entQuestion.ByOrder())
	if filter.CategoryId != 0 {
		query = query.Where(entQuestion.CategoryIDEQ(filter.CategoryId))
	}

	exec, err := query.Offset(pagination.Offset).Limit(pagination.Limit).All(ctx)
	if err != nil {
		return nil, nil, err
	}

	count, _ := query.Count(ctx)
	pagination.Count = int(count)
	pagination = tools.Paging(pagination)

	var questions []*question.Question

	for _, v := range exec {
		questions = append(questions, &question.Question{
			ID:          v.ID,
			Text:        v.Text,
			Language:    v.Language,
			Order:       v.Order,
			Example:     v.Example,
			Description: v.Description,
			CategoryID:  v.CategoryID,
		})
	}

	return questions, pagination, nil
}

func (r *Repository) GetDetail(ctx context.Context, order int, language string) (*question.Question, error) {

	query := r.db.Question.Query().
		Where(
			entQuestion.OrderEQ(order),
			entQuestion.LanguageEQ(language),
		)

	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, nil
	}

	exec, err := query.First(ctx)
	if err != nil {
		return nil, err
	}

	res := &question.Question{
		ID:          exec.ID,
		Text:        exec.Text,
		Language:    exec.Language,
		Order:       exec.Order,
		Example:     exec.Example,
		Description: exec.Description,
	}

	return res, nil
}

func (r *Repository) Create(ctx context.Context, form *question.Form) error {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Question.Create().
		SetText(form.Text).
		SetOrder(form.Order).
		SetLanguage(form.Language).
		SetDescription(form.Description).
		SetCategoryID(form.Category).
		SetExample(form.Example).
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
