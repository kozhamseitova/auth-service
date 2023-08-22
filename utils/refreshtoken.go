package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashRefreshToken(refreshToken string) (string, error) {
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedToken), nil
}

func CheckRefreshToken(refreshToken, hashedRefreshToken string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedRefreshToken), []byte(refreshToken))
	if err != nil {
		return err
	}

	return nil
}
