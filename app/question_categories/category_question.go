package question_categories

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name" form:"category_name"`
	Order        int    `json:"order"`
	Language     string `json:"language"`
}

type Form struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
	Order        int    `json:"order"`
	Language     string `json:"language"`
}

type Filter struct {
	ID             int `json:"id" form:"id[]"`
	UserId         int `json:"user_id" form:"user_id"`
	FormResponseId int `json:"form_response_id" form:"form_response_id"`
	QuestionId     int `json:"question_id" form:"question_id"`
	// Language string `json:"language"`
}

type Usecase interface {
	// GetDetail(c *gin.Context) (*Response, error)
	Create(c *gin.Context) error
	// GetAll(c *gin.Context, filter *Filter) ([]*Response, *tools.Pagination, error) // UpdateUser(c *gin.Context) error
	// DeleteUser(c *gin.Context) error
}

type Repository interface {
	// GetDetail(ctx context.Context, id int) (*Response, error)
	Create(ctx context.Context, form *Form) error
	GetDetail(ctx context.Context, language string, qcOrder int) (*Response, error)
}
