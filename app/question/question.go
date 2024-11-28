package question

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ID          int    `json:"id"`
	Text        string `json:"text"`
	Language    string `json:"language"`
	Order       int    `json:"order"`
	Description string `json:"description"`
}

type Form struct {
	ID          int    `json:"-"`
	Text        string `json:"text"`
	Language    string `json:"language"`
	Order       int    `json:"order"`
	Description string `json:"description"`
}

type Filter struct {
	ID       int    `json:"id" form:"id[]"`
	Order    int    `json:"order"`
	Language string `json:"language"`
}

type QuestionUsecase interface {
	GetDetail(c *gin.Context) (*Response, error)
	Create(c *gin.Context) error
	// GetAllUser(c *gin.Context) ([]*Response, *tools.Pagination, error)
	// UpdateUser(c *gin.Context) error
	// DeleteUser(c *gin.Context) error
}

type QuestionRepository interface {
	GetDetail(ctx context.Context, order int, language string) (*Response, error)
	Create(ctx context.Context, form *Form) error
	// GetAllUser(ctx context.Context, pagination *tools.Pagination, filter *Filter) ([]*UserResponse, *tools.Pagination, error)
	// CheckUserIdentifier(ctx context.Context, identifier string) (*UserResponseSensitiveCase, error)
	// UpdateUser(ctx context.Context, form *UserForm) error
	// DeleteUser(ctx context.Context, ID uuid.UUID) error
}
