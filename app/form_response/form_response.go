package form_response

import (
	"context"
	"go-gin/app/tools"

	"github.com/gin-gonic/gin"
	googleUUID "github.com/google/uuid"
)

type Response struct {
	ID              int             `json:"id"`
	TotalScore      uint8           `json:"total_score"`
	DepressionLevel string          `json:"depression_level"`
	UserId          googleUUID.UUID `json:"user_id" form:"user_id"`
}

type Form struct {
	ID              int             `json:"-"`
	TotalScore      uint8           `json:"total_score"`
	DepressionLevel string          `json:"depression_level"`
	UserId          googleUUID.UUID `json:"user_id" form:"user_id"`
	Username        string          `json:"username"`
}

type Filter struct {
	ID       int    `json:"id" form:"id[]"`
	UserId   string `json:"user_id" form:"user_id"`
	Username string `json:"username"`
	// Language string `json:"language"`
}

type FormResponseUsecase interface {
	GetDetail(c *gin.Context) (*Response, error)
	Create(c *gin.Context) error
	GetAll(c *gin.Context, filter *Filter) ([]*Response, *tools.Pagination, error) // UpdateUser(c *gin.Context) error
	// DeleteUser(c *gin.Context) error
}

type FormResponseRepository interface {
	GetDetail(ctx context.Context, id int) (*Response, error)
	Create(ctx context.Context, form *Form) error
	GetAll(ctx context.Context, pagination *tools.Pagination, filter *Filter) ([]*Response, *tools.Pagination, error)
	// CheckUserIdentifier(ctx context.Context, identifier string) (*UserResponseSensitiveCase, error)
	// UpdateUser(ctx context.Context, form *UserForm) error
	// DeleteUser(ctx context.Context, ID uuid.UUID) error
}
