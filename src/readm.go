package scraper

import (
	"fmt"
	"log"
	"net/http"
	"shampoo-scraper/src/model"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var readmBaseURL = "https://www.readm.org"

//Gets the last 5 pages of updates from the site and returns
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

//Gets all the mangas on the web site.
func getAllMangasOnreadm() []model.Manga {
	stringBase := "#ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var listaMangas []model.Manga

	for _, l := range stringBase {
		letter := string(l)
		var AllMangasPageURL string

		if letter == "#" {
			AllMangasPageURL = fmt.Sprintf(readmBaseURL + "/manga-list/")
		} else {
			AllMangasPageURL = fmt.Sprintf(readmBaseURL+"/manga-list/%s", letter)
		}

		res, err := http.Get(AllMangasPageURL)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer res.Body.Close()
		doc, err := goquery.NewDocumentFromReader(res.Body)

		//TODO:find a way to not break the application
		if err != nil {
			log.Fatal(err.Error())
		}

		doc.Find("li.segment-poster-sm").Each(
			func(i int, s *goquery.Selection) {
				node := s.Find("div.poster")
				path, _ := node.Find("a").Attr("href")
				title := node.Find("h2.truncate").Text()
				//status := node.Find("p.poster-meta").Text()//TODO:Checke if it will be needed.
				imgPath, _ := node.Find("img.lazy-wide").Attr("data-src")

				listaMangas = append(listaMangas,
					model.Manga{Title: title, Path: path, SiteURL: readmBaseURL, ImgURL: imgPath})
			},
		)
	}
	return listaMangas
}

//check if a string is inside another
func isIn(phrase string, str string) bool {

	p := strings.ToLower(phrase)
	s := strings.ToLower(str)

	lPhr := len(phrase)
	lStr := len(str)

	if lStr < lPhr {
		return false
	}

	for i := 0; i < lStr; i++ {
		if p[0] == s[i] {
			for j := i; j < lStr; j++ {
				if s[j] == p[j-i] {
					if j-i == lPhr-1 {
						return true
					}
				} else {
					break
				}
			}
		}
	}
	return false
}

//Searches on the readm website for a manga with the specified phrase.
func SearchMangareadm(title string) []model.Manga {
	var searchResults []model.Manga
	mangas := getAllMangasOnreadm()

	for _, m := range mangas {
		if isIn(title, m.Title) {
			searchResults = append(searchResults, m)
		}
	}
	return searchResults
}
