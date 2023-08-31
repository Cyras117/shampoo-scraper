package scrapers

import (
	"fmt"
	"log"
	hw "shampoo-scraper/src/htmlWrapper"
	"shampoo-scraper/src/model"
	"shampoo-scraper/src/utils"
	"strings"

	"golang.org/x/net/html"
)

const asuraConfigKey = "asuraConfig"

var asuraConfig model.StiteConfig

/*
Load Config file for asura
*/
func loadasuraConfig() {
	asuraConfig = utils.ConfigLoader()[asuraConfigKey]
	log.Println("Asura config File Loaded!!")
}

/*
Verify if config file is loaded already
*/
func checkAsuraConfigFileIsLoaded() {
	if asuraConfig == nil {
		loadasuraConfig()
	} else {
		fmt.Println("Asura config File already Loaded!!")
	}
}

/*
Returs the value from a html node attribute.
*/
func returnValFromKey(node *html.Node, key string) string {
	var actualSeriesUrl string

	for _, a := range node.Attr {
		if a.Key == key {
			{
				actualSeriesUrl = a.Val
				break
			}
		}
	}
	return actualSeriesUrl
}

/*
Gets the URL from a serie by its name on asura site
*/
func AsuraFindSerieUrlByName(name string) string {
	checkAsuraConfigFileIsLoaded()
	reqName := strings.ReplaceAll(name, " ", asuraConfig["separator"])
	rootNode := hw.GetUrl(fmt.Sprintf(asuraConfig["baseUrl"] + asuraConfig["searchUrl"] + reqName))
	node := hw.SearchNodeByAtrrFirstMatch(rootNode, "a", "title", name)
	return returnValFromKey(node, "href")
}

/*
Gets the URL from the last charpter of a serie on asura site
*/
func AsuraGetLastChByUrl(urlSerie string) string {
	return returnValFromKey(hw.SearchForElementFirstMatch(hw.GetUrl(urlSerie), "div|class|eph-num", "a"), "href")
}

/*
Gets the URL from the n charpter of a serie on asura site1
*/
func AsuraGetChUrlByNameNumber(name string, number int) string {
	return returnValFromKey(hw.SearchForElementFirstMatch(hw.GetUrl(AsuraFindSerieUrlByName(name)), fmt.Sprintf("li|data-num|%d", number), "a"), "href")
}

/*
Gets the URL from the first charpter of a serie on asura site
*/
func AsuraGetFirstChByUrl(urlSerie string) string {
	return returnValFromKey(hw.SearchForElementFirstMatch(hw.GetUrl(urlSerie), "li|data-num|1", "a"), "href")
}

/*
Gets the URL from the last charpter of a serie on asura site through its name
*/
func AsuraGetLastChByName(nameSerie string) string {
	return returnValFromKey(hw.SearchForElementFirstMatch(hw.GetUrl(AsuraFindSerieUrlByName(nameSerie)), "div|class|eph-num", "a"), "href")
}

/*
Gets the URL from the First charpter of a serie on asura site through its name
*/
func AsuraGetFirstChByName(nameSerie string) string {
	return returnValFromKey(hw.SearchForElementFirstMatch(hw.GetUrl(AsuraFindSerieUrlByName(nameSerie)), "li|data-num|1", "a"), "href")
}

/*
Gets the URL from the n charpter of a serie on asura site through its name
*/
func AsuraGetLestChNumber(urlSerie string) string {
	return hw.SearchForElementFirstMatch(hw.GetUrl(urlSerie), "div|class|eph-num", "span|class|chapternum").FirstChild.Data
}
