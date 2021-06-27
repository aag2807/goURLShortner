package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	maphandler := MapHandler(pathsToUrls, mux)

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := YAMLHandler([]byte(yaml), maphandler)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	_ = yamlHandler
	fmt.Println("starting 8080")
	http.ListenAndServe(":8080", maphandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello humans")
}
