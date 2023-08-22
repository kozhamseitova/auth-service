package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Manager struct {
	client *mongo.Client
}

func NewRepository(client *mongo.Client) *Manager {
	return &Manager{
		client: client,
	}
}