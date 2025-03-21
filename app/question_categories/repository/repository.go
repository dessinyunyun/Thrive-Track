package repository

import (
	"context"
	"go-gin/app/question_categories"
	"go-gin/database/ent"
	entQC "go-gin/database/ent/question_category"
)

type Repository struct {
	db *ent.Client
}

func NewCategoryQuestionRepository(db *ent.Client) *Repository {
	return &Repository{
		db: db,
	}
}

// func (r *Repository) GetAll(ctx context.Context, pagination *tools.Pagination, filter *history_answer.Filter) ([]*history_answer.Response, *tools.Pagination, error) {
// 	query := r.db.History_Answer.Query()

// 	// Tambahkan kondisi filter jika ada
// 	if filter != nil {
// 		if filter.FormResponseId != 0 {
// 			query = query.Where(entHistoryResponse.FormResponseIDEQ(filter.FormResponseId))
// 		}
// 		if filter.QuestionId != 0 {
// 			query = query.Where(entHistoryResponse.QuestionIDEQ(filter.QuestionId))
// 		}
// 		if filter.QuestionId != 0 {
// 			query = query.Where(entHistoryResponse.QuestionIDEQ(filter.QuestionId))
// 		}
// 	}

// 	// Filter by DeletedAt is Nil (Always Applied)
// 	query = query.Where(entHistoryResponse.DeletedAtIsNil())

// 	historyAnswer, err := query.Offset(pagination.Offset).Limit(pagination.Limit).All(ctx)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	count, _ := query.Count(ctx)
// 	pagination.Count = int(count)
// 	pagination = tools.Paging(pagination)

// 	var historyAnswers []*history_answer.Response
// 	for _, v := range historyAnswer {
// 		historyAnswers = append(historyAnswers, &history_answer.Response{
// 			ID:             v.ID,
// 			FormResponseId: v.FormResponseID,
// 			QuestionId:     v.QuestionID,
// 			Answer:         v.Answer,
// 		})
// 	}

// 	return historyAnswers, pagination, nil
// }

func (r *Repository) GetDetail(ctx context.Context, language string, qcOrder int) (*question_categories.Response, error) {
	query := r.db.Question_Category.Query().
		Where(
			entQC.LanguageEQ(language),
			entQC.OrderEQ(qcOrder),
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
	return &question_categories.Response{
		ID:           exec.ID,
		CategoryName: exec.Name,
		Order:        exec.Order,
		Language:     exec.Language, // Ganti sesuai dengan nama kolom di entitas Question    // Ganti sesuai dengan nama kolom di entitas Question
	}, nil
}

func (r *Repository) Create(ctx context.Context, form *question_categories.Form) error {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Question_Category.Create().
		SetLanguage(form.Language).
		SetName(form.CategoryName).
		SetOrder(form.Order).
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
