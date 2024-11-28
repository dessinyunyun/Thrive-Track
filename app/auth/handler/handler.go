package handler

import (
	"go-gin/app/auth"
	"go-gin/app/tools"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	uc  auth.AuthUsecase
	log *logrus.Entry
}

func AuthRoute(uc auth.AuthUsecase, r *gin.RouterGroup, log *logrus.Entry) {
	h := AuthHandler{
		uc:  uc,
		log: log,
	}

	v2 := r.Group("auth")

	v2.POST("/register", h.Register)
	v2.POST("/login", h.Login)
	// v2.PUT("/:id", h.UpdateUser)
	// v2.DELETE("/:id", h.DeleteUser)
}

// @Tags Example
// @Summary Create Example
// @Description Create Example
// @Router /example [post]
// @Accept json
// @Produce json
// @Param request body example.ExampleForm true "Payload Body for Create [RAW]"
func (h *AuthHandler) Register(c *gin.Context) {
	err := h.uc.Register(c)
	if err != nil {
		h.log.Errorf("create User handlers: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, tools.Response{
		Status:  "success",
		Message: "success register",
	})
}

func (h *AuthHandler) Login(c *gin.Context) {

	result, err := h.uc.Login(c)
	if err != nil {
		h.log.Errorf("get detail User handlers: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusCreated, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "success Login",
	})
}
