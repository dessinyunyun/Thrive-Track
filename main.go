package main

import (
	"context"
	"go-gin/app/router"
	"go-gin/database/connection"
	"go-gin/docs"
	"go-gin/mail"
	"go-gin/redis"
	"go-gin/worker"
	"log"
	"os"
	"time"

	_ "time/tzdata"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	godotenv.Load()
}

func main() {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatalf("Gagal memuat lokasi: %v", err)
	}

	// Menetapkan zona waktu lokal
	time.Local = loc

	log := logrus.NewEntry(logrus.StandardLogger())
	docs.SwaggerInfo.BasePath = "/" + os.Getenv("PREFIX_API")

	gin.SetMode(os.Getenv("MODE"))
	r := gin.Default()
	ctx := context.Background()
	db := connection.ConnectionDB(ctx, log, os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	gomail := mail.GomailConnection(log, os.Getenv("EMAIL"), os.Getenv("EMAIL_PASSWORD"), os.Getenv("EMAIL_SMTP"), os.Getenv("EMAIL_PORT"))
	redis := redis.NewRedisClient(log)
	worker := worker.NewWorker(redis, gomail, log)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if db == nil {
		os.Exit(1)
	}

	rh := &router.Handlers{
		Ctx:    ctx,
		DB:     db,
		R:      r,
		Log:    log,
		Gomail: gomail,
		Redis:  redis,
		Worker: worker,
	}

	go worker.Start(ctx)

	rh.Routes()

	r.Run()
	select {}

}
