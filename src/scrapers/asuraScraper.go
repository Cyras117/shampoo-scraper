package scrapers

import (
	"fmt"
	"log"
	"net/http"
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

func AsuraFindSerieUrlByName(name string) string {
	checkAsuraConfigFileIsLoaded()
	reqName := strings.ReplaceAll(name, " ", asuraConfig["separator"])
	rootNode := hw.GetUrl(fmt.Sprintf(asuraConfig["baseUrl"] + asuraConfig["searchUrl"] + reqName))
	node := hw.SearchFirstNodeOccurrence(rootNode, "title", name, "a")
	var actualSeriesUrl string

	for _, a := range node.Attr {
		if a.Key == "href" {
			{
				actualSeriesUrl = a.Val
			}
		}
	}
	return actualSeriesUrl
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
	n := hw.SearchFirstNodeOccurrence(rootNode, "class", "eph-num", "div")
	return n.FirstChild.NextSibling.Attr[0].Val
}
