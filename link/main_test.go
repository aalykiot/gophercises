package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

// Compairs two slices of type Link
func compareSlice(want, got []Link, t *testing.T) {
	if len(want) != len(got) {
		t.Errorf("Expected '%+v', got '%+v'", want, got)
	}

	for i, v := range want {
		if v != got[i] {
			t.Errorf("Expected '%+v', got '%+v'", v, got[i])
			break
		}
	}
}

func TestParseHtml(t *testing.T) {
	t.Run("parse an html anchor tag", func(t *testing.T) {
		str := `
			<a href="/dog">
				<span>Something in a span</span>
				Text not in a span
				<b>Bold text!</b>
			</a>
		`

		doc, _ := html.Parse(strings.NewReader(str))
		got := ParseHTML(doc)

		want := []Link{
			Link{
				Href: "/dog",
				Text: "Something in a span Text not in a span Bold text!",
			},
		}

		compareSlice(want, got, t)
	})

	t.Run("Parse an html document", func(t *testing.T) {
		str := `
			<html>
			<head>
			<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
			</head>
			<body>
			<h1>Social stuffs</h1>
			<div>
				<a href="https://www.twitter.com/joncalhoun">
				Check me out on twitter
				<i class="fa fa-twitter" aria-hidden="true"></i>
				</a>
				<a href="https://github.com/gophercises">
				Gophercises is on <strong>Github</strong>!
				</a>
			</div>
			</body>
			</html>
		`

		doc, _ := html.Parse(strings.NewReader(str))
		got := ParseHTML(doc)

		want := []Link{
			Link{
				Href: "https://www.twitter.com/joncalhoun",
				Text: "Check me out on twitter",
			},
			Link{
				Href: "https://github.com/gophercises",
				Text: "Gophercises is on Github!",
			},
		}

		compareSlice(want, got, t)
	})

	t.Run("parse an html document with comments", func(t *testing.T) {
		str := `
			<html>
			<body>
			<a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
			</body>
			</html>
		`

		doc, _ := html.Parse(strings.NewReader(str))
		got := ParseHTML(doc)

		want := []Link{
			Link{
				Href: "/dog-cat",
				Text: "dog cat",
			},
		}

		compareSlice(want, got, t)
	})
}
