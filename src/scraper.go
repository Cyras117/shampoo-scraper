package scraper

import (
	"fmt"
	"net/http"
	"shampoo-scraper/src/model"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var readmBaseURL = "https://www.readm.org"

//Returns the last n(MAX: 5) pages of updates from the site.
func GetLatestUpdatesPage(n int) []model.Manga {
	var listaMangas []model.Manga
	qtdPages := n + 1

	if n > 5 || n <= 0 {
		qtdPages = 5
	}

	for i := 1; i < qtdPages; i++ {
		latestUpdatePageURL := readmBaseURL + fmt.Sprintf("/latest-releases/%d", i)

		res, err := http.Get(latestUpdatePageURL)
		errLogOutput(err)

		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)
		errLogOutput(err)

		doc.Find("ul.clearfix.latest-updates li.segment-poster-sm").Each(
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

//Returns a list with all the mangas on the site.
func GetAllMangas() []model.Manga {
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
		errLogOutput(err)

		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)
		errLogOutput(err)

		doc.Find("li.segment-poster-sm").Each(
			func(i int, s *goquery.Selection) {
				node := s.Find("div.poster")
				path, _ := node.Find("a").Attr("href")
				title := node.Find("h2.truncate").Text()
				imgPath, _ := node.Find("img.lazy-wide").Attr("data-src")

				listaMangas = append(listaMangas,
					model.Manga{Title: title, Path: path, SiteURL: readmBaseURL, ImgURL: imgPath})
			},
		)
	}
	return listaMangas
}

//Returns a list of mangas with the passed phrase in the title.
func SearchManga(title string) []model.Manga {
	var searchResults []model.Manga
	mangas := GetAllMangas()

	for _, m := range mangas {
		if isIn(title, m.Title) {
			searchResults = append(searchResults, m)
		}
	}
	return searchResults
}

//TODO:Optmize it to go straight to the manga page
func SearchMangaWithPath(path string) model.Manga {
	var res model.Manga
	mangas := GetAllMangas() // takes too long
	//TODO:Treat in case nothing
	for _, m := range mangas {
		if isIn(path, m.Path) {
			res = m
		}
	}

	return res
}

//Returns a list of mangas with the last released chapter.
func GetLastChMangasList(list []model.Manga) []model.Manga {
	var resList []model.Manga
	for _, manga := range list {

		resList = append(resList, GetLastChManga(manga))
	}

	return resList
}

//Returns a manga with the last released chapter.
func GetLastChManga(manga model.Manga) model.Manga {

	fullURL := manga.SiteURL + manga.Path

	res, err := http.Get(fullURL)
	errLogOutput(err)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	errLogOutput(err)
	lastep := doc.Find(".episodes-box a[data-tab]").First().Text()

	lastep = strings.ReplaceAll(lastep, "CH", "")
	lastep = strings.ReplaceAll(lastep, " ", "")
	lastep = strings.Split(lastep, "-")[0]

	if lastep == "" {
		lastep = "0"
	}

	lep, err := strconv.ParseFloat(lastep, 64)
	errLogOutput(err)

	return model.Manga{
		Title: manga.Title, Path: manga.Path, SiteURL: manga.SiteURL,
		ImgURL: manga.ImgURL, CurrentCh: lep, LastReadCh: manga.LastReadCh}
}
