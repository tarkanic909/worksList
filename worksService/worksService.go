package worksService

import (
	"encoding/json"
	"worksList/getClient"
)

const url = "https://openlibrary.org/authors/"

type Work struct {
	Title    string
	Revision int
}

type WorksResponse struct {
	Entries []Work
}

func GetWorks(authorKey string) []Work {
	var response WorksResponse

	body := getClient.Get(url + authorKey + "/works.json")

	json.Unmarshal([]byte(body), &response)

	return response.Entries

}
