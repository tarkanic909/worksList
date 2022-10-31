package searchService

import (
	"encoding/json"
	"strings"
	"worksList/getClient"
)

const url = "https://openlibrary.org/search.json?title="

type Doc struct {
	Title       string
	AuthorsName []string `json:"author_name"`
	AuthorsKey  []string `json:"author_key"`
}

type SearchResponse struct {
	Docs []Doc
}

func Search(query string) []Doc {
	var response SearchResponse

	body := getClient.Get(url + strings.ReplaceAll(query, " ", "+"))
	json.Unmarshal([]byte(body), &response)

	return response.Docs
}
