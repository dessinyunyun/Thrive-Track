package mailer

import (
	"context"
)

type EmailForm struct {
	To        string    `json:"to"`
	From      string    `json:"from"`
	Body      string    `json:"body"`
	Subject   string    `json:"subject"`
	EmailData EmailData `json:"email_data"`
}

type EmailData struct {
	Username      string `json:"username"`
	ActivationURL string `json:"activation_url"`
}

type MailerUsecase interface {
	ActivatedEmail(emailForm EmailForm) error
	SendEmail(emailForm EmailForm) error
}

type MailerRepository interface {
	SendEmail(ctx context.Context, emailForm EmailForm) error
}

var (
	QueueEmail = "queue-email"
)
