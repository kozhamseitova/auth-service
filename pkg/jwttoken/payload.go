package jwttoken

import (
	"time"

	"github.com/google/uuid"
	"github.com/kozhamseitova/auth-service/utils"
)


type JWTPayload struct {
	ID        uuid.UUID `json:"id"`
	UserId    string     `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (j JWTPayload) Valid() error {
	if time.Now().After(j.ExpiredAt) {
		return utils.ErrExpiredToken
	}

	return nil
}