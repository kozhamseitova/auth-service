package repository

import (
	"github.com/kozhamseitova/auth-service/internal/config"
	"github.com/kozhamseitova/auth-service/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type Manager struct {
	client *mongo.Client
	config config.DBConfig
	logger logger.Logger
}

func NewRepository(client *mongo.Client, config config.DBConfig, logger logger.Logger) *Manager {
	return &Manager{
		client: client,
		config: config,
		logger: logger,
	}
}