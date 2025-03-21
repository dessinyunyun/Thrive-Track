package handler

import (
	"go-gin/app/question"
	"go-gin/app/tools"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	uc  question.QuestionUsecase
	log *logrus.Entry
}

func QuestionRoute(uc question.QuestionUsecase, r *gin.RouterGroup, log *logrus.Entry) {
	h := Handler{
		uc:  uc,
		log: log,
	}

	v2 := r.Group("question")

	v2.GET("/", h.GetAll)
	v2.GET("/:language/:order", h.GetDetail)
	v2.POST("", h.Create)

}

// @Tags Question
// @Summary Get All Question
// @Description Get All Question
// @Router /question [get]
// @Accept json
// @Produce json
// @param id query array false "ID in Array"
// @param order query string false "Filter by order"
// @param language query string false "Filter by language"
// @param category_id query string false "Filter by category_id"
// @param limit query integer false "Limit of pagination"
// @param page query integer false "Page of pagination"
func (h *Handler) GetAll(c *gin.Context) {
	result, pagination, err := h.uc.GetAll(c)
	if err != nil {
		if err != nil {
			question.ErrorHandler(c, h.log, err)
			return
		}
		return
	}

	c.JSON(http.StatusOK, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get All Question",
		Meta:    pagination,
	})
}

// @Tags Question
// @Summary Get Detail Question
// @Description Get Detail Question by language and order. lang(en/id)
// @Router /question/{language}/{order} [get]
// @Accept json
// @Produce json
// @param language path string true "language (en or id)"
// @param order path string true "order"
func (h *Handler) GetDetail(c *gin.Context) {
	result, err := h.uc.GetDetail(c)
	if err != nil {
		if err != nil {
			question.ErrorHandler(c, h.log, err)
			return
		}
		return
	}

	c.JSON(http.StatusOK, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get Detail Question",
	})
}

// @Tags Question
// @Summary Create Question
// @Description Create Question
// @Router /Question [post]
// @Accept json
// @Produce json
// @Param request body question.Form true "Payload Body for Create [RAW]"
func (h *Handler) Create(c *gin.Context) {
	err := h.uc.Create(c)
	if err != nil {
		if err != nil {
			question.ErrorHandler(c, h.log, err)
			return
		}
		return
	}

	c.JSON(http.StatusCreated, tools.Response{
		Status:  "success",
		Message: "Create Question",
	})
}
