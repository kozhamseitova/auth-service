package handler

import (
	"github.com/kozhamseitova/auth-service/internal/config"
	"github.com/kozhamseitova/auth-service/internal/service"
)

type Handler struct {
	service service.Service
	config     *config.Config
}

func NewHandler(service service.Service, config *config.Config) *Handler {
	return &Handler{
		service: service,
		config:     config,
	}
}