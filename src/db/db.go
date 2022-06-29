package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getClient() *mongo.Client {
	//TODO:Check Context cancel
	//TODO:Check Disconnect
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err.Error())
	}
	return client
}

func GetCollection() *mongo.Collection {
	return getClient().Database("test").Collection("user")
}
