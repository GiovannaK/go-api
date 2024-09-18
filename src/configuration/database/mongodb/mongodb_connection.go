package mongodb

import (
	"context"
	"os"

	"github.com/GiovannaK/go-api/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGO_URI           = "MONGO_URI"
	MONGO_USER_DATABASE = "MONGO_USER_DATABASE"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongo_uri := os.Getenv(MONGO_URI)
	mongo_database := os.Getenv(MONGO_USER_DATABASE)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongo_uri))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Connected to MongoDB")
	return client.Database(mongo_database), nil

}
