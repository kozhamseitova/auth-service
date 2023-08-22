package repository

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetUserById(ctx context.Context, id uuid.UUID) error
	GetUserByRefreshToken(ctx context.Context, refreshToken string) error
	UpdateRefreshToken(ctx context.Context) error
}