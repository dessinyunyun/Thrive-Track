package question

import (
	"context"
	"go-gin/app/tools"

	"github.com/gin-gonic/gin"
)

type Question struct {
	ID          int    `json:"id,omitempty"`
	Text        string `json:"text,omitempty"`
	Language    string `json:"language,omitempty"`
	Description string `json:"description,omitempty"`
	Example     string `json:"example,omitempty"`
	Order       int    `json:"order,omitempty"`
	CategoryID  int    `json:"category_id,omitempty"`
}

type QuestionResponse struct {
	CategoryName string      `json:"category_name,omitempty"`
	Questions    []*Question `json:"questions,omitempty"`
}

type Form struct {
	ID          int    `json:"-"`
	Text        string `json:"text"`
	Language    string `json:"language"`
	Order       int    `json:"order"`
	Description string `json:"description"`
	Example     string `json:"example"`
	Category    int    `json:"category_id"`
}

type Filter struct {
	ID                    int    `json:"id" form:"id"`
	Order                 int    `json:"order" form:"order"`
	Language              string `json:"language" form:"language"`
	CategoryId            int    `json:"category_id" form:"category_id"`
	QuestionCategoryOrder int    `json:"question_category_order" form:"question_category_order"`
}

type QuestionUsecase interface {
	GetDetail(c *gin.Context) (*Question, error)
	Create(c *gin.Context) error
	GetAll(c *gin.Context) (*QuestionResponse, *tools.Pagination, error)
}

type QuestionRepository interface {
	GetDetail(ctx context.Context, order int, language string) (*Question, error)
	Create(ctx context.Context, form *Form) error
	GetAll(ctx context.Context, pagination *tools.Pagination, filter *Filter) ([]*Question, *tools.Pagination, error)
}
