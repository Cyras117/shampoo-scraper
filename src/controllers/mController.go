package controllers

import (
	"fmt"
	scraper "shampoo-scraper/src"
	"shampoo-scraper/src/db"
	"shampoo-scraper/src/model"

	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson"
)

func FindManga(path string) model.Manga {
	var result []model.Manga
	c := db.GetMangaCollection()

	cursor, err := c.Find(db.GetContext(), bson.M{"path": path})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(db.GetContext(), &result); err != nil {
		panic(err)
	}
	defer cursor.Close(db.GetContext())

	if len(result) == 0 {
		return model.Manga{}
	}
	return result[0]
}

func CreateManga(manga model.Manga) {
	res := FindManga(manga.Path)

	if cmp.Equal(res, model.Manga{}) {
		_, err := db.GetMangaCollection().InsertOne(db.GetContext(), manga)
		if err != nil {
			panic(err)
		}
		return
	}
	println("Manga already registered!")
}

func DeleteManga(manga model.Manga) {
	mangaRsult := FindManga(manga.Path)

	if cmp.Equal(mangaRsult, model.Manga{}) {
		println("Manga not found")
		return
	}

	_, err := db.GetMangaCollection().DeleteOne(db.GetContext(), bson.M{"path": manga.Path})
	if err != nil {
		panic(err)
	}
}

func UpdateManga(manga model.Manga) {
	mangaUpdated := scraper.GetMangaWithPath(manga.Path)
	c := db.GetMangaCollection()

	mangaRsult := FindManga(manga.Path)
	if cmp.Equal(mangaRsult, model.Manga{}) {
		println("Manga not found")
		return
	}

	_, err := c.UpdateOne(db.GetContext(), bson.M{"path": manga.Path}, bson.M{"$set": mangaUpdated})
	if err != nil {
		panic(err)
	}
}

//TODO DO NOT RUN THIS
func AddAllMangas() {
	allMangas := scraper.GetAllMangas()

	for _, manga := range allMangas {
		CreateManga(manga)
		fmt.Printf("Created manga: %v\n", manga)
	}
}
