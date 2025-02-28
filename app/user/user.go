package user

import (
	"context"
	"go-gin/app/tools"

	"github.com/gin-gonic/gin"

	googleUUID "github.com/google/uuid"
)

type UserResponse struct {
	ID       googleUUID.UUID `json:"id"`
	Name     string          `json:"name"`
	Username string          `json:"username"`
	Email    *string         `json:"email"`
}

type UserResponseSensitiveCase struct {
	ID       googleUUID.UUID `json:"id"`
	Name     string          `json:"name"`
	Username string          `json:"username"`
	Email    *string         `json:"email"`
	Password string          `json:"password"`
}

type UserForm struct {
	ID       googleUUID.UUID `json:"-"`
	Name     string          `json:"name" binding:"required"`
	Username string          `json:"username" binding:"required"`
	Email    string          `json:"email" binding:"required"`
	Password string          `json:"password" binding:"required"`
}

type FilterUser struct {
	ID       googleUUID.UUID `json:"id" form:"id[]"`
	Name     string          `json:"name" form:"name"`
	Username string          `json:"username" form:"username"`
	Email    string          `json:"email" form:"email"`
}

type IdentifierForm struct {
	Identifier string `form:"identifier" binding:"required"`
}

type IdentifierResponse struct {
	Identifier string `json:"identifier" binding:"required"`
}

type UserUsecase interface {
	GetAllUser(c *gin.Context) ([]*UserResponse, *tools.Pagination, error)
	GetDetailUser(c *gin.Context) (*UserResponse, error)
	CreateUser(c *gin.Context) error
	UpdateUser(c *gin.Context) error
	DeleteUser(c *gin.Context) error
	CheckUserIdentifier(c *gin.Context) (*IdentifierResponse, error)
}

type UserRepository interface {
	GetAllUser(ctx context.Context, pagination *tools.Pagination, filter *FilterUser) ([]*UserResponse, *tools.Pagination, error)
	GetDetailUser(ctx context.Context, ID googleUUID.UUID) (*UserResponse, error)
	CheckUserIdentifier(ctx context.Context, identifier string) (*UserResponseSensitiveCase, error)
	CreateUser(ctx context.Context, form *UserForm) error
	UpdateUser(ctx context.Context, form *UserForm) error
	DeleteUser(ctx context.Context, ID googleUUID.UUID) error
}
