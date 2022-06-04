package scraper

import (
	"fmt"
	"log"
	"net/http"
	"shampoo-scraper/src/model"

	"github.com/PuerkitoBio/goquery"
)

var readmBaseURL = "https://www.readm.org"

func GetreadmLatestUpdatesPage() []model.Manga {
	var listaMangas []model.Manga

	for i := 1; i < 6; i++ {
		latestUpdatePageURL := readmBaseURL + fmt.Sprintf("/latest-releases/%d", i)

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

		doc.Find("ul.clearfix").Find("li.segment-poster-sm").Each(
			func(i int, s *goquery.Selection) {
				node := s.Find("h2")
				title := node.Find("a").Text()
				path, _ := node.Find("a").Attr("href")
				node = s.Find("img")
				imgPath, _ := node.Attr("data-src")

				listaMangas = append(listaMangas,
					model.Manga{Title: title, Path: path, SiteURL: readmBaseURL, ImgURL: imgPath})
			},
		)
	}
	return listaMangas
}
