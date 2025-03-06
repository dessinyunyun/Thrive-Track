package mail

import (
	"strconv"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

func GomailConnection(log *logrus.Entry, email, emailPassword, emailSMTP, port string) *gomail.Dialer {
	var portInt int
	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.Errorf("Invalid port number: %v", err)
		return nil
	}

	client := gomail.NewDialer(emailSMTP, portInt, email, emailPassword)

	if _, err := client.Dial(); err != nil {
		log.Errorf("Gagal terhubung ke server SMTP: %v", err)
	}
	return client
}
