package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/kozhamseitova/auth-service/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "user"

func (m *Manager) Create(ctx context.Context, user *entity.User) (string, error) {
	userCollection := m.client.Database(m.config.DBName).Collection(collectionName)

	id := uuid.New().String()
	
	newUser := entity.User{
		ID: id,
		Name: user.Name,
		Password: user.Password,
	}

	insertResult, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("%v", insertResult.InsertedID), nil
}


func (m *Manager) GetUserByName(ctx context.Context, name string) (*entity.User, error) {
	userCollection := m.client.Database(m.config.DBName).Collection(collectionName)
	
	filter := bson.D{{"name", name}}
	
	user := new(entity.User)
	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, err
		}
	}
	
	return user, nil
}

func (m *Manager) GetUserById(ctx context.Context, id string) (*entity.User, error) {
	userCollection := m.client.Database(m.config.DBName).Collection(collectionName)
	
	filter := bson.D{{"_id", id}}
	
	user := new(entity.User)
	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, err
		}
	}
	
	return user, nil
}

// func (m *Manager) GetUserByRefreshToken(ctx context.Context, refreshToken string) error {

// }

func (m *Manager) UpdateRefreshToken(ctx context.Context, id, refreshToken string) error {
	userCollection := m.client.Database(m.config.DBName).Collection(collectionName)

	filter := bson.D{{"_id", id}}

	update := bson.D{{"$set", bson.D{{"refresh_token", refreshToken}}},
					{"$setOnInsert", bson.D{{"refresh_token", refreshToken}}},}

	_, err := userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
	
}