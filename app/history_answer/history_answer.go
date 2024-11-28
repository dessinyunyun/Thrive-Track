package history_answer

import (
	"context"
	"go-gin/app/tools"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ID             int `json:"id"`
	FormResponseId int `json:"form_response_id" form:"form_response_id"`
	QuestionId     int `json:"question_id"`
	Answer         int `json:"answer"`
}

type Form struct {
	ID             int `json:"-"`
	FormResponseId int `json:"form_response_id" form:"form_response_id"`
	QuestionId     int `json:"question_id"`
	Answer         int `json:"answer"`
}

type Filter struct {
	ID             int `json:"id" form:"id[]"`
	UserId         int `json:"user_id" form:"user_id"`
	FormResponseId int `json:"form_response_id" form:"form_response_id"`
	QuestionId     int `json:"question_id" form:"question_id"`
	// Language string `json:"language"`
}

type HistoryAnswerUsecase interface {
	GetDetail(c *gin.Context) (*Response, error)
	Create(c *gin.Context) error
	GetAll(c *gin.Context, filter *Filter) ([]*Response, *tools.Pagination, error) // UpdateUser(c *gin.Context) error
	// DeleteUser(c *gin.Context) error
}

type HistoryAnswerRepository interface {
	GetDetail(ctx context.Context, id int) (*Response, error)
	Create(ctx context.Context, form *Form) error
	GetAll(ctx context.Context, pagination *tools.Pagination, filter *Filter) ([]*Response, *tools.Pagination, error)
	// CheckUserIdentifier(ctx context.Context, identifier string) (*UserResponseSensitiveCase, error)
	// UpdateUser(ctx context.Context, form *UserForm) error
	// DeleteUser(ctx context.Context, ID uuid.UUID) error
}
