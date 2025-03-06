package handler

import (
	"fmt"
	"go-gin/app/mailer"
	"go-gin/app/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MailerHandler struct {
	uc  mailer.MailerUsecase
	log *logrus.Entry
}

func MailerRoute(uc mailer.MailerUsecase, r *gin.RouterGroup, log *logrus.Entry) {
	h := MailerHandler{
		uc:  uc,
		log: log,
	}

	v2 := r.Group("mailer", middleware.AuthMiddleware(log))
	fmt.Println(h)
	fmt.Println(v2)
	// v2.GET("/", h.GetAll)
	// v2.GET("/:id", h.GetDetail)
	// v2.POST("", h.SendEmail)

}

// func (h *MailerHandler) SendEmail(c *gin.Context) {
// 	err := h.uc.ActivatedEmail(c)
// 	if err != nil {
// 		h.log.Errorf("create User handlers: %v", err)
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"message": err,
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, tools.Response{
// 		Status:  "success",
// 		Message: "Create User",
// 	})
// }
