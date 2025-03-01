package handler

import (
	"errors"
	"go-gin/app/auth"
	"go-gin/app/tools"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

// @Tags Auth
// @Summary Register
// @Description Register
// @Router /auth/register [post]
// @Accept json
// @Produce json
// @Param request body auth.RegisterForm true "Payload Body for Register [RAW]"
func (h *AuthHandler) Register(c *gin.Context) {
	err := h.uc.Register(c)
	if err != nil {

		if errors.Is(err, auth.ErrEmailandUsernameAlreadyExist) || errors.Is(err, auth.ErrEmailAlreadyExist) || errors.Is(err, auth.ErrUsernameAlreadyExist) {
			c.AbortWithStatusJSON(http.StatusBadRequest, tools.Response{
				Status:  "error",
				Message: err.Error(),
			})
			return
		}

		h.log.Errorf("create User handlers: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, tools.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, tools.Response{
		Status:  "success",
		Message: "success register",
	})
}

// @Tags Auth
// @Summary Login
// @Description Login
// @Router /auth/login [post]
// @Accept json
// @Produce json
// @Param request body auth.LoginForm true "Payload Body for Login [RAW]"
func (h *AuthHandler) Login(c *gin.Context) {
	result, err := h.uc.Login(c)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			validationErr := tools.ValidationErrors(validationErrors)
			c.AbortWithStatusJSON(http.StatusBadRequest, tools.Response{
				Status:  "error",
				Message: "Validation error",
				Data:    validationErr,
			})
			return
		} else if errors.Is(err, auth.ErrIdentityNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, tools.Response{
				Status:  "error",
				Message: err.Error(),
			})
			return
		} else if errors.Is(err, auth.ErrWrongPassword) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, tools.Response{
				Status:  "error",
				Message: err.Error(),
			})
			return
		}

		h.log.Errorf("Login handlers err: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, tools.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, tools.Response{
		Status:  "success",
		Message: "success login",
		Data:    result,
	})
}
