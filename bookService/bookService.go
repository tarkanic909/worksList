package bookService

import (
	"encoding/json"
	"worksList/getClient"
)

const bookUrl = "https://openlibrary.org/books/"

type BookOlid struct {
	Title   string `json:"title"`
	Authors []struct {
		Key string `json:"key"`
	} `json:"authors"`
	Type struct {
		Key string
	}
}

func GetBookByOLID(olid string) BookOlid {
	var response BookOlid

	body := getClient.Get(bookUrl + olid + ".json")

	json.Unmarshal([]byte(body), &response)

	return response
}
