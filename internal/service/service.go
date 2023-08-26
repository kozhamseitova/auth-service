package service

import (
	"context"

	"github.com/kozhamseitova/auth-service/internal/entity"
)

type Service interface {
	Create(ctx context.Context) (string, error)
	Login(ctx context.Context, id string) (*entity.Token, error)
	Refresh(ctx context.Context, user entity.User) (*entity.Token, error)
	VerifyToken(ctx context.Context, token string) (string, error)
}