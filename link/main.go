package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func parseLink(n *html.Node) string {
	if n.FirstChild == nil {
		return n.Data
	}

	var text string

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += parseLink(c)
	}

	return text
}

func parseHTML(n *html.Node, links *[]Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		var link Link
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				link.Href = attr.Val
			}
		}
		link.Text = parseLink(n)

		*links = append(*links, link)
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseHTML(c, links)
	}
}

func main() {
	file, err := os.Open("test.html")
	checkError(err)

	doc, err := html.Parse(file)
	checkError(err)

	var links []Link
	parseHTML(doc, &links)

	fmt.Println(links)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
