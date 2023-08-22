package htmlwrapper

import (
	"net/http"
	"shampoo-scraper/src/utils"

	"golang.org/x/net/html"
)

func GetUrl(url string) *html.Node {
	res, err := http.Get(url)
	if err != nil {
		utils.ErrLogOutput(err)
	}
	defer res.Body.Close()
	node, err := html.Parse(res.Body)
	if err != nil {
		utils.ErrLogOutput(err)
	}
	return node
}
