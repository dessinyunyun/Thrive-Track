package auth

import (
	"errors"
	"go-gin/app/tools"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

var (
	ErrIdentityNotFound    = errors.New("Email or username not found. Please register.")
	ErrWrongPassword       = errors.New("Authentication failed: wrong password.")
	ErrEmailExists         = errors.New("Email already exists.")
	ErrUsernameExists      = errors.New("Username already exists.")
	ErrEmailUsernameExists = errors.New("Email and username already exist.")
	ErrUserNotFound        = errors.New("User not found.")
	ErrInvalidToken        = errors.New("Invalid token.")
	ErrATalreadyUsed       = errors.New("The activation token has already been used.")
	ErrATnotFound          = errors.New("User activation token not found.")
	ErrAccountNotActivated = errors.New("Account has not been activated. Please check your email for the activation link to complete the process.")
)

func ErrorHandler(c *gin.Context, log *logrus.Entry, err error) {
	switch {
	case isJWTValidationError(err):
		handleJWTValidationError(c, err)
	case isValidationError(err):
		handleValidationError(c, err)
	case isAuthError(err):
		handleAuthError(c, err)
	default:
		handleInternalServerError(c, log, err)
	}
}

func isJWTValidationError(err error) bool {
	var ve *jwt.ValidationError
	return errors.As(err, &ve)
}

func handleJWTValidationError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, tools.Response{
		Status:  "error",
		Message: err.Error(),
	})
}

func isValidationError(err error) bool {
	_, ok := err.(validator.ValidationErrors)
	return ok
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

func isAuthError(err error) bool {
	return errors.Is(err, ErrEmailUsernameExists) ||
		errors.Is(err, ErrEmailExists) ||
		errors.Is(err, ErrUsernameExists) ||
		errors.Is(err, ErrWrongPassword) ||
		errors.Is(err, ErrInvalidToken) ||
		errors.Is(err, ErrATalreadyUsed) ||
		errors.Is(err, ErrAccountNotActivated) ||
		errors.Is(err, ErrIdentityNotFound) ||
		errors.Is(err, ErrUserNotFound) ||
		errors.Is(err, ErrATnotFound)
}

func handleAuthError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, ErrEmailUsernameExists) || errors.Is(err, ErrEmailExists) || errors.Is(err, ErrUsernameExists):
		c.AbortWithStatusJSON(http.StatusBadRequest, tools.Response{
			Status:  "error",
			Message: err.Error(),
		})
	case errors.Is(err, ErrWrongPassword) || errors.Is(err, ErrInvalidToken) || errors.Is(err, ErrATalreadyUsed):
		c.AbortWithStatusJSON(http.StatusUnauthorized, tools.Response{
			Status:  "error",
			Message: err.Error(),
		})
	case errors.Is(err, ErrAccountNotActivated):
		c.AbortWithStatusJSON(http.StatusForbidden, tools.Response{
			Status:  "error",
			Message: err.Error(),
		})
	case errors.Is(err, ErrIdentityNotFound) || errors.Is(err, ErrUserNotFound) || errors.Is(err, ErrATnotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, tools.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
}

func handleInternalServerError(c *gin.Context, log *logrus.Entry, err error) {
	log.Errorf("Auth handlers err: %v", err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, tools.Response{
		Status:  "error",
		Message: err.Error(),
	})
}
