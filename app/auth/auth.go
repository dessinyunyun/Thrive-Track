package auth

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type RegisterForm struct {
	ID       uuid.UUID `json:"-"`
	Name     string    `json:"name" binding:"required"`
	Username string    `json:"username" binding:"required"`
	Email    string    `json:"email" binding:"required"`
	Password string    `json:"password" binding:"required"`
}

type LoginForm struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token        string            `json:"token"`
	RefreshToken string            `json:"refreshtoken"`
	User         LoginResponseUser `json:"user"`
}

type LoginResponseUser struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type AuthUsecase interface {
	Register(c *gin.Context) error
	Login(c *gin.Context) (*LoginResponse, error)
	// UpdateExample(c *gin.Context) error
	// DeleteExample(c *gin.Context) error
}

var (
	ErrIdentityNotFound = errors.New("Email or Username not found. Please register")
	ErrWrongPassword    = errors.New("Authentication failed: wrong password")

	ErrEmailAlreadyExist    = errors.New("Email already exist")
	ErrUsernameAlreadyExist = errors.New("Username already exist")

	ErrEmailandUsernameAlreadyExist = errors.New("Email and Username already exist")
)
