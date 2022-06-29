package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetContext() context.Context {
	//TODO:Check the cancel return
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
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

func GetUserCollection() *mongo.Collection {
	return getClient().Database("test").Collection("user")
}
func GetMangaCollection() *mongo.Collection {
	return getClient().Database("test").Collection("manga_list")
}
