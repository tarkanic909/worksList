package bookService

import (
	"encoding/json"
	"fmt"
	"worksList/getClient"
)

const bookUrl = "https://openlibrary.org/books/"

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

type BookOlid struct {
	Title   string `json:"title"`
	Authors []struct {
		Key string `json:"key"`
	} `json:"authors"`
}

func GetBookByOLID(olid string) Book {
	var response BookOlid

	body := getClient.Get(bookUrl + olid + ".json")
	fmt.Println(string(body))

	json.Unmarshal([]byte(body), &response)
	fmt.Println(response)

	return Book{Title: response.Title}
}
