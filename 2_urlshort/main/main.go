package main

import (
	"fmt"
	"net/http"
	"github.com/gophercises/2_urlshort"
	"flag"
	"io/ioutil"
)

func main() {

	var useBoltDb bool
	flag.BoolVar(&useBoltDb, "boltdb", false, "Set to TRUE if use BoltDB")
	flag.Parse()

	mux := defaultMux()
	if useBoltDb {
		boltDbHandler, err := __urlshort.BoltDBHandler("2_urlshort/inputs/bolt.db", mux)
		if err != nil {
			panic(err)
		}

		fmt.Println("Starting the server on :8080 (boltdb)")
		http.ListenAndServe(":8080", boltDbHandler)
	} else {
		//Build the MapHandler using the mux as the fallback
		pathsToUrls := map[string]string{
			"/one": "https://godoc.org/github.com/gophercises/2_urlshort",
			"/two": "https://godoc.org/gopkg.in/yaml.v2",
		}
		mapHandler := __urlshort.MapHandler(pathsToUrls, mux)

		// Build the YAMLHandler using the mapHandler as the
		// fallback
		yamlBytes, err := ioutil.ReadFile("2_urlshort/inputs/yamlInput.yaml")
		if err != nil {
			panic(err)
		}

		yamlHandler, err := __urlshort.YAMLHandler(yamlBytes, mapHandler)
		if err != nil {
			panic(err)
		}

		jsonBytes, err := ioutil.ReadFile("2_urlshort/inputs/jsonInput.json")
		if err != nil {
			panic(err)
		}
		jsonHandler, err := __urlshort.JSONHandler(jsonBytes, yamlHandler)
		if err != nil {
			panic(err)
		}
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", jsonHandler)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world! (from fallback)")
}