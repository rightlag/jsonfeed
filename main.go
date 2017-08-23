package main

import (
	"flag"
	"io/ioutil"
	"log"
)

var (
	file string
)

func init() {
	flag.StringVar(&file, "file", "", "")
}

func main() {
	flag.Parse()
	if file == "" {
		log.Fatal("main: argument -file is required")
	}
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	var parser FeedParser
	root, err := parser.Parse(b)
	if err != nil {
		log.Fatal(err)
	}
	visitor := NewValidationVisitor()
	root.Accept(visitor)
	if visitor.HasErrors() {
		for _, err := range visitor.Errors() {
			log.Println(err)
		}
	}
}
