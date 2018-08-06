package main

import (
	"fmt"
	"net/http"

	"github.com/gophercises/urlshort"
	"io/ioutil"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/one": "https://godoc.org/github.com/gophercises/urlshort",
		"/two":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /three
  url: https://github.com/gophercises/urlshort
- path: /four
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadFile("jsonInput.json")
	if err != nil {
		panic(err)
	}
	jsonHandler, err := urlshort.JSONHandler(bytes, yamlHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world! (from fallback)")
}