package auth

import (
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
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
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
