package scraper

import (
	"fmt"
	"log"
	"net/http"
	"shampoo-scraper/src/model"

	"github.com/PuerkitoBio/goquery"
)

var mangaforestBaseURL = "https://mangaforest.com"

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
