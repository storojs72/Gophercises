package urlshort

import (
	"net/http"
	"gopkg.in/yaml.v2"
	"encoding/json"
	"github.com/boltdb/bolt"
	"time"
	"fmt"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if pathsToUrls[r.URL.Path] == ""{
			fallback.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, pathsToUrls[r.URL.Path], http.StatusTemporaryRedirect)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var configuration []struct {
		Path string
		Url string
	}
	err := yaml.Unmarshal(yml, &configuration)
	if err != nil {
		return nil, err
	}

	pathsMap := make(map[string]string, 0)
	for _, element := range configuration {
		pathsMap[element.Path] = element.Url
	}

	return MapHandler(pathsMap, fallback), nil
}

func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var configuration struct {
		RedirectInfo []struct{
			Path string
			Url string
		}
	}
	err := json.Unmarshal(jsn, &configuration)
	if err != nil {
		return nil, err
	}

	pathsMap := make(map[string]string, 0)
	for _, element := range configuration.RedirectInfo {
		pathsMap[element.Path] = element.Url
	}

	return MapHandler(pathsMap, fallback), nil
}

func BoltDBHandler(pathToBoltDB string, fallback http.Handler) (http.HandlerFunc, error) {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := bolt.Open(pathToBoltDB, 0600, &bolt.Options{Timeout:1 * time.Second})
		if err != nil {
			panic(err)
		}
		defer db.Close()

		var retreivedUrl []byte
		err = db.Update(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte("URLs_and_Paths"))
			retreivedUrl = bucket.Get([]byte(r.URL.Path))
			return nil
		}); if err != nil {
			fmt.Println(err)
			fallback.ServeHTTP(w, r)
		}

		if retreivedUrl == nil {
			fallback.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, string(retreivedUrl), http.StatusTemporaryRedirect)
		}

	}, nil
}