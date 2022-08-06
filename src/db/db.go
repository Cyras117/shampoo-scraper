package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const database = "test"

func GetContext() context.Context {
	//TODO:Check the cancel return
	ctx := context.Background()
	return ctx
}

func getClient() *mongo.Client {
	//TODO:Check Context cancel
	//TODO:Check Disconnect

	client, err := mongo.Connect(GetContext(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err.Error())
	}
	return client
}

func GetMangaCollection() *mongo.Collection {
	return getClient().Database(database).Collection("manga")
}

func GetListCollection() *mongo.Collection {
	return getClient().Database(database).Collection("lists")
}
