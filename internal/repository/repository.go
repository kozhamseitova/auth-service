package repository

import (
	"context"

	"github.com/kozhamseitova/auth-service/internal/entity"
	// "github.com/google/uuid"
)

type Repository interface {	
	Create(ctx context.Context, user *entity.User) (string, error)
	GetUserByName(ctx context.Context, name string) (*entity.User, error)
	GetUserById(ctx context.Context, id string) (*entity.User, error)
	// GetUserByRefreshToken(ctx context.Context, refreshToken string) (string, error)
	UpdateRefreshToken(ctx context.Context, id, refreshToken string) error
}