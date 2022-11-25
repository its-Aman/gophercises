package main

import (
	"flag"
	"fmt"
	"net/http"

	"url_shortner/urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlFlag := flag.String("yaml", yaml, "The yaml file containing urls as path and url format")
	yamlHandler, err := urlshort.YAMLHandler([]byte(*yamlFlag), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")

	json := `
	[
		{
			"path": "/ggl",
			"url" : "https://www.google.com/"
		}
	]
	`
	jsonFlag := flag.String("json", json, "JSON file containing urls as path and url format")

	jsonHandler, err := urlshort.JSONHandler([]byte(*jsonFlag), yamlHandler)

	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
