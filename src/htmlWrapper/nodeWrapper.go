package htmlwrapper

import (
	"shampoo-scraper/src/utils"
	"strings"

	"golang.org/x/net/html"
)

// Check the [node] contains the same data attribute and value
func checkMatch(node *html.Node, data, atrr, value string) bool {
	if node.Attr == nil {
		return false
	}
	if !strings.EqualFold(strings.ToLower(node.Data), strings.ToLower(data)) {
		return false
	}
	for i := 0; i < len(node.Attr); i++ {
		natrr := strings.ToLower(node.Attr[i].Key)
		nvalue := strings.ToLower(node.Attr[i].Val)
		atr := strings.ToLower(atrr)
		val := strings.ToLower(value)
		if natrr == atr {
			if utils.IsIn(val, nvalue) {
				return true
			}
		}

	}
	return false
}

/*
Searches for a  node inside another node

elements ex: "data|attr|value","data"
*/
func SearchForElementFirstMatch(node *html.Node, elements ...string) *html.Node {
	if elements == nil {
		return node
	}
	auxNode := node
	for _, element := range elements {
		elementValues := strings.Split(element, "|")
		if len(elementValues) < 2 {
			auxNode = SearchNodeByDataFirstMatch(auxNode, elementValues[0])
		} else {
			auxNode = SearchNodeByAtrrFirstMatch(auxNode, elementValues[0], elementValues[1], elementValues[2])
		}
	}
	return auxNode
}

/*
Searches for the first match of a node with the provided data,
if not found it returns nill
*/
func SearchNodeByDataFirstMatch(node *html.Node, data string) *html.Node {
	nodeResult := node

	if nodeResult.Data == data {
		return nodeResult
	}

	if node.FirstChild != nil {
		nodeResult = SearchNodeByDataFirstMatch(node.FirstChild, data)
		if nodeResult != nil {
			return nodeResult
		}
	}

	if node.NextSibling != nil {
		nodeResult = SearchNodeByDataFirstMatch(node.NextSibling, data)
		if nodeResult != nil {
			return nodeResult
		}
	}

	return nil
}

/*
Searches for nodes with the provided data, attribute and value provided
*/
func SearchNodesByAtrr(node *html.Node, data, atrr, value string, result *[]html.Node) {
	if checkMatch(node, atrr, value, data) {
		*result = append(*result, *node)
	}

	if node.FirstChild != nil {
		SearchNodesByAtrr(node.FirstChild, data, atrr, value, result)
	}

	if node.NextSibling != nil {
		SearchNodesByAtrr(node.NextSibling, data, atrr, value, result)
	}
}

/*
Searches for nodes with the provided data, attribute and value provided
if not found it returns nill
*/
func SearchNodeByAtrrFirstMatch(node *html.Node, data, atrr, value string) *html.Node {
	nodeResult := node

	if checkMatch(node, data, atrr, value) {
		return nodeResult
	}

	if node.FirstChild != nil {
		nodeResult = SearchNodeByAtrrFirstMatch(node.FirstChild, data, atrr, value)
		if nodeResult != nil {
			return nodeResult
		}
	}

	if node.NextSibling != nil {
		nodeResult = SearchNodeByAtrrFirstMatch(node.NextSibling, data, atrr, value)
		if nodeResult != nil {
			return nodeResult
		}
	}
	return nil
}
