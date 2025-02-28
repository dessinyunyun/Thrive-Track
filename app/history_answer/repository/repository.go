package repository

import (
	"context"
	"go-gin/app/history_answer"
	"go-gin/app/tools"
	"go-gin/database/ent"
	entHistoryResponse "go-gin/database/ent/history_answer"
)

type Repository struct {
	db *ent.Client
}

func NewHistoryAnswerRepository(db *ent.Client) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAll(ctx context.Context, pagination *tools.Pagination, filter *history_answer.Filter) ([]*history_answer.Response, *tools.Pagination, error) {
	query := r.db.History_Answer.Query()

	// Tambahkan kondisi filter jika ada
	if filter != nil {
		if filter.FormResponseId != 0 {
			query = query.Where(entHistoryResponse.FormResponseIDEQ(filter.FormResponseId))
		}
		if filter.QuestionId != 0 {
			query = query.Where(entHistoryResponse.QuestionIDEQ(filter.QuestionId))
		}
		if filter.QuestionId != 0 {
			query = query.Where(entHistoryResponse.QuestionIDEQ(filter.QuestionId))
		}
	}

	// Filter by DeletedAt is Nil (Always Applied)
	query = query.Where(entHistoryResponse.DeletedAtIsNil())

	historyAnswer, err := query.Offset(pagination.Offset).Limit(pagination.Limit).All(ctx)
	if err != nil {
		return nil, nil, err
	}

	count, _ := query.Count(ctx)
	pagination.Count = int(count)
	pagination = tools.Paging(pagination)

	var historyAnswers []*history_answer.Response
	for _, v := range historyAnswer {
		historyAnswers = append(historyAnswers, &history_answer.Response{
			ID:             v.ID,
			FormResponseId: v.FormResponseID,
			QuestionId:     v.QuestionID,
			Answer:         v.Answer,
		})
	}

	return historyAnswers, pagination, nil
}

func (r *Repository) GetDetail(ctx context.Context, formResponseid int) (*history_answer.Response, error) {

	query := r.db.History_Answer.Query().
		Where(
			entHistoryResponse.FormResponseIDEQ(formResponseid),
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
	return &history_answer.Response{
		ID:             exec.ID,
		FormResponseId: exec.FormResponseID,
		QuestionId:     exec.QuestionID,
		Answer:         exec.Answer, // Ganti sesuai dengan nama kolom di entitas Question    // Ganti sesuai dengan nama kolom di entitas Question
	}, nil
}

func (r *Repository) Create(ctx context.Context, form *history_answer.Form) error {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return err
	}

	_, err = tx.History_Answer.Create().
		SetFormResponseID(form.FormResponseId).
		SetQuestionID(form.QuestionId).
		SetAnswer(form.Answer).
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
