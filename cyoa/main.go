package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []option `json:"options"`
}

type option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type storyHandler struct {
	Chapters map[string]chapter
	Template *template.Template
}

func (s storyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	// If path is '/' redirect to '/intro'
	if path == "" {
		http.Redirect(w, r, "/intro", 301)
		return
	}

	chapter := s.Chapters[path]

	// Send the template to response writer
	s.Template.Execute(w, chapter)
}

func parseJSON(jsonFile []byte) (map[string]chapter, error) {
	var data map[string]chapter

	err := json.Unmarshal(jsonFile, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	// Reading the json file as []byte
	jsonFile, err := ioutil.ReadFile("gopher.json")
	checkError(err)

	// Decoding json data to map[string]chapter
	chapters, err := parseJSON(jsonFile)
	checkError(err)

	// Building the html template
	tmpl, err := template.ParseFiles("layout.html")
	checkError(err)

	// Create the http handler
	handler := storyHandler{
		Chapters: chapters,
		Template: tmpl,
	}

	// Start http server
	log.Println("Server started on port :8082")
	log.Fatal(http.ListenAndServe(":8082", handler))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
