package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Link describes the format of an anchor tag
type Link struct {
	href string
	text string
}

// Parsing link's subtree to get the text
func parseLink(n *html.Node) string {
	// If the node is a comment return empty string
	if n.Type == html.CommentNode {
		return ""
	}

	// When on a text node return the text
	if n.Type == html.TextNode {
		return n.Data
	}

	var text string

	// Concatenating the text of the link
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += parseLink(c)
	}

	return strings.Join(strings.Fields(text), " ")
}

// ParseHTML will parse all the html links to an array using dfs
func ParseHTML(n *html.Node, links *[]Link) {
	// If we found a link element explore that subtree
	if n.Type == html.ElementNode && n.Data == "a" {
		var link Link

		// Getting the href attribute
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				link.href = attr.Val
			}
		}

		// Getting the text of the link
		link.text = parseLink(n)

		*links = append(*links, link)
		return
	}

	// Continuing the dfs on the tree
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ParseHTML(c, links)
	}
}

func main() {
	file, err := os.Open("test.html")
	checkError(err)

	doc, err := html.Parse(file)
	checkError(err)

	var links []Link
	ParseHTML(doc, &links)

	fmt.Println(links)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
