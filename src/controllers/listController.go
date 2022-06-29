package controllers

import (
	"context"
	"fmt"
	"shampoo-scraper/src/db"
	"shampoo-scraper/src/model"

	"go.mongodb.org/mongo-driver/bson"
)

//TODO:Not working yet
func CreateList(userName, listName string) {
	var list model.List
	list.Name = listName
	c := db.GetCollection()
	res, err := c.UpdateOne(context.Background(), bson.M{"name": userName}, bson.M{"lists": list})

	fmt.Printf("err: %v\n", err)
	fmt.Printf("res: %v\n", res)
}

//TODO:Not working yet
func AddMangaToList(manga model.Manga, ListName string, userName string) {
	c := db.GetCollection()
	find := c.FindOne(context.Background(), bson.M{"name": userName})
	fmt.Printf("\nResult find: %v\n", find)
	//c.UpdateOne(context.Background(), bson.M{"name": userName})
}
