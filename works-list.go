package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	exitCode := 0
	bookArg := flag.String("book", "Lord of the rings", "book name")
	sortArg := flag.String("sort", "asc", "sort by count of revision asc/desc")

	defer func() {
		os.Exit(exitCode)
	}()

	flag.Parse()

	if *sortArg != "asc" && *sortArg != "desc" {
		fmt.Println("Bad sort argument!", "Use asc or desc!")
		exitCode = 1
	}

	fmt.Println("book: ", *bookArg)
	fmt.Println("sort: ", *sortArg)
}
