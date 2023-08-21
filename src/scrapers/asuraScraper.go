package scrapers

import (
	"log"
	"net/http"
	nodeWrapper "shampoo-scraper/src/htmlWrapper"
	"shampoo-scraper/src/utils"

	"golang.org/x/net/html"
)

const asuraConfigKey = "asuraConfig"

var asuraConfig interface{}

func loadasuraConfig() {
	readmConfig = utils.ConfigLoader()[asuraConfigKey]
	log.Println("Asura config File Loaded!!\n")
}

func checkAsuraConfigFileIsLoaded() {
	if readmConfig == nil {
		loadasuraConfig()
	}
	return
}

func AsuraGetLestEpUrl(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		utils.ErrLogOutput(err)
	}
	defer resp.Body.Close()
	rootNode, err := html.Parse(resp.Body)
	if err != nil {
		utils.ErrLogOutput(err)
	}
	n := nodeWrapper.SearchFirstNodeOccurrence(rootNode, "class", "eph-num", "div")
	return n.FirstChild.NextSibling.Attr[0].Val
}
