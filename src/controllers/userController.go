package controllers

import (
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
	result, err := db.GetUserCollection().InsertOne(db.GetContext(), user)
	if err != nil {
		panic(err.Error())
	}

	println("User created: ", result.InsertedID)
}

func FindUsers(name string) []model.User {
	var users []model.User
	c := db.GetUserCollection()

	cursor, err := c.Find(db.GetContext(), bson.M{"name": name})
	if err != nil {
		panic(err.Error())
	}
	if err = cursor.All(db.GetContext(), &users); err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(db.GetContext())
	return users
}

func DeleteUser(name string) {
	res, _ := db.GetUserCollection().DeleteOne(db.GetContext(), bson.M{"name": name})
	fmt.Printf("res: %v\n", res)
}

//TODO:Think about it
func UpdateUser(id, name string, filter model.User) {

}
