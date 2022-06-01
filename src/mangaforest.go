package scraper

import (
	"log"
	"net/http"
	"shampoo-scraper/src/model"

	"github.com/PuerkitoBio/goquery"
)

var mangaforestBaseURL = "https://mangaforest.com"

//TODO:get also de upates from the second page
func GetmangaforestLatestUpdatesPage() []model.Manga {
	latestUpdatePageURL := mangaforestBaseURL + "/latest"
	var listaMangas []model.Manga
	res, err := http.Get(latestUpdatePageURL)

	//TODO:find a way to not break the application
	if err != nil {
		log.Fatal(err.Error())
	}
	defer res.Body.Close()
	//TODO:find a way to not break the application
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Statuscode error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	//TODO:find a way to not break the application
	if err != nil {
		log.Fatal(err.Error())
	}

	doc.Find("div.book-item").Find("div.book-detailed-item").Each(
		func(i int, s *goquery.Selection) {
			node := s.Find("div.thumb")
			title, titleExists := node.Find("a").Attr("title")
			path, pathExist := node.Find("a").Attr("href")

			//TODO:Change to only omit the atributes that doesn't exist~.
			if !pathExist || !titleExists {
				listaMangas = append(listaMangas,
					model.Manga{Title: title, Path: "", SiteURL: mangaforestBaseURL, ImgURL: ""})
				return
			}
			listaMangas = append(listaMangas,
				model.Manga{Title: title, Path: path, SiteURL: mangaforestBaseURL, ImgURL: ""})
		},
	)
	//TODO: Try to change it for []struct instead of an [][]string
	return listaMangas
}
