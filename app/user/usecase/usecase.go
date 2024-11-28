package usecase

import (
	"context"
	"go-gin/app/tools"
	"go-gin/app/user"

	"github.com/gin-gonic/gin"
	googleUUID "github.com/google/uuid"
)

type UserUsecase struct {
	repo user.UserRepository
	ctx  context.Context
}

func NewUserUsecase(repo user.UserRepository, ctx context.Context) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		ctx:  ctx,
	}
}

func (uc *UserUsecase) GetAllUser(c *gin.Context) ([]*user.UserResponse, *tools.Pagination, error) {
	pagination, err := tools.Paginate(c)
	if err != nil {
		return nil, nil, err
	}

	filter := new(user.FilterUser)
	if err := c.ShouldBindQuery(filter); err != nil {
		return nil, nil, err
	}

	result, pagination, err := uc.repo.GetAllUser(uc.ctx, pagination, filter)
	if err != nil {
		return nil, nil, err
	}

	return result, pagination, nil
}

func (uc *UserUsecase) GetDetailUser(c *gin.Context) (*user.UserResponse, error) {
	uuid, err := googleUUID.Parse(c.Param("id"))
	if err != nil {
		return nil, err
	}

	result, err := uc.repo.GetDetailUser(uc.ctx, uuid)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uc *UserUsecase) CreateUser(c *gin.Context) error {
	var createUser *user.UserForm
	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		return err
	}

	err = uc.repo.CreateUser(uc.ctx, createUser)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUsecase) UpdateUser(c *gin.Context) error {
	var updateUser *user.UserForm
	err := c.ShouldBindJSON(&updateUser)
	if err != nil {
		return err
	}

	ID, err := googleUUID.Parse(c.Param("id"))
	if err != nil {
		return err
	}

	updateUser.ID = ID

	err = uc.repo.UpdateUser(uc.ctx, updateUser)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUsecase) DeleteUser(c *gin.Context) error {
	ID, err := googleUUID.Parse(c.Param("id"))
	if err != nil {
		return err
	}

	err = uc.repo.DeleteUser(uc.ctx, ID)
	if err != nil {
		return err
	}
	return nil
}
