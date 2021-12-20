package main

import (
	"flag"
	"fmt"
	"github.com/hinaka/gophercises/urlshort"
	"net/http"
	"os"
)

func main() {
	yamlFilename := flag.String("yaml", "", "a yaml file that contains shortened Path and its original URL")
	jsonFilename := flag.String("json", "", "a json file that contains shortened Path and its original URL")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	if *yamlFilename != "" {
		yamlData := readFile(*yamlFilename)
		yamlHandler, err := urlshort.YAMLHandler(yamlData, mapHandler)
		if err != nil {
			panic(err)
		}
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", yamlHandler)
	} else if *jsonFilename != "" {
		jsonData := readFile(*jsonFilename)
		jsonHandler, err := urlshort.JSONHandler(jsonData, mapHandler)
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
	fmt.Fprintln(w, "Hello, world!")
}

func readFile(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		exit(fmt.Sprintf("Failed to read file: %s\n", filename))
	}
	return data
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}