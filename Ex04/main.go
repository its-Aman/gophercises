package main

import (
	"HTML_Link_Parser/HTMLparser"
	"flag"
	"log"
	"os"
)

func main() {
	htmlFile := flag.String("HTML file name", "ex01.html", "HTML File to parse")
	flag.Parse()

	file, err := os.Open(*htmlFile)

	if err != nil {
		log.Fatal("Error while parsing the file", err)
		panic(err)
	}

	defer file.Close()
	HTMLparser.Parse(file)
}
