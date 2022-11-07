package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"worksList/bookService"
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

func checkSortArgs(revArg *string, authArg *string) {
	if *revArg != "asc" && *revArg != "desc" || *authArg != "asc" && *authArg != "desc" {
		fmt.Println("Bad sort argument!", "Use asc or desc!")
		os.Exit(1)
	}
}

func main() {

	var authors []Author
	var isPrint string

	// take cli arguments
	olidArg := flag.String("olid", "OL32011221M", "Open Library ID")
	revArg := flag.String("revision", "asc", "sort by count of revision asc/desc")
	authArg := flag.String("author", "asc", "sort by count of revision asc/desc")

	flag.Parse()

	// check sorting arguments
	checkSortArgs(revArg, authArg)

	book := bookService.GetBookByOLID(strings.ToUpper(strings.TrimSpace(*olidArg)))

	if book.Title == "" {
		fmt.Println("No book found!")
		os.Exit(0)
	}

	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, book.Title)
	fmt.Printf("Title: %v ", colored)
	fmt.Println()

	// populate authors slice
	fmt.Println("Author(s):")
	for _, author := range book.Authors {
		works := worksService.GetAllWorks(author.Key)
		name := worksService.GetAuthorByKey(author.Key)
		fmt.Println(name, len(works), "work(s)")
		authors = append(authors, Author{Name: name, Works: works})
	}

	fmt.Print("Would you like to print list in stdout ? [Y/n] : ")
	fmt.Scanln(&isPrint)
	isPrint = strings.ToLower(strings.TrimSpace(isPrint))
	if isPrint == "" {
		isPrint = "y"
	}

	sortByName(authors, *authArg)
	sortByRevision(authors, *revArg)

	out, err := yaml.Marshal(authors)

	if err != nil {
		fmt.Println("Can not format output to yaml format!", err.Error())
		os.Exit(0)
	}

	if isPrint == "y" || isPrint == "yes" {
		fmt.Println(string(out))
	}
}
