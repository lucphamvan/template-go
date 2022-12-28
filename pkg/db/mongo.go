package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectMongoDb() error {
	if os.Getenv("MONGO_URL") == "" {
		MONGO_URL = "mongodb://localhost:27017"
	}

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URL))
	if err != nil {
		return err
	}

	Client = client
	return nil
}

func DisconnectMongoDb() {
	_ = Client.Disconnect(context.Background())
}
