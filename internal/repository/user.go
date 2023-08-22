package repository

import (
	"context"
	"fmt"

	"github.com/kozhamseitova/auth-service/internal/entity"
)

const collectionName = "user"

func (m *Manager) Create(ctx context.Context) (string, error) {
	userCollection := m.client.Database("auth-service").Collection("user")

	// Create a new user document
	newUser := entity.User{
		Name: "john_doe",
	}

	// Insert the user document
	insertResult, err := userCollection.InsertOne(context.Background(), newUser)
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("%v", insertResult.InsertedID), nil
}