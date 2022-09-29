package scraper

import (
	"fmt"
	"net/http"
	"shampoo-scraper/src/model"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const readmBaseURL = "https://www.readm.org"

// Returns the last n(MAX: 5) pages of updates from the site.
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

// Returns a list with all the mangas on the site.
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

// Returns a list of mangas with the passed phrase in the title.
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

// Grab information of that manga page
func GetMangaWithPath(path string) model.Manga {
	var mangaRsult model.Manga

	pathResult := strings.ReplaceAll(path, "/", "")

	mangaURL := fmt.Sprint(readmBaseURL, "/manga/", pathResult)

	res, err := http.Get(mangaURL)
	errLogOutput(err)

	//TODO add treatment in case manga does not exist

	doc, err := goquery.NewDocumentFromReader(res.Body)
	errLogOutput(err)

	defer res.Body.Close()

	doc.Find("#router-view").Each(
		func(i int, s *goquery.Selection) {
			title := s.Find("#router-view > div > div.ui.grid > div.left.floated.sixteen.wide.tablet.eight.wide.computer.column > a > h1").Text()
			imgPath, _ := s.Find("#series-profile-image-wrapper > img").Attr("src")
			lastCh := s.Find("#seasons-menu > div > a.item.active").Text()
			//Qtd chapters
			//totalChs := s.Find("#series-profile-content-wrapper > article > div.media-meta > table > tbody > tr > td:nth-child(2) > div:nth-child(2)")

			lastCh = strings.ReplaceAll(lastCh, "CH", "")
			lastCh = strings.ReplaceAll(lastCh, " ", "")
			lastCh = strings.Split(lastCh, "-")[0]
			if lastCh == "" {
				lastCh = "0"
			}

			lastChRes, err := strconv.ParseFloat(lastCh, 64)
			errLogOutput(err)

			mangaRsult.Path = path
			mangaRsult.Title = title
			mangaRsult.ImgURL = imgPath
			mangaRsult.SiteURL = readmBaseURL
			mangaRsult.CurrentCh = lastChRes
		},
	)

	return mangaRsult
}

// Returns a list of mangas with the last released chapter.
func GetLastChMangasList(list []model.Manga) []model.Manga {
	var resList []model.Manga
	for _, manga := range list {

		resList = append(resList, GetMangaWithPath(manga.Path))
	}

	return resList
}
