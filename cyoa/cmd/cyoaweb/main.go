package main

import (
	"flag"
	"fmt"
	"github.com/hinaka/gophercises/cyoa"
	"os"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON filename with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	fmt.Printf("%+v\n", story)
}
