package worksService

import (
	"encoding/json"
	"worksList/getClient"
)

const url = "https://openlibrary.org/"

type Work struct {
	Title    string
	Revision int
}

type WorksResponse struct {
	Entries []Work
	Links   struct {
		Next string
	}
}

func getWorks(next string, allWorks *[]Work) {
	var response WorksResponse

	if next != "" {
		body := getClient.Get(url + next)
		json.Unmarshal([]byte(body), &response)
		*allWorks = append(*allWorks, response.Entries...)
		getWorks(response.Links.Next, allWorks)
	}

}

func GetAllWorks(authorKey string) []Work {
	var response WorksResponse
	var allWorks []Work

	body := getClient.Get(url + authorKey + "/works.json")

	json.Unmarshal([]byte(body), &response)
	allWorks = append(allWorks, response.Entries...)

	if response.Links.Next != "" {
		getWorks(response.Links.Next, &allWorks)
	}

	return allWorks

}

type Author struct {
	Name string
}

func GetAuthorByKey(authorKey string) string {
	var response Author

	body := getClient.Get(url + authorKey + ".json")

	json.Unmarshal([]byte(body), &response)

	return response.Name

}
