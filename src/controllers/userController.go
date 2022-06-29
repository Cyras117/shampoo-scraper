package controllers

import (
	"context"
	"fmt"
	"log"
	"shampoo-scraper/src/db"
	"shampoo-scraper/src/model"

	"go.mongodb.org/mongo-driver/bson"
)

//TODO:Make sure that the user name is unique
func CreateUser(name string) {
	var user model.User
	user.Name = name
	result, err := db.GetCollection().InsertOne(context.Background(), user)
	if err != nil {
		panic(err.Error())
	}

	println("User created: ", result.InsertedID)
}

func FindUser(name string) {
	var users []model.User
	c := db.GetCollection()

	res, err := c.Find(context.Background(), bson.M{"name": name})
	if err != nil {
		panic(err.Error())
	}
	if err = res.All(context.Background(), &users); err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)

}

//TODO:Not working yet
func DeleteUser(name string) {
	res, _ := db.GetCollection().DeleteOne(context.Background(), bson.M{"name": name})
	fmt.Printf("res: %v\n", res)
}
