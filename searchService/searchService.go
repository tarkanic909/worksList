package searchService

import (
	"encoding/json"
	"net/url"
	"worksList/getClient"
)

const searchUrl = "https://openlibrary.org/search.json?title="

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

	body := getClient.Get(searchUrl + url.QueryEscape(query))
	json.Unmarshal([]byte(body), &response)

	return response.Docs
}
