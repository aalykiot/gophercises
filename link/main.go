package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type link struct {
	Href string
	Text string
}

// Parsing link's subtree to get the text
func parseLink(n *html.Node) string {
	// If the node is a comment return empty string
	if n.Type == html.CommentNode {
		return ""
	}

	// When on a leaf node return the text
	if n.FirstChild == nil {
		return strings.TrimSpace(n.Data)
	}

	var text string

	// Concatenating the text of the link
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += parseLink(c)
	}

	return text
}

// Parsing all the html tree using dfs
func parseHTML(n *html.Node, links *[]link) {
	// If we found a link element explore that subtree
	if n.Type == html.ElementNode && n.Data == "a" {
		var lnk link

		// Getting the href attribute
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				lnk.Href = attr.Val
			}
		}

		// Getting the text of the link
		lnk.Text = parseLink(n)

		*links = append(*links, lnk)
		return
	}

	// Continuing the dfs on the tree
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseHTML(c, links)
	}
}

func main() {
	file, err := os.Open("test.html")
	checkError(err)

	doc, err := html.Parse(file)
	checkError(err)

	var links []link
	parseHTML(doc, &links)

	fmt.Println(links)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
