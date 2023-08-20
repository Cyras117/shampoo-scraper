package scrapers

import (
	"net/http"
	nodeWrapper "shampoo-scraper/src/htmlWrapper"
	"shampoo-scraper/src/utils"

	"golang.org/x/net/html"
)

// criar uma funçao para carregar isso do yaml
const baseURL string = "https://asura.nacm.xyz/"

//https://asura.nacm.xyz/?s=The+extra

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
