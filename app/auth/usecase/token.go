package usecase

import (
	"context"
	"fmt"
	"go-gin/app/auth"
	"go-gin/database/ent"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	googleUUID "github.com/google/uuid"
)

func (uc *AuthUsecase) GenerateJWT(user ent.User, expiredToken time.Duration, expiredRefreshToken time.Duration) (*auth.Authenticate, error) {
	// For Token
	claimsToken := &auth.CustomClaims{
		User: &user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiredToken).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsToken)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	// For Refresh Token
	claimsRefreshToken := &auth.CustomClaims{
		User: &user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiredRefreshToken).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefreshToken)

	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("REFRESH_SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	now := time.Now()
	expiredTokenTime := now.Add(expiredToken)
	expiredRefreshTokenTime := now.Add(expiredRefreshToken)

	fmt.Println("expiredRefreshToken", expiredRefreshToken)
	fmt.Println("expiredTokenTime", expiredTokenTime)

	res := &auth.Authenticate{
		User: &user,
		Token: auth.Token{
			AccessToken:         tokenString,
			RefreshToken:        refreshTokenString,
			AccessTokenExpired:  expiredTokenTime,
			RefreshTokenExpired: expiredRefreshTokenTime,
		},
	}

	err = uc.repo.CreateToken(uc.ctx, user.ID, res.Token)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type ActivationClaims struct {
	UserID googleUUID.UUID `json:"user_id"`
	jwt.StandardClaims
}

func (uc *AuthUsecase) GenerateActivationToken(ctx context.Context, email, userID string, expiryTime time.Duration) (*string, error) {
	userIdParse, err := uuid.Parse(userID)
	claims := &ActivationClaims{
		UserID: userIdParse,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiryTime).Unix(), // Masa berlaku token
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("ACTIVATION_SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	err = uc.repo.CreateAT(ctx, tokenString, userID)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func (uc *AuthUsecase) ValidateActivationToken(tokenString string) (*googleUUID.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenString, &ActivationClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACTIVATION_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*ActivationClaims); ok && token.Valid {
		user, err := uc.userRepo.GetDetailUser(uc.ctx, claims.UserID)
		if err != nil {
			return nil, err
		}

		if user == nil {
			return nil, auth.ErrUserNotFound
		}

		AT, err := uc.GetDetailAT(claims.UserID)
		if err != nil {
			return nil, err
		}

		if AT.Isused {
			return nil, auth.ErrATalreadyUsed
		}

		err = uc.repo.UsedAT(uc.ctx, claims.UserID)
		if err != nil {
			return nil, err
		}

		return &claims.UserID, nil
	}

	return nil, auth.ErrInvalidToken
}

func (uc *AuthUsecase) GetDetailAT(userID googleUUID.UUID) (*ent.Activation_token, error) {
	res, err := uc.repo.GetDetailAT(uc.ctx, userID)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, auth.ErrATnotFound
	}

	return res, nil
}

func (uc *AuthUsecase) GetDetailToken(userID googleUUID.UUID) (*ent.Token, error) {
	res, err := uc.repo.GetDetailToken(uc.ctx, userID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
