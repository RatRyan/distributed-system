package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func Init() error {
	uri := "mongodb://database:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var err error
	DB, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	return nil
}

func GetUserColl() *mongo.Collection {
	return DB.Database("data").Collection("users")
}

func GetGameColl() *mongo.Collection {
	return DB.Database("data").Collection("games")
}

func GetOfferColl() *mongo.Collection {
	return DB.Database("data").Collection("offers")
}
