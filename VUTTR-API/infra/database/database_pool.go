package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewConnect(db_connect_url string, ctx context.Context) *mongo.Client {

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db_connect_url))

	if err != nil {
		panic(err)
	}

	return client
}
