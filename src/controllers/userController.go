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
	result, err := db.GetUserCollection().InsertOne(db.GetContext(), bson.M{"name": name, "lists": bson.A{}})
	if err != nil {
		panic(err.Error())
	}
	println("User created: ", result.InsertedID)
}

//TODO:Pass Filter
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

//TODO:Pass Filter
func DeleteUser(name string) {
	res, _ := db.GetUserCollection().DeleteOne(db.GetContext(), bson.M{"name": name})
	fmt.Printf("res: %v\n", res)
}

func UpdateUserName(name, nameUpdate string) {
	c := db.GetUserCollection()

	user := FindUsers(name)[0]
	user.Name = nameUpdate
	res, err := c.UpdateOne(db.GetContext(), bson.D{{Key: "_id", Value: user.ID}}, bson.D{{Key: "$set", Value: user}})
	fmt.Printf("res.MatchedCount: %v\n", res.MatchedCount)
	fmt.Printf("err: %v\n", err)
}

//TODO:Think about it
func UpdateUserLists(filter, updateModel model.User) {
	c := db.GetUserCollection()

	user := FindUsers(filter.Name)[0]
	fmt.Println(user)
	user.Lists = append(user.Lists, updateModel.Lists...)
	fmt.Println(user.Name)
	res, err := c.UpdateOne(db.GetContext(), filter, bson.M{"$set": user})

	fmt.Printf("err: %v\n", err)
	fmt.Printf("res: %v\n", res)
}
