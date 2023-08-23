package service

import (
	"github.com/kozhamseitova/auth-service/internal/config"
	"github.com/kozhamseitova/auth-service/internal/repository"
	"github.com/kozhamseitova/auth-service/pkg/jwttoken"
)

type Manager struct {
	repository repository.Repository
	token *jwttoken.JWTToken
	config config.TokenConfig
}

func NewService(repository repository.Repository, token *jwttoken.JWTToken, config config.TokenConfig) *Manager {
	return &Manager{
		repository: repository,
		token: token,
		config: config,
	}
}