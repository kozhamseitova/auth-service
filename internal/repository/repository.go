package repository

import "context"

// "github.com/google/uuid"

type Repository interface {	
	Create(ctx context.Context) (string, error)

	// GetUserById(ctx context.Context, id uuid.UUID) error
	// GetUserByRefreshToken(ctx context.Context, refreshToken string) error
	// UpdateRefreshToken(ctx context.Context) error
}