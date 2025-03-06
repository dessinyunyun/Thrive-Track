package usecase

import (
	"go-gin/app/mailer"

	"github.com/stretchr/testify/mock"
)

type MockEmailUsecase struct {
	mock.Mock
}

func (uc *MockEmailUsecase) ActivatedEmail(emailForm mailer.EmailForm) error {

	return nil
}

func (uc *MockEmailUsecase) SendEmail(emailForm mailer.EmailForm) error {

	return nil
}
