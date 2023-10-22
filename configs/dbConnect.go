package configs

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDatabase() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:123@localhost:27017/?authSource=admin"));

	if err != nil {
		println(Red, "database connection error", Reset);
	}

	err = client.Ping(ctx, readpref.Primary());

	if err != nil {
		println(Red, "error", err, Reset);
	}

	println(Green, "Connect database successfully", Reset);

	return client
}