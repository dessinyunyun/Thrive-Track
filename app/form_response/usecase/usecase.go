package usecase

import (
	"context"
	"go-gin/app/form_response"
	"go-gin/app/tools"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Usecase struct {
	repo form_response.FormResponseRepository
	ctx  context.Context
}

func (uc *Usecase) GetAll(c *gin.Context, filter *form_response.Filter) ([]*form_response.Response, *tools.Pagination, error) {
	pagination, err := tools.Paginate(c)
	if err != nil {
		return nil, nil, err
	}

	if err := c.ShouldBindQuery(filter); err != nil {
		return nil, nil, err
	}

	result, pagination, err := uc.repo.GetAll(uc.ctx, pagination, filter)
	if err != nil {
		return nil, nil, err
	}

	return result, pagination, nil
}

func NewFormResponseUsecase(repo form_response.FormResponseRepository, ctx context.Context) *Usecase {
	return &Usecase{
		repo: repo,
		ctx:  ctx,
	}
}

func (uc *Usecase) GetDetail(c *gin.Context) (*form_response.Response, error) {
	formResponseId := c.Param("id")

	formResponseIdInt, _ := strconv.Atoi(formResponseId)

	result, err := uc.repo.GetDetail(uc.ctx, formResponseIdInt)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uc *Usecase) Create(c *gin.Context) error {
	var create *form_response.Form
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
