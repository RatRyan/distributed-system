package service

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func filterById(id string) (*primitive.M, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objId}
	return &filter, nil
}

func insertObj[T any](obj *T, c *mongo.Collection) (*primitive.ObjectID, error) {
	result, err := c.InsertOne(context.Background(), obj)
	if err != nil {
		return nil, err
	}

	newObjId, _ := result.InsertedID.(primitive.ObjectID)
	return &newObjId, nil
}

func getObjById[T any](id string, c *mongo.Collection) (*T, error) {
	filter, _ := filterById(id)

	var foundObj T
	err := c.FindOne(context.Background(), filter).Decode(&foundObj)
	if err != nil {
		return nil, err
	}

	return &foundObj, nil
}

func updateObj[T any](obj *T, id string, c *mongo.Collection) error {
	filter, _ := filterById(id)
	update := bson.M{"$set": obj}

	result, err := c.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("no documents were replaced")
	}

	return nil
}

func deleteObj[T any](id string, c *mongo.Collection) error {
	filter, _ := filterById(id)

	result, err := c.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("no documents were deleted")
	}

	return nil
}
