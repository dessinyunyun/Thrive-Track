package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"go-gin/app/mailer"
	"go-gin/worker"

	"github.com/go-redis/redis/v8"
	"gopkg.in/gomail.v2"
)

type EmailRepository struct {
	gomail *gomail.Dialer
	redis  *redis.Client
}

func NewEmailRepository(gomail *gomail.Dialer, redis *redis.Client) *EmailRepository {
	return &EmailRepository{gomail: gomail, redis: redis}
}

func (r *EmailRepository) SendEmail(ctx context.Context, emailForm mailer.EmailForm) error {

	msg := gomail.NewMessage()
	msg.SetHeader("From", emailForm.From)
	msg.SetHeader("To", emailForm.To)
	msg.SetHeader("Subject", emailForm.Subject)
	msg.SetBody("text/html", emailForm.Body)

	if err := r.gomail.DialAndSend(msg); err != nil {
		job := worker.Job{
			Type: mailer.QueueEmail,
			Data: emailForm,
		}
		jobJSON, _ := json.Marshal(job)

		fmt.Println("failed to send email: %w", err)
		return r.redis.RPush(ctx, mailer.QueueEmail, jobJSON).Err()

	}
	return nil

}
