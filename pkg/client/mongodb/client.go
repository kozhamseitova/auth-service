package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	host     string
	port     string
	dbName   string
	client *mongo.Client
}

func Connect(opts ...Option) (*mongo.Client, error) {
	m := new(MongoDB)

	for _, opt := range opts {
		opt(m)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://" + m.host + m.port))
	if err != nil {
		return nil, fmt.Errorf("mongo connect err: %w", err)
	}

	// coll := client.Database(m.dbName).Collection(m.collectionName)

	return client, nil
}

func (m MongoDB) Close() {
	if m.client != nil {
		m.client.Disconnect(context.TODO())
	}
}