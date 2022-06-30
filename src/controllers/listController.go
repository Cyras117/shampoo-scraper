package controllers

import (
	"fmt"
	"shampoo-scraper/src/db"
	"shampoo-scraper/src/model"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateList(userName, listName string) {
	user := FindUsers(userName)[0]
	c := db.GetUserCollection()
	res, err := c.UpdateOne(db.GetContext(), bson.M{"_id": user.ID}, bson.M{"$push": bson.M{"lists": bson.M{"name": listName, "colection": bson.A{}}}})
	fmt.Printf("res: %v\n", res)
	fmt.Printf("err: %v\n", err)
}

//TODO:Not working yet
func AddMangaToList(manga model.Manga, ListName string, userName string) {
	c := db.GetUserCollection()
	find := c.FindOne(db.GetContext(), bson.M{"name": userName})
	fmt.Printf("\nResult find: %v\n", find)
	//c.UpdateOne(context.Background(), bson.M{"name": userName})
}
