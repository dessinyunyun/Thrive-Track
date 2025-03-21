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

// func (h *Handler) GetAll(c *gin.Context) {
// 	var filter history_answer.Filter

// 	if err := c.ShouldBindQuery(&filter); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"message": "Invalid query parameters: " + err.Error(),
// 		})
// 		return
// 	}

// 	result, pagination, err := h.uc.GetAll(c, &filter)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data":    result,
// 		"status":  "success",
// 		"message": "Get All History Answer",
// 		"meta":    pagination,
// 	})
// }

// func (h *Handler) GetDetail(c *gin.Context) {
// 	result, err := h.uc.GetDetail(c)
// 	if err != nil {
// 		h.log.Errorf("get detail History Answer handlers: %v", err)
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"message": err,
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, tools.Response{
// 		Data:    result,
// 		Status:  "success",
// 		Message: "Get Detail History Answer",
// 	})
// }

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
