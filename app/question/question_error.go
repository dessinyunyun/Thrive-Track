package question

import (
	"errors"
	"go-gin/app/tools"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

var (
	ErrQuestionCategoryNotFound = errors.New("question category not found")
)

func ErrorHandler(c *gin.Context, log *logrus.Entry, err error) {
	switch {
	case isValidationError(err):
		handleValidationError(c, err)
	case isQuestionError(err):
		handleQuestionError(c, err)
	default:
		handleInternalServerError(c, log, err)
	}
}

func isValidationError(err error) bool {
	_, ok := err.(validator.ValidationErrors)
	return ok
}

func isQuestionError(err error) bool {
	return errors.Is(err, ErrQuestionCategoryNotFound)
}

func handleValidationError(c *gin.Context, err error) {
	validationErrors := err.(validator.ValidationErrors)
	validationErr := tools.ValidationErrors(validationErrors)
	c.AbortWithStatusJSON(http.StatusBadRequest, tools.Response{
		Status:  "error",
		Message: err.Error(),
		Data:    validationErr,
	})
}

func handleQuestionError(c *gin.Context, err error) {
	switch {
	// case errors.Is(err, ErrEmailUsernameExists) || errors.Is(err, ErrEmailExists) || errors.Is(err, ErrUsernameExists):
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, tools.Response{
	// 		Status:  "error",
	// 		Message: err.Error(),
	// 	})
	// case errors.Is(err, ErrWrongPassword) || errors.Is(err, ErrInvalidToken) || errors.Is(err, ErrATalreadyUsed):
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, tools.Response{
	// 		Status:  "error",
	// 		Message: err.Error(),
	// 	})
	// case errors.Is(err, ErrAccountNotActivated):
	// 	c.AbortWithStatusJSON(http.StatusForbidden, tools.Response{
	// 		Status:  "error",
	// 		Message: err.Error(),
	// 	})
	case errors.Is(err, ErrQuestionCategoryNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, tools.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
}

func handleInternalServerError(c *gin.Context, log *logrus.Entry, err error) {
	log.Errorf("question handlers err: %v", err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, tools.Response{
		Status:  "error",
		Message: err.Error(),
	})
}
