package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) (string, error) {
	if len(pwd) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(h), err
}

func CheckPassword(hash, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
