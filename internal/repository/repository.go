package repository

import (
	"context"

	"github.com/kozhamseitova/auth-service/internal/entity"
)

type Repository interface {	
	Create(ctx context.Context) (string, error)
	GetUserById(ctx context.Context, id string) (*entity.User, error)
	UpdateRefreshToken(ctx context.Context, id, refreshToken string) error
}