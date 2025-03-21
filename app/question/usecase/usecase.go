package usecase

import (
	"context"
	"fmt"
	"go-gin/app/question"
	"go-gin/app/question_categories"
	"go-gin/app/tools"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Usecase struct {
	repo   question.QuestionRepository
	qcRepo question_categories.Repository
	ctx    context.Context
}

func NewQuestionUsecase(repo question.QuestionRepository, qcRepo question_categories.Repository, ctx context.Context) *Usecase {
	return &Usecase{
		repo:   repo,
		ctx:    ctx,
		qcRepo: qcRepo,
	}
}

func (uc *Usecase) GetAll(c *gin.Context) (*question.QuestionResponse, *tools.Pagination, error) {
	pagination, err := tools.Paginate(c)
	if err != nil {
		return nil, nil, err
	}

	var filter question.Filter
	if err := c.ShouldBindQuery(&filter); err != nil {
		return nil, nil, err
	}

	categoryQuestion, err := uc.qcRepo.GetDetail(uc.ctx, filter.Language, filter.QuestionCategoryOrder)
	if err != nil {
		return nil, nil, err
	}

	filter.CategoryId = categoryQuestion.ID

	result, pagination, err := uc.repo.GetAll(uc.ctx, pagination, &filter)
	if err != nil {
		return nil, nil, err
	}

	finalRes := &question.QuestionResponse{
		Questions:    result,
		CategoryName: categoryQuestion.CategoryName,
	}

	return finalRes, pagination, nil
}

func (uc *Usecase) GetDetail(c *gin.Context) (*question.Question, error) {
	language := c.Param("language")
	order := c.Param("order")

	fmt.Println("tessss", language, order)
	orderInt, err := strconv.Atoi(order)

	result, err := uc.repo.GetDetail(uc.ctx, orderInt, language)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uc *Usecase) Create(c *gin.Context) error {
	var createQuestion *question.Form
	err := c.ShouldBindJSON(&createQuestion)
	if err != nil {
		return err
	}

	err = uc.repo.Create(uc.ctx, createQuestion)
	if err != nil {
		return err
	}

	return nil
}

// func (uc *Usecase) GetAllUser(c *gin.Context) ([]*user.UserResponse, *tools.Pagination, error) {
// 	pagination, err := tools.Paginate(c)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	filter := new(user.FilterUser)
// 	if err := c.ShouldBindQuery(filter); err != nil {
// 		return nil, nil, err
// 	}

// 	result, pagination, err := uc.repo.GetAllUser(uc.ctx, pagination, filter)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	return result, pagination, nil
// }

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
