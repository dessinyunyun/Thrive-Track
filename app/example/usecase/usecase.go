package usecase

import (
	"context"
	"go-gin/app/example"
	"go-gin/app/tools"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ExampleUsecase struct {
	repo example.ExampleRepository
	ctx  context.Context
}

func NewExampleUsecase(repo example.ExampleRepository, ctx context.Context) *ExampleUsecase {
	return &ExampleUsecase{
		repo: repo,
		ctx:  ctx,
	}
}

func (uc *ExampleUsecase) GetAllExample(c *gin.Context) ([]*example.ExampleResponse, *tools.Pagination, error) {
	pagination, err := tools.Paginate(c)
	if err != nil {
		return nil, nil, err
	}

	filter := new(example.FilterExample)
	if err := c.ShouldBindQuery(filter); err != nil {
		return nil, nil, err
	}

	result, pagination, err := uc.repo.GetAllExample(uc.ctx, pagination, filter)
	if err != nil {
		return nil, nil, err
	}

	return result, pagination, nil
}

func (uc *ExampleUsecase) GetDetailExample(c *gin.Context) (*example.ExampleResponse, error) {
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return nil, err
	}

	result, err := uc.repo.GetDetailExample(uc.ctx, uuid)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uc *ExampleUsecase) CreateExample(c *gin.Context) error {
	var createExample *example.ExampleForm
	err := c.ShouldBindJSON(&createExample)
	if err != nil {
		return err
	}

	err = uc.repo.CreateExample(uc.ctx, createExample)
	if err != nil {
		return err
	}

	return nil
}

func (uc *ExampleUsecase) UpdateExample(c *gin.Context) error {
	var updateExample *example.ExampleForm
	err := c.ShouldBindJSON(&updateExample)
	if err != nil {
		return err
	}

	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return err
	}

	updateExample.ID = ID

	err = uc.repo.UpdateExample(uc.ctx, updateExample)
	if err != nil {
		return err
	}

	return nil
}

func (uc *ExampleUsecase) DeleteExample(c *gin.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return err
	}

	err = uc.repo.DeleteExample(uc.ctx, ID)
	if err != nil {
		return err
	}
	return nil
}
