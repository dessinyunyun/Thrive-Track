package mail

import (
	"github.com/stretchr/testify/mock"
	"gopkg.in/gomail.v2"
)

// Define an interface that includes the methods you use from gomail.Dialer
type MailDialer interface {
	DialAndSend(...*gomail.Message) error
}

// Create a mock struct that implements the MailDialer interface
type MockMailDialer struct {
	mock.Mock
}

// Implement the DialAndSend method
func (m *MockMailDialer) DialAndSend(msgs ...*gomail.Message) error {
	args := m.Called(msgs)
	return args.Error(0)
}
