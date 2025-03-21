package usecase

import (
	"context"
	"go-gin/app/question_categories"

	"github.com/gin-gonic/gin"
)

type Usecase struct {
	repo question_categories.Repository
	ctx  context.Context
}

func NewCategoryQuestionUsecase(repo question_categories.Repository, ctx context.Context) *Usecase {
	return &Usecase{
		repo: repo,
		ctx:  ctx,
	}
}

// func (uc *Usecase) GetAll(c *gin.Context, filter *history_answer.Filter) ([]*history_answer.Response, *tools.Pagination, error) {
// 	pagination, err := tools.Paginate(c)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	if err := c.ShouldBindQuery(filter); err != nil {
// 		return nil, nil, err
// 	}

// 	result, pagination, err := uc.repo.GetAll(uc.ctx, pagination, filter)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	return result, pagination, nil
// }

// func (uc *Usecase) GetDetail(c *gin.Context) (*history_answer.Response, error) {
// 	historyAnswerId := c.Param("id")

// 	historyAnswerIdInt, _ := strconv.Atoi(historyAnswerId)

// 	result, err := uc.repo.GetDetail(uc.ctx, historyAnswerIdInt)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }

func (uc *Usecase) Create(c *gin.Context) error {
	var create *question_categories.Form
	err := c.ShouldBindJSON(&create)
	if err != nil {
		return err
	}

	err = uc.repo.Create(uc.ctx, create)
	if err != nil {
		return err
	}

	return nil
}

// func (uc *UserUsecase) GetDetailUser(c *gin.Context) (*user.UserResponse, error) {
// 	uuid, err := uuid.Parse(c.Param("id"))
// 	if err != nil {
// 		return nil, err
// 	}

// 	result, err := uc.repo.GetDetailUser(uc.ctx, uuid)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }

// func (uc *UserUsecase) UpdateUser(c *gin.Context) error {
// 	var updateUser *user.UserForm
// 	err := c.ShouldBindJSON(&updateUser)
// 	if err != nil {
// 		return err
// 	}

// 	ID, err := uuid.Parse(c.Param("id"))
// 	if err != nil {
// 		return err
// 	}

// 	updateUser.ID = ID

// 	err = uc.repo.UpdateUser(uc.ctx, updateUser)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (uc *UserUsecase) DeleteUser(c *gin.Context) error {
// 	ID, err := uuid.Parse(c.Param("id"))
// 	if err != nil {
// 		return err
// 	}

// 	err = uc.repo.DeleteUser(uc.ctx, ID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
