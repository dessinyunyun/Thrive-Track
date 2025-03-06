package handler

import (
	"go-gin/app/auth"
	"go-gin/app/tools"
	"net/http"

	"github.com/gin-gonic/gin"
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

	v1 := r.Group("auth")

	v1.POST("/register", h.Register)
	v1.POST("/login", h.Login)
	v1.PATCH("/refresh-token", h.RefreshToken)
	v1.PATCH("/activated-client", h.ActivatedClient)
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
		auth.ErrorHandler(c, h.log, err)
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
		auth.ErrorHandler(c, h.log, err)
		return
	}

	c.JSON(http.StatusOK, tools.Response{
		Status:  "success",
		Message: "success login",
		Data:    result,
	})
}

// @Tags Auth
// @Summary RefreshToken
// @Description RefreshToken
// @Router /auth/refresh-token [patch]
// @Accept json
// @Produce json
// @Param request body auth.RefreshTokenForm true "Payload Body for RefreshToken [RAW]"
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	result, err := h.uc.RefreshToken(c)
	if err != nil {
		auth.ErrorHandler(c, h.log, err)
		return
	}

	c.JSON(http.StatusOK, tools.Response{
		Status:  "success",
		Message: "success refresh token",
		Data:    result,
	})
}

// @Tags Auth
// @Summary ActivatedClient
// @Description ActivatedClient with token
// @Router /auth/activated-client [patch]
// @Accept json
// @Produce json
// @Param request body auth.ActivatedTokenForm true "Payload Body for Patch [RAW]"
func (h *AuthHandler) ActivatedClient(c *gin.Context) {
	err := h.uc.ActivateUser(c)
	if err != nil {
		auth.ErrorHandler(c, h.log, err)
		return
	}

	c.JSON(http.StatusOK, tools.Response{
		Status:  "success",
		Message: "success activated",
	})
}
