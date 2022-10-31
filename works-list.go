package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"worksList/searchService"
	"worksList/worksService"

	"gopkg.in/yaml.v3"
)

type Author struct {
	Name  string
	Works []worksService.Work
}

func sortByName(arr []Author) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Name < arr[j].Name
	})
}

func sortByRevision(arr []Author, sortType string) {
	for _, author := range arr {
		sort.Slice(author.Works, func(i, j int) bool {
			return author.Works[i].Revision < author.Works[j].Revision
		})
	}
}

func main() {
	exitCode := 0
	var firstFound searchService.Doc
	var authors []Author
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

	books := searchService.Search(*bookArg)

	if len(books) == 0 {
		fmt.Println("No book found!")
		exitCode = 1
	}

	firstFound = books[0]
	for i, authorKey := range firstFound.AuthorsKey {
		works := worksService.GetWorks(authorKey)
		authors = append(authors, Author{Name: firstFound.AuthorsName[i], Works: works})
	}

	sortByName(authors)
	sortByRevision(authors, *sortArg)

	out, _ := yaml.Marshal(authors)

	fmt.Println(string(out))
}
