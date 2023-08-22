package service

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	Login(ctx context.Context, id uuid.UUID) (string, string, error)
	Refresh(ctx context.Context, refreshToken string) (string, string, error)
}