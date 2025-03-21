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
