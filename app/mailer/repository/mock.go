package repository

import (
	"context"
	"go-gin/app/mailer"

	"github.com/go-redis/redis/v8"
	"gopkg.in/gomail.v2"
)

type MockRepository struct {
	gomail *gomail.Dialer
	redis  *redis.Client
}

func (m *MockRepository) SendEmail(ctx context.Context, emailForm mailer.EmailForm) error {
	// args := m.Called(ctx, emailForm)
	return nil
}
