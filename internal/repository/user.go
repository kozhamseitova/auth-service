package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/kozhamseitova/auth-service/internal/entity"
	"github.com/kozhamseitova/auth-service/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "user"

func (m *Manager) Create(ctx context.Context) (string, error) {
	userCollection := m.client.Database(m.config.DBName).Collection(collectionName)

	id := uuid.New().String()
	
	newUser := entity.User{
		ID: id,
	}

	insertResult, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		m.logger.Errorf(ctx, "[Create] err: %v", err)
		return "", utils.ErrInternalError
	}
	return fmt.Sprintf("%v", insertResult.InsertedID), nil
}

func (m *Manager) GetUserById(ctx context.Context, id string) (*entity.User, error) {
	userCollection := m.client.Database(m.config.DBName).Collection(collectionName)
	
	filter := bson.D{{"_id", id}}
	
	user := new(entity.User)
	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		m.logger.Errorf(ctx, "[GetUserById] err: %v", err)
		if errors.Is(err, mongo.ErrNoDocuments){
			return nil, utils.ErrUserNotFound
		}
		return nil, utils.ErrInternalError
	}
	
	return user, nil
}

func (m *Manager) GetByRefreshToken(ctx context.Context, refreshToken string) (*entity.User, error) {
	userCollection := m.client.Database(m.config.DBName).Collection(collectionName)
	
	filter := bson.D{{"refresh_token", refreshToken}}
	
	user := new(entity.User)
	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		m.logger.Errorf(ctx, "[GetByRefreshToken] err: %v", err)
		if errors.Is(err, mongo.ErrNoDocuments){
			return nil, utils.ErrUserNotFound
		}
		return nil, utils.ErrInternalError
	}
	
	return user, nil
}

func (m *Manager) UpdateRefreshToken(ctx context.Context, id, refreshToken string) error {
	userCollection := m.client.Database(m.config.DBName).Collection(collectionName)

	filter := bson.D{{"_id", id}}

	update := bson.D{{"$set", bson.D{{"refresh_token", refreshToken}}}}
	
	_, err := userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		m.logger.Errorf(ctx, "[UpdateRefreshToken] err: %v", err)
		return utils.ErrInternalError
	}

	return nil
	
}