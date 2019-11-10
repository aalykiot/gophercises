package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type route struct {
	Path string
	URL  string
}

func mapHandler(urlPaths map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if url, ok := urlPaths[r.URL.Path]; ok {
			http.Redirect(w, r, url, 301)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

func buildMap(routes []route) map[string]string {
	urlPaths := make(map[string]string)

	for _, route := range routes {
		urlPaths[route.Path] = route.URL
	}

	return urlPaths
}

func yamlHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var data []route

	err := yaml.Unmarshal(yml, &data)
	if err != nil {
		return nil, err
	}

	urlPaths := buildMap(data)

	mapHandler := mapHandler(urlPaths, fallback)

	return mapHandler, nil
}

func newDefaultMux() http.Handler {
	mux := http.NewServeMux()
	muxFunc := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "Go server works!")
	}
	mux.HandleFunc("/", muxFunc)
	return mux
}

func main() {
	yamlFilename := flag.String("yaml", "urls.yml", "The filename of the yml file with the routes")
	flag.Parse()

	mux := newDefaultMux()

	yamlFile, err := ioutil.ReadFile(*yamlFilename)
	if err != nil {
		panic(err)
	}

	mapHandler, err := yamlHandler(yamlFile, mux)
	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(":8081", mapHandler))
}
