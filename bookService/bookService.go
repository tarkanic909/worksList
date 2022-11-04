package bookService

import (
	"encoding/json"
	"fmt"
	"worksList/getClient"
)

const bookUrl = "https://openlibrary.org/api/books?bibkeys=ISBN:"

type bookResponse struct {
	authors []struct {
		name string
	}
}

func getBookByISBN(isbn string) {
	var response bookResponse
	fmt.Println("hello from book service")

	body := getClient.Get(bookUrl + isbn + "&format=json&jscmd=data")

	json.Unmarshal([]byte(body), &response)

	// return response

}
