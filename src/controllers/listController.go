package controllers

import (
	"shampoo-scraper/src/db"
	"shampoo-scraper/src/model"

	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson"
)

/*
	TODO Modify this theres is no user.
*/

func FindList(name string) model.List {
	var lists []model.List

	c := db.GetListCollection()

	cursor, err := c.Find(db.GetContext(), bson.M{"name": name})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(db.GetContext(), &lists); err != nil {
		panic(err)
	}
	defer cursor.Close(db.GetContext())
	if len(lists) == 0 {
		return model.List{}
	}
	return lists[0]
}

func CreateList(list model.List) {
	listFound := FindList(list.Name)
	if !cmp.Equal(listFound, model.List{}) {
		println("List already exists!")
		return
	}
	if len(list.Collection) == 0 {
		_, err := db.GetListCollection().InsertOne(db.GetContext(), bson.D{{Key: "name", Value: list.Name}, {Key: "Collection", Value: bson.A{}}})
		if err != nil {
			panic(err)
		}
		return
	}
	_, err := db.GetListCollection().InsertOne(db.GetContext(), list)
	if err != nil {
		panic(err)
	}
}

func DeleteList(list model.List) {
	ListFound := FindList(list.Name)
	if cmp.Equal(ListFound, model.List{}) {
		println("List not found!")
		return
	}
	_, err := db.GetListCollection().DeleteOne(db.GetContext(), bson.M{"name": list.Name})
	if err != nil {
		panic(err)
	}
}

func UpdateListName(list model.List, name string) {
	c := db.GetListCollection()

	ListUpdate := list

	ListUpdate.Name = name

	ListFound := FindList(list.Name)

	if cmp.Equal(ListFound, model.List{}) {
		println("List not found")
		return
	}

	_, err := c.UpdateOne(db.GetContext(), bson.M{"name": list.Name}, bson.M{"$set": ListUpdate})
	if err != nil {
		panic(err)
	}
}

func AddMangaToList(listName, mangaPath string) {
	FindListRes := FindList(listName)
	FindMangaRes := FindManga(mangaPath)

	//TODO Verify if manga already exists in list

	if cmp.Equal(FindMangaRes, model.Manga{}) {

		//TODO:Check on readm.org if the manga exests if so, add it to manga database
		println("List not found!")
		return
	}

	if cmp.Equal(FindListRes, model.List{}) {
		println("List not found!")
		return
	}

	FindListRes.Collection = append(FindListRes.Collection, mangaPath)

	c := db.GetListCollection()

	_, err := c.UpdateOne(db.GetContext(), bson.M{"name": FindListRes.Name}, bson.M{"$set": FindListRes})
	if err != nil {
		panic(err)
	}
}

//TODO
func RemoveMangaFromList(listName, mangaPath string) {

}
