package service

import (
	"context"

	"github.com/kozhamseitova/auth-service/internal/repository"
)

type Manager struct {
	ctx context.Context
	repository repository.Repository
}

func NewService(ctx context.Context, repository repository.Repository) *Manager {
	return &Manager{
		ctx: ctx,
		repository: repository,
	}
}