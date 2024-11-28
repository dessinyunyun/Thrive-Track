package usecase

import (
	"context"
	"go-gin/app/auth"
	"go-gin/app/middleware"
	"go-gin/app/user"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	repo user.UserRepository
	ctx  context.Context
}

func NewAuthUsecase(repo user.UserRepository, ctx context.Context) *AuthUsecase {
	return &AuthUsecase{
		repo: repo,
		ctx:  ctx,
	}
}

func (uc *AuthUsecase) Register(c *gin.Context) error {
	var createUser *user.UserForm
	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		return err
	}
	password := HashPassword(createUser.Password)
	createUser.Password = password

	err = uc.repo.CreateUser(uc.ctx, createUser)
	if err != nil {
		return err
	}

	return nil
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func CheckComparePass(clientProvidePassword string, userPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(clientProvidePassword))
	if err != nil {
		return err
	}
	return nil
}

func (uc *AuthUsecase) Login(c *gin.Context) (*auth.LoginResponse, error) {
	var loginForm *auth.LoginForm
	err := c.ShouldBindJSON(&loginForm)
	if err != nil {
		return nil, err
	}

	user, err := uc.repo.CheckUserIdentifier(uc.ctx, loginForm.Identifier)
	if err != nil {
		return nil, err
	}

	CheckComparePass(loginForm.Password, user.Password)
	jwtWrapper := middleware.JwtWrapper{
		SecretKey:         os.Getenv("PREFIX_API"),
		Issuer:            "MHService",
		ExpirationMinutes: 1,
		ExpirationHours:   12,
	}

	signedToken, err := jwtWrapper.GenerateToken(*user.Email)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": "error signing token",
		})
		c.Abort()
		return nil, err
	}
	result := &auth.LoginResponse{
		Token:        signedToken,
		RefreshToken: signedToken,
		User: auth.LoginResponseUser{ // Membuat objek LoginResponseUser
			ID:       user.ID.String(), // Jika user.ID adalah UUID, pastikan menggunakan .String() untuk mengonversi menjadi string
			Name:     user.Name,
			Username: user.Username,
			Email:    *user.Email,
		},
	}

	return result, nil
}
