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

		var got []Link
		doc, _ := html.Parse(strings.NewReader(str))
		ParseHTML(doc, &got)

		want := []Link{
			Link{
				href: "/dog",
				text: "Something in a span Text not in a span Bold text!",
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

		var got []Link
		doc, _ := html.Parse(strings.NewReader(str))
		ParseHTML(doc, &got)

		want := []Link{
			Link{
				href: "https://www.twitter.com/joncalhoun",
				text: "Check me out on twitter",
			},
			Link{
				href: "https://github.com/gophercises",
				text: "Gophercises is on Github!",
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

		var got []Link
		doc, _ := html.Parse(strings.NewReader(str))
		ParseHTML(doc, &got)

		want := []Link{
			Link{
				href: "/dog-cat",
				text: "dog cat",
			},
		}

		compareSlice(want, got, t)
	})
}
