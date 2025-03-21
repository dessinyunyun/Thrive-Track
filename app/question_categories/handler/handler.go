package handler

import (
	"go-gin/app/middleware"
	"go-gin/app/question_categories"
	"go-gin/app/tools"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	uc  question_categories.Usecase
	log *logrus.Entry
}

func CategoryQuestionRoute(uc question_categories.Usecase, r *gin.RouterGroup, log *logrus.Entry) {
	h := Handler{
		uc:  uc,
		log: log,
	}

	v2 := r.Group("category-question", middleware.AuthMiddleware(log))

	// v2.GET("/", h.GetAll)
	// v2.GET("/:id", h.GetDetail)
	v2.POST("", h.Create)

}

// @Tags Question Category
// @Summary Create Question Category
// @Description Create Question Category
// @Router /category-question [post]
// @Accept json
// @Produce json
// @Param request body question_categories.Form true "Payload Body for Create [RAW]"
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
