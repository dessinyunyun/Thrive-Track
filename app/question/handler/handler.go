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

	v2.GET("/", h.GetDetail)
	v2.GET("/:language/:order", h.GetDetail)
	v2.POST("", h.Create)

}

// @Tags Example
// @Summary Get Detail Example
// @Description Get Detail Example by ID
// @Router /example/{id} [get]
// @Accept json
// @Produce json
// @param id path string true "ID"
func (h *Handler) GetDetail(c *gin.Context) {
	result, err := h.uc.GetDetail(c)
	if err != nil {
		h.log.Errorf("get detail Question handlers: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get Detail Question",
	})
}

// @Tags Example
// @Summary Create Example
// @Description Create Example
// @Router /example [post]
// @Accept json
// @Produce json
// @Param request body example.ExampleForm true "Payload Body for Create [RAW]"
func (h *Handler) Create(c *gin.Context) {
	err := h.uc.Create(c)
	if err != nil {
		h.log.Errorf("create Question handlers: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, tools.Response{
		Status:  "success",
		Message: "Create Question",
	})
}