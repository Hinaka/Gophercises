package main

import (
	"flag"
	"fmt"
	"github.com/hinaka/gophercises/urlshort"
	"net/http"
	"os"
)

func main() {
	yamlFilename := flag.String("yaml", "", "a yaml file that contains shortened URL and its original Path")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	if *yamlFilename != "" {
		file, err := os.ReadFile(*yamlFilename)
		if err != nil {
			exit(fmt.Sprintf("Failed to read yaml file: %s\n", *yamlFilename))
		}
		yamlHandler, err := urlshort.YAMLHandler(file, mapHandler)
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", yamlHandler)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}