package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}
func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseJSON(jsonBytes)
	if err != nil {
		return nil, err
	}

	return MapHandler(buildMap(pathUrls), fallback), nil
}

func parseJSON(jsonBytes []byte) ([]pathUrls, error) {
	var pathUrls []pathUrls
	err := json.Unmarshal(jsonBytes, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYAML(yml)

	if err != nil {
		return nil, err
	}

	return MapHandler(buildMap(pathUrls), fallback), nil
}

func buildMap(urls []pathUrls) map[string]string {
	pathToUrls := make(map[string]string)
	for _, pu := range urls {
		pathToUrls[pu.Path] = pu.URL
	}

	fmt.Println(pathToUrls)
	return pathToUrls
}

func parseYAML(data []byte) ([]pathUrls, error) {
	var pathUrls []pathUrls
	err := yaml.Unmarshal(data, &pathUrls)

	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

type pathUrls struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
