package middleware

import (
	"go-gin/app/auth"
	"go-gin/app/tools"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AuthMiddleware(log *logrus.Entry) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		parts := strings.Fields(tokenString)

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, tools.Response{
				Message: "Token is missing",
			})
			c.Abort()
			return
		}

		if len(parts) < 2 {
			c.JSON(http.StatusUnauthorized, tools.Response{
				Message: "Token is missing",
				Status:  "error",
			})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(parts[1], &auth.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, tools.Response{
				Message: "Invalid token",
				Status:  "error",
			})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*auth.CustomClaims); ok && token.Valid {

			userID, _ := tools.ConvertUUID(string(claims.User.ID.ID()))
			c.Set("user_id", userID)

			c.Next()

		} else {
			c.JSON(http.StatusUnauthorized, tools.Response{
				Message: "Unauthorized",
				Status:  "error",
			})
			c.Abort()
		}
	}
}
