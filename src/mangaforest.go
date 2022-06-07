package scraper

import (
	"fmt"
	"log"
	"net/http"
	"shampoo-scraper/src/model"

	"github.com/PuerkitoBio/goquery"
)

var mangaforestBaseURL = "https://mangaforest.com"

//Gets the last 5 pages of updates from the site and returns
func GetmangaforestLatestUpdatesPage() []model.Manga {
	var listaMangas []model.Manga

	for i := 1; i < 6; i++ {
		latestUpdatePageURL := mangaforestBaseURL + fmt.Sprintf("/latest?page=%d", i)

		res, err := http.Get(latestUpdatePageURL)

		//TODO:find a way to not break the application
		if err != nil {
			log.Fatal(err.Error())
		}
		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)
		//TODO:find a way to not break the application
		if err != nil {
			log.Fatal(err.Error())
		}

		doc.Find("div.book-item").Find("div.book-detailed-item").Each(
			func(i int, s *goquery.Selection) {
				node := s.Find("div.thumb")
				title, _ := node.Find("a").Attr("title")
				path, _ := node.Find("a").Attr("href")

				listaMangas = append(listaMangas,
					model.Manga{Title: title, Path: path, SiteURL: mangaforestBaseURL, ImgURL: ""})
			},
		)
	}
	return listaMangas
}

//Searches on the mangaforest api for a manga with the specified phrase.
func SearchMangamangaforest(phrase string) []model.Manga {
	searchURL := "https://mangaforest.com/api/manga/search?q="
	var mangas []model.Manga
	requestURL := searchURL + phrase

	res, err := http.Get(requestURL)
	if err != nil {
		log.Panic(err.Error())
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Panic(err.Error())
	}
	defer res.Body.Close()
	doc.Find("div.name").Each(
		func(i int, s *goquery.Selection) {
			title := s.Find("a").Text()
			path, _ := s.Find("a").Attr("href")
			mangas = append(mangas,
				model.Manga{Title: title, Path: path, SiteURL: mangaforestBaseURL, ImgURL: ""})
		},
	)
	return mangas
}
