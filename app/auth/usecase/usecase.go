package usecase

import (
	"context"
	"errors"
	"go-gin/app/auth"
	"go-gin/app/mailer"
	"go-gin/app/user"
	"go-gin/database/ent"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	repo          auth.AuthRepository
	userRepo      user.UserRepository
	ctx           context.Context
	mailerUsecase mailer.MailerUsecase
}

func NewAuthUsecase(repo auth.AuthRepository, userRepo user.UserRepository, ctx context.Context, mailerUsecase mailer.MailerUsecase) *AuthUsecase {
	return &AuthUsecase{
		repo:          repo,
		userRepo:      userRepo,
		ctx:           ctx,
		mailerUsecase: mailerUsecase,
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

	user, err := uc.userRepo.CheckEmailAndUsernameExist(uc.ctx, &createUser.Email, &createUser.Username)
	if err != nil {
		return err
	}
	if user != nil {
		if *user.Email == createUser.Email && user.Username == createUser.Username {
			return auth.ErrEmailUsernameExists
		}
		if *user.Email == createUser.Email {
			return auth.ErrEmailExists
		}
		if user.Username == createUser.Username {
			return auth.ErrUsernameExists
		}
	}

	res, err := uc.userRepo.CreateUser(uc.ctx, createUser)
	if err != nil {
		if ent.IsConstraintError(err) {
			switch {
			case strings.Contains(err.Error(), "email"):
				return auth.ErrEmailExists
			case strings.Contains(err.Error(), "username"):
				return auth.ErrEmailExists
			}

		}
		return err
	}

	activationToken, err := uc.GenerateActivationToken(c.Request.Context(), *res.Email, res.ID.String(), 24*time.Hour) // Masa berlaku 24 jam
	if err != nil {
		return err
	}

	mailForm := mailer.EmailForm{
		To: createUser.Email,
		EmailData: mailer.EmailData{
			ActivationURL: os.Getenv("FRONTEND_URL") + *activationToken,
			Username:      createUser.Username,
		},
	}
	go uc.mailerUsecase.ActivatedEmail(mailForm)

	return nil
}

func (uc *AuthUsecase) Login(c *gin.Context) (*auth.Authenticate, error) {
	var loginForm auth.LoginForm
	err := c.ShouldBindJSON(&loginForm)
	if err != nil {
		return nil, err
	}

	user, err := uc.userRepo.CheckEmailAndUsernameExist(uc.ctx, &loginForm.Email, &loginForm.Username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, auth.ErrIdentityNotFound
	}

	if !user.Active {
		return nil, auth.ErrAccountNotActivated
	}

	err = CheckComparePass(loginForm.Password, user.Password)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, auth.ErrWrongPassword
		} else {
			return nil, err
		}
	}

	timeExpiredToken := time.Minute * time.Duration(60)           // 1 hour
	timeExpiredRefreshToken := time.Minute * time.Duration(10080) // 7 days

	signedToken, err := uc.GenerateJWT(*user, timeExpiredToken, timeExpiredRefreshToken)
	if err != nil {
		return nil, err
	}

	result := &auth.Authenticate{
		User: user,
		Token: auth.Token{
			AccessToken:  signedToken.Token.AccessToken,
			RefreshToken: signedToken.Token.RefreshToken,
		},
	}

	return result, nil
}

func (uc *AuthUsecase) RefreshToken(c *gin.Context) (*auth.Authenticate, error) {
	var refreshToken *auth.RefreshTokenForm
	err := c.ShouldBindJSON(&refreshToken)
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(refreshToken.RefreshToken, &auth.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*auth.CustomClaims); ok && token.Valid {

		token, err := uc.GetDetailToken(claims.User.ID)
		if err != nil {
			return nil, err
		}

		if token == nil || token.Revoked {
			return nil, auth.ErrInvalidToken
		}

		err = uc.repo.RevokedRefreshToken(uc.ctx, claims.User.ID)
		if err != nil {
			return nil, err
		}

		user, err := uc.userRepo.GetDetailUser(uc.ctx, claims.User.ID)
		if err != nil {
			return nil, err
		}
		if user == nil {
			return nil, auth.ErrUserNotFound
		}

		timeExpiredToken := time.Minute * time.Duration(60)           // 1 hour
		timeExpiredRefreshToken := time.Minute * time.Duration(10080) // 7 days

		signedToken, err := uc.GenerateJWT(*user, timeExpiredToken, timeExpiredRefreshToken)
		if err != nil {
			return nil, err
		}

		result := &auth.Authenticate{
			User: user,
			Token: auth.Token{
				AccessToken:  signedToken.Token.AccessToken,
				RefreshToken: signedToken.Token.RefreshToken,
			},
		}

		return result, nil
	}

	return nil, auth.ErrInvalidToken
}

func (uc *AuthUsecase) ActivateUser(c *gin.Context) error {
	var token *auth.ActivatedTokenForm
	err := c.ShouldBindJSON(&token)
	if err != nil {
		return err
	}

	userID, err := uc.ValidateActivationToken(token.Token)
	if err != nil {
		return err
	}

	// Aktifkan user di database
	err = uc.userRepo.ActivatedUser(uc.ctx, *userID)
	if err != nil {
		return err
	}

	return nil
}
