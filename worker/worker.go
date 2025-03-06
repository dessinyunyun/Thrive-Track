package worker

import (
	"context"
	"encoding/json"
	"go-gin/app/mailer"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

type Job struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type Worker struct {
	redisClient *redis.Client
	queueNames  []string
	gomail      *gomail.Dialer
	log         *logrus.Entry
}

func NewWorker(redisClient *redis.Client, gomail *gomail.Dialer, log *logrus.Entry) *Worker {
	return &Worker{
		redisClient: redisClient,
		queueNames:  []string{mailer.QueueEmail},
		gomail:      gomail,
		log:         log,
	}
}

func (w *Worker) Start(ctx context.Context) {
	for {
		result, err := w.redisClient.BLPop(ctx, 5*time.Second, w.queueNames...).Result()
		if err != nil {
			if err == redis.Nil {
				continue
			}
			log.Printf("Failed to fetch job from Redis: %v\n", err)
			time.Sleep(3 * time.Second) // Tunggu sebelum mencoba lagi
			continue
		}

		var job Job
		if err := json.Unmarshal([]byte(result[1]), &job); err != nil {
			log.Printf("Failed to parse job: %v\n", err)
			continue
		}

		w.processJob(ctx, job)
	}
}

func (w *Worker) processJob(ctx context.Context, job Job) {
	switch job.Type {
	case mailer.QueueEmail:
		jsonData, err := json.Marshal(job.Data)
		if err != nil {
			return
		}

		var emailForm mailer.EmailForm
		if err := json.Unmarshal(jsonData, &emailForm); err != nil {
			return
		}

		msg := gomail.NewMessage()
		msg.SetHeader("From", emailForm.From)
		msg.SetHeader("To", emailForm.To)
		msg.SetHeader("Subject", emailForm.Subject)
		msg.SetBody("text/html", emailForm.Body)

		if err := w.gomail.DialAndSend(msg); err != nil {
			job := Job{
				Type: mailer.QueueEmail,
				Data: emailForm,
			}
			jobJSON, _ := json.Marshal(job)

			w.redisClient.RPush(ctx, "dlq", jobJSON).Err()
			w.log.Fatalf("dlq email add. fatal error: %v", err)
		}

	default:
		log.Printf("Unknown job type: %s\n", job.Type)
	}
}
