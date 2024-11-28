package handler

import (
	"fmt"
	"go-gin/app/history_answer"
	"go-gin/app/tools"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	uc  history_answer.HistoryAnswerUsecase
	log *logrus.Entry
}

func HistoryAnswerRoute(uc history_answer.HistoryAnswerUsecase, r *gin.RouterGroup, log *logrus.Entry) {
	h := Handler{
		uc:  uc,
		log: log,
	}

	v2 := r.Group("history-answer")

	v2.GET("/", h.GetAll)
	v2.GET("/:id", h.GetDetail)
	v2.POST("", h.Create)

}

func (h *Handler) GetAll(c *gin.Context) {
	var filter history_answer.Filter

	if err := c.ShouldBindQuery(&filter); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid query parameters: " + err.Error(),
		})
		return
	}
	fmt.Println("tse", filter)

	result, pagination, err := h.uc.GetAll(c, &filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    result,
		"status":  "success",
		"message": "Get All History Answer",
		"meta":    pagination,
	})
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
		h.log.Errorf("get detail History Answer handlers: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get Detail History Answer",
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
		h.log.Errorf("create History Answer handlers: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, tools.Response{
		Status:  "success",
		Message: "Create History Answer",
	})
}
