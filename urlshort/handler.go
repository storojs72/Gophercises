package urlshort

import (
	"net/http"
	"strings"
	"fmt"
	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	result := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		if strings.EqualFold(r.URL.Path, "/") {
			fallback.ServeHTTP(w, r)
			return
		}
		http.Redirect(w, r, pathsToUrls[r.URL.Path], http.StatusMovedPermanently)
	}
	return result
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
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	return MapHandler(parsedYaml, fallback), nil
}

func parseYAML(yml []byte) (map[string]string, error) {
	result := make(map[string]string, 0)
	var configuration Config
	err := yaml.Unmarshal(yml, &configuration)
	if err != nil {
		return nil, err
	}
	for _, element := range configuration {
		result[element.Path] = element.Url
	}
	return result, nil
}

type Config []struct {
	Path string
	Url string
}

