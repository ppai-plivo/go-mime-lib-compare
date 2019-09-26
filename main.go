package main

import (
	"flag"
	"log"
)

func main() {

	var testDir string
	flag.StringVar(&testDir, "addr", "testdata", "Path to directory containing test files")
	flag.Parse()

	if err := printContentTypes(testDir); err != nil {
		log.Fatalf("printing content types failed: %s\n", err.Error())
	}
}
