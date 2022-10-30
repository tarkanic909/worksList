package main

import (
	"flag"
	"fmt"
)

func main() {
	bookArg := flag.String("book", "Lord of the rings", "book name")
	authorArg := flag.String("author", "asc", "sort by author name asc/desc")
	revisionArg := flag.String("revision", "asc", "sort by count of revision asc/desc")

	flag.Parse()

	fmt.Println("book: ", *bookArg)
	fmt.Println("author: ", *authorArg)
	fmt.Println("revision: ", *revisionArg)
}
