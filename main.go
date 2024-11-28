package main

import (
	"context"
	"go-gin/app/router"
	"go-gin/database/connection"
	"go-gin/docs"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	godotenv.Load()
}

func main() {
	docs.SwaggerInfo.BasePath = "/" + os.Getenv("PREFIX_API")

	gin.SetMode(os.Getenv("MODE"))
	r := gin.Default()
	ctx := context.Background()
	log := logrus.NewEntry(logrus.StandardLogger())
	db := connection.ConnectionDB(ctx, log, os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	if db == nil {
		os.Exit(1)
	}

	rh := &router.Handlers{
		Ctx: ctx,
		DB:  db,
		R:   r,
		Log: log,
	}

	rh.Routes()

	r.Run()

}
