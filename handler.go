package main

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		if dest, ok := pathToUrls[path]; ok {
			http.Redirect(res, req, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(res, req)
	}
}

func YAMLHandler(yamls []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yamls, &pathUrls)
	if err != nil {
		return nil, err
	}

	pathToUrls := make(map[string]string)

	for _, val := range pathUrls {
		pathToUrls[val.path] = val.url
	}

	return MapHandler(pathToUrls, fallback), nil
}

type pathUrl struct {
	path string `yaml:"path"`
	url  string `yaml:"url"`
}
