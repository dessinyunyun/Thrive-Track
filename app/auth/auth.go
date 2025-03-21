package auth

import (
	"context"
	"go-gin/database/ent"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	googleUUID "github.com/google/uuid"
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

type LoginResponseUser struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Authenticate struct {
	User  *ent.User `json:"user"`
	Token Token     `json:"token"`
}

type Token struct {
	AccessToken         string    `json:"access_token"`
	RefreshToken        string    `json:"refresh_token"`
	AccessTokenExpired  time.Time `json:"-"`
	RefreshTokenExpired time.Time `json:"-"`
}

type RefreshTokenForm struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type ActivatedTokenForm struct {
	Token string `json:"token"`
}

type JwtWrapper struct {
	SecretKey         string
	Issuer            string
	ExpirationMinutes int64
	ExpirationHours   int64
}

type CustomClaims struct {
	User               *ent.User `json:"user"`
	jwt.StandardClaims `json:"standard_claims"`
}

type AuthUsecase interface {
	Register(c *gin.Context) error
	Login(c *gin.Context) (*Authenticate, error)
	RefreshToken(c *gin.Context) (*Authenticate, error)
	ActivateUser(c *gin.Context) error
	GetDetailAT(userID googleUUID.UUID) (*ent.Activation_token, error)
	GetDetailToken(userID googleUUID.UUID) (*ent.Token, error)
}

type AuthRepository interface {
	CreateAT(ctx context.Context, token, userID string) error
	UsedAT(ctx context.Context, userID googleUUID.UUID) error
	GetDetailAT(ctx context.Context, userID googleUUID.UUID) (*ent.Activation_token, error)

	CreateToken(ctx context.Context, userID googleUUID.UUID, token Token) error
	GetDetailToken(ctx context.Context, userID googleUUID.UUID) (*ent.Token, error)
	RevokedRefreshToken(ctx context.Context, userID googleUUID.UUID) error
}
