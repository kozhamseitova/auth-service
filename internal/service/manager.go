package service

import (
	"github.com/kozhamseitova/auth-service/internal/repository"
	"github.com/kozhamseitova/auth-service/pkg/jwttoken"
)

type Manager struct {
	repository repository.Repository
	token *jwttoken.JWTToken
}

func NewService(repository repository.Repository, token *jwttoken.JWTToken) *Manager {
	return &Manager{
		repository: repository,
		token: token,
	}
}