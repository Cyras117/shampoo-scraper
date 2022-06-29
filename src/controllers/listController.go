package controllers

import (
	"fmt"
	"shampoo-scraper/src/db"
	"shampoo-scraper/src/model"

	"go.mongodb.org/mongo-driver/bson"
)

//TODO:Not working yet
func CreateList(userName, listName string) {
	var list model.List
	list.Name = listName
	mc := db.GetMangaCollection()
	user := FindUsers(userName)

	//TODO:Delete this
	fmt.Println(user[0].ID)

	//res, err := mc.UpdateOne(db.GetContext(), bson.M{"_id": user[0].ID}, bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "lists", Value: listName}}}})
	res, err := mc.UpdateByID(db.GetContext(), user[0].ID, bson.D{{"$set", bson.M{"lists": list}}})

	fmt.Printf("err: %v\n", err)
	fmt.Printf("res: %v\n", res)
}

//TODO:Not working yet
func AddMangaToList(manga model.Manga, ListName string, userName string) {
	c := db.GetUserCollection()
	find := c.FindOne(db.GetContext(), bson.M{"name": userName})
	fmt.Printf("\nResult find: %v\n", find)
	//c.UpdateOne(context.Background(), bson.M{"name": userName})
}
