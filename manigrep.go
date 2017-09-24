package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kieron-pivotal/manigrep/grepper"
)

func main() {
	filenamePtr := flag.String("file", "", "Input YAML file")
	valSearchPtr := flag.String("val-search", "", "Text to search for in values")
	flag.Parse()

	grepper, err := grepper.New(*filenamePtr)
	if err != nil {
		log.Fatal(err)
	}

	matches, err := grepper.SearchVal(*valSearchPtr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for match := range matches {
		fmt.Println(match)
	}
}
