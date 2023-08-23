package repository

import (
	"github.com/kozhamseitova/auth-service/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Manager struct {
	client *mongo.Client
	config config.DBConfig
}

func NewRepository(client *mongo.Client, config config.DBConfig) *Manager {
	return &Manager{
		client: client,
		config: config,
	}
}