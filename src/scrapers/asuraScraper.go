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

func loadasuraConfig() {
	asuraConfig = utils.ConfigLoader()[asuraConfigKey]
	log.Println("Asura config File Loaded!!")
}

func checkAsuraConfigFileIsLoaded() {
	if asuraConfig == nil {
		loadasuraConfig()
	}
}

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

func AsuraFindSerieUrlByName(name string) string {
	checkAsuraConfigFileIsLoaded()
	reqName := strings.ReplaceAll(name, " ", asuraConfig["separator"])
	rootNode := hw.GetUrl(fmt.Sprintf(asuraConfig["baseUrl"] + asuraConfig["searchUrl"] + reqName))
	node := hw.SearchNodeByAtrrFirstMatch(rootNode, "a", "title", name)
	return returnValFromKey(node, "href")
}

func AsuraGetLestEpUrl(url string) string {
	return returnValFromKey(hw.SearchForElementFirstMatch(hw.GetUrl(url), "div|class|eph-num", "a"), "href")
}

func AsuraGetLestEpNumber(url string) string {
	return hw.SearchForElementFirstMatch(hw.GetUrl(url), "div|class|eph-num", "span|class|chapternum").FirstChild.Data
}
