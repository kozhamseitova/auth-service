package service

import (
	"github.com/kozhamseitova/auth-service/internal/config"
	"github.com/kozhamseitova/auth-service/internal/repository"
	"github.com/kozhamseitova/auth-service/pkg/jwttoken"
	"github.com/kozhamseitova/auth-service/pkg/logger"
)

type Manager struct {
	repository repository.Repository
	token *jwttoken.JWTToken
	config config.TokenConfig
	logger logger.Logger
}

func NewService(repository repository.Repository, token *jwttoken.JWTToken, config config.TokenConfig, logger logger.Logger) *Manager {
	return &Manager{
		repository: repository,
		token: token,
		config: config,
		logger: logger,
	}
}