package utils

import (
	"encoding/base64"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

func GenerateRefreshToken() (string, error) {
	n := rand.Intn(11) + 10
	refreshToken := make([]byte, n)
	_, err := rand.Read(refreshToken)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(refreshToken), nil
}
