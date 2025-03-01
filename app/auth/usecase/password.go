package usecase

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CheckComparePass(clientProvidePassword string, userPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(clientProvidePassword))
	if err != nil {

		return err

	}
	return nil
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}
