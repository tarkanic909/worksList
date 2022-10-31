package getClient

import (
	"io"
	"net/http"
)

func Get(url string) []byte {

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	return body
}
