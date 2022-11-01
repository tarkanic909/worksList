package bookservice

import (
	"encoding/json"
	"errors"
	"worksList/getClient"
)

const bookUrl = "https://openlibrary.org/books/"

type BookResponse struct {
	Error string
}

func GetById(id string) (BookResponse, error) {
	var response BookResponse

	body := getClient.Get(bookUrl + id + ".json")
	json.Unmarshal([]byte(body), &response)

	if response.Error == "notfound" {
		return BookResponse{}, errors.New("not found")
	}

	return response, nil

}
