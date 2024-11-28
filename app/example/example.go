package example

import (
	"context"
	"go-gin/app/tools"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ExampleResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Email    *string   `json:"email"`
}

type ExampleForm struct {
	ID       uuid.UUID `json:"-"`
	Name     string    `json:"name" binding:"required"`
	Username string    `json:"username" binding:"required"`
	Email    string    `json:"email" binding:"required"`
}

type FilterExample struct {
	ID       uuid.UUID `json:"id" form:"id[]"`
	Name     string    `json:"name" form:"name"`
	Username string    `json:"username" form:"username"`
	Email    string    `json:"email" form:"email"`
}

type ExampleUsecase interface {
	GetAllExample(c *gin.Context) ([]*ExampleResponse, *tools.Pagination, error)
	GetDetailExample(c *gin.Context) (*ExampleResponse, error)
	CreateExample(c *gin.Context) error
	UpdateExample(c *gin.Context) error
	DeleteExample(c *gin.Context) error
}

type ExampleRepository interface {
	GetAllExample(ctx context.Context, pagination *tools.Pagination, filter *FilterExample) ([]*ExampleResponse, *tools.Pagination, error)
	GetDetailExample(ctx context.Context, ID uuid.UUID) (*ExampleResponse, error)
	CreateExample(ctx context.Context, form *ExampleForm) error
	UpdateExample(ctx context.Context, form *ExampleForm) error
	DeleteExample(ctx context.Context, ID uuid.UUID) error
}
