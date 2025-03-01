package usecase

import (
	"context"
	"errors"
	"go-gin/app/auth"
	"go-gin/app/middleware"
	"go-gin/app/user"
	"go-gin/database/ent"
	"os"
	"strings"

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

	user, err := uc.repo.CheckEmailAndUsernameExist(uc.ctx, &createUser.Email, &createUser.Username)
	if err != nil {
		return err
	}
	if user != nil {
		if *user.Email == createUser.Email && user.Username == createUser.Username {
			return auth.ErrEmailandUsernameAlreadyExist
		}
		if *user.Email == createUser.Email {
			return auth.ErrEmailAlreadyExist
		}
		if user.Username == createUser.Username {
			return auth.ErrUsernameAlreadyExist
		}
	}

	err = uc.repo.CreateUser(uc.ctx, createUser)
	if err != nil {
		if ent.IsConstraintError(err) {
			switch {
			case strings.Contains(err.Error(), "email"):
				return auth.ErrEmailAlreadyExist
			case strings.Contains(err.Error(), "username"):
				return auth.ErrUsernameAlreadyExist
			}

		}
		return err
	}

	return nil
}

func (uc *AuthUsecase) Login(c *gin.Context) (*auth.LoginResponse, error) {
	var loginForm auth.LoginForm
	err := c.ShouldBindJSON(&loginForm)
	if err != nil {
		return nil, err
	}

	user, err := uc.repo.CheckEmailAndUsernameExist(uc.ctx, &loginForm.Email, &loginForm.Username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, auth.ErrIdentityNotFound
	}

	err = CheckComparePass(loginForm.Password, user.Password)
	jwtWrapper := middleware.JwtWrapper{
		SecretKey:         os.Getenv("PREFIX_API"),
		Issuer:            "MHService",
		ExpirationMinutes: 1,
		ExpirationHours:   12,
	}

	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, auth.ErrWrongPassword
		} else {
			return nil, err
		}
	}

	signedToken, err := jwtWrapper.GenerateToken(*user.Email)
	if err != nil {
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

	// result := &auth.LoginResponse{
	// 	Token:        "",
	// 	RefreshToken: "",
	// 	User: auth.LoginResponseUser{ // Membuat objek LoginResponseUser
	// 		ID:       "", // Jika user.ID adalah UUID, pastikan menggunakan .String() untuk mengonversi menjadi string
	// 		Name:     "",
	// 		Username: "",
	// 		Email:    "",
	// 	},
	// }

	return result, nil
}
