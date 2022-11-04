package bookService

import (
	"encoding/json"
	"regexp"
	"worksList/getClient"
)

const bookUrl = "https://openlibrary.org/api/books?bibkeys=ISBN:"

type BookResponse struct {
	Book Book
}

type Book struct {
	Title   string `json:"title"`
	Authors []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"authors"`
}

func GetBookByISBN(isbn string) Book {
	var response BookResponse
	// r := regexp.MustCompile(`\w+:[0-9]\d{1,13}`)
	r := regexp.MustCompile(`\bISBN:\d{1,13}`)

	body := getClient.Get(bookUrl + isbn + "&format=json&jscmd=data")

	// replace ISBN:number with book
	result := r.ReplaceAllString(string(body), "Book")
	json.Unmarshal([]byte(result), &response)

	return response.Book

}
