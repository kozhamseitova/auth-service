package service

import (
	"context"

	"github.com/kozhamseitova/auth-service/internal/entity"
)

// "github.com/google/uuid"

type Service interface {
	Create(ctx context.Context, user *entity.User) (string, error)
	Login(ctx context.Context, name, password string) (string, string, error)
	Refresh(ctx context.Context, id, refreshToken string) (string, string, error)
}