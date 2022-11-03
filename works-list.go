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

func sortByName(arr []Author, sortType string) {
	sort.Slice(arr, func(i, j int) bool {
		if sortType == "desc" {
			return arr[i].Name > arr[j].Name
		}
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

func checkSortArgs(revArg *string, authArg *string) {
	if *revArg != "asc" && *revArg != "desc" || *authArg != "asc" && *authArg != "desc" {
		fmt.Println("Bad sort argument!", "Use asc or desc!")
		exit(1)
	}
}

func main() {

	var firstFound searchService.Doc
	var authors []Author
	var books []searchService.Doc
	var isPrint string

	// take cli arguments
	bookArg := flag.String("book", "Lord of the rings", "book name")
	revArg := flag.String("revision", "asc", "sort by count of revision asc/desc")
	authArg := flag.String("author", "asc", "sort by count of revision asc/desc")

	flag.Parse()

	// check sorting arguments
	checkSortArgs(revArg, authArg)

	// search for books
	books = searchService.SearchByTitle(*bookArg)

	if len(books) == 0 {
		fmt.Println("No book found!")
		exit(0)
	}

	// get first item
	firstFound = books[0]
	fmt.Print("First found: ")
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, firstFound.Title)
	fmt.Printf("Title: %v ", colored)
	fmt.Printf("Authors: %v \n", strings.Join(firstFound.AuthorsName, ", "))

	fmt.Print("Would you like to print list in stdout ? [Y/n] : ")
	fmt.Scanln(&isPrint)
	isPrint = strings.ToLower(strings.TrimSpace(isPrint))
	if isPrint == "" {
		isPrint = "y"
	}

	// populate authors slice
	for i, authorKey := range firstFound.AuthorsKey {
		works := worksService.GetWorks(authorKey)
		authors = append(authors, Author{Name: firstFound.AuthorsName[i], Works: works})
	}

	sortByName(authors, *authArg)
	sortByRevision(authors, *revArg)

	out, err := yaml.Marshal(authors)

	if err != nil {
		fmt.Println("Can not format output to yaml format!")
		exit(0)
	}

	if isPrint == "y" || isPrint == "yes" {
		fmt.Println(string(out))
	}
}
