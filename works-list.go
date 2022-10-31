package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
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
			if sortType == "desc" {
				return author.Works[i].Revision > author.Works[j].Revision
			}
			return author.Works[i].Revision < author.Works[j].Revision
		})
	}
}

func exit(code int) {
	os.Exit(code)
}

func main() {

	var firstFound searchService.Doc
	var authors []Author
	var books []searchService.Doc
	var isPrint string

	// take cli arguments
	bookArg := flag.String("book", "Lord of the rings", "book name")
	sortArg := flag.String("sort", "asc", "sort by count of revision asc/desc")

	flag.Parse()

	// check sort argument
	if *sortArg != "asc" && *sortArg != "desc" {
		fmt.Println("Bad sort argument!", "Use asc or desc!")
		exit(1)
	}

	// search for books
	books = searchService.Search(*bookArg)

	if len(books) == 0 {
		fmt.Println("No book found!")
		exit(1)
	}

	// get first item
	firstFound = books[0]
	fmt.Print("First found: ")
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, firstFound.Title)
	fmt.Printf("Title: %v ", colored)
	fmt.Printf("Authors: %v \n", strings.Join(firstFound.AuthorsName, ", "))

	fmt.Print("Would you like to print list in stdout ? [Y/n] : ")
	fmt.Scanln(&isPrint)
	if strings.TrimSpace(isPrint) == "" {
		isPrint = "y"
	}

	// populate authors slice
	for i, authorKey := range firstFound.AuthorsKey {
		works := worksService.GetWorks(authorKey)
		authors = append(authors, Author{Name: firstFound.AuthorsName[i], Works: works})
	}

	sortByName(authors)
	sortByRevision(authors, *sortArg)

	out, _ := yaml.Marshal(authors)

	if isPrint == "y" || isPrint == "yes" {
		fmt.Println(string(out))
	}
}
