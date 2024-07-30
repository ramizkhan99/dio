package db

import (
	"context"

	"github.com/ramizkhan99/dio-be/internal/config"
	"github.com/ramizkhan99/dio-be/pkg/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() *mongo.Client {
	Client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.GetDBUrl()))
	if err != nil {
		logger.GetSugaredLogger().Panicf("failed to connect to database: %w", err)
	}

	logger.GetSugaredLogger().Infof("Connected to DB!!")
	return Client
}

func GetDBClient() *mongo.Client {
	return Client
}

func Disconnect() {
	if err := Client.Disconnect(context.TODO()); err != nil {
		logger.GetSugaredLogger().Panicf("failed to disconnect from database: %w", err)
	}
	logger.GetSugaredLogger().Infof("Database disconnected sucessfully")
}

func Insert[T any](collectionName string, entity T) (string, error) {
	collection := Connect().Database("dio").Collection(collectionName)
	insertResult, err := collection.InsertOne(context.TODO(), entity)
	if err != nil {
		logger.GetSugaredLogger().Errorf("failed to insert document :: %v", err)
		return "", err
	}
	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

func Delete(collection mongo.Collection, entity interface{}) (interface{}, error) {
	return nil, nil
}

func GetUsersCollection() *mongo.Collection {
	return Connect().Database("dio").Collection("users")
}
