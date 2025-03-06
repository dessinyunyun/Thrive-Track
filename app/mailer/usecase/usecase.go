package usecase

import (
	"bytes"
	"context"
	"go-gin/app/mailer"
	"go-gin/app/mailer/repository"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/go-redis/redis/v8"
)

type EmailUsecase struct {
	emailRepo *repository.EmailRepository
	redis     *redis.Client
	ctx       context.Context
	wg        sync.WaitGroup
}

type EmailJob struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func NewEmailUsecase(emailRepo *repository.EmailRepository, redis *redis.Client, ctx context.Context) *EmailUsecase {
	usecase := &EmailUsecase{
		emailRepo: emailRepo,
		redis:     redis,
		ctx:       ctx,
	}

	return usecase
}

func (uc *EmailUsecase) ActivatedEmail(emailForm mailer.EmailForm) error {

	data := mailer.EmailData{
		Username:      emailForm.EmailData.Username,
		ActivationURL: emailForm.EmailData.ActivationURL,
	}

	body, err := LoadTemplate("user_invitation", data)
	if err != nil {
		return err
	}

	emailForm.Body = body
	emailForm.From = "noreply@aissed-projects.my.id"
	emailForm.Subject = "Finish Registration with GopherSocial"

	return uc.SendEmail(emailForm)
}

func (uc *EmailUsecase) SendEmail(emailForm mailer.EmailForm) error {
	err := uc.emailRepo.SendEmail(uc.ctx, emailForm)
	if err != nil {
		return err
	}
	return nil
}

func LoadTemplate(templateName string, data interface{}) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Println("Gagal mendapatkan direktori kerja:", err)
		return "", err
	}
	templatePath := filepath.Join(cwd, "app", "mailer", "template", templateName+".tmpl")

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println("Error parsing template:", err)
		return "", err
	}

	var bodyBuffer bytes.Buffer
	if err := t.ExecuteTemplate(&bodyBuffer, "body", data); err != nil {
		log.Println("Error executing template:", err)
		return "", err
	}

	return bodyBuffer.String(), nil
}
