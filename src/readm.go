package scraper

import (
	"log"
	"net/http"
	"shampoo-scraper/src/model"

	"github.com/PuerkitoBio/goquery"
)

var readmBaseURL = "https://www.readm.org"

//TODO:get also de upates from the second page
func GetreadmLatestUpdatesPage() []model.Manga {
	latestUpdatePageURL := readmBaseURL + "/latest-releases"
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
	//TODO: Try to change it for []struct instead of an [][]string
	doc.Find("ul.clearfix").Find("li.segment-poster-sm").Each(
		func(i int, s *goquery.Selection) {
			node := s.Find("h2")
			title := node.Find("a").Text()
			path, pathExist := node.Find("a").Attr("href")
			node = s.Find("img")
			imgPath, imgPathExist := node.Attr("data-src")

			//TODO:Change to only omit the atributes that doesn't exist~.
			if !pathExist || !imgPathExist {
				listaMangas = append(listaMangas,
					model.Manga{Title: title, Path: "", SiteURL: readmBaseURL, ImgURL: ""})
				return
			}
			listaMangas = append(listaMangas,
				model.Manga{Title: title, Path: path, SiteURL: readmBaseURL, ImgURL: imgPath})
		},
	)

	return listaMangas
}
