package openlibrary

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/valyala/fasthttp"
)

type openlibraryResponse struct {
	Docs []struct {
		Key string `json:"key"`
	} `json:"docs"`
}

func GetBook(title string, author string) (string, error) {
	url := fmt.Sprintf("https://openlibrary.org/search.json?title=%s&author=%s&limit=3",
		url.QueryEscape(title),
		url.QueryEscape(author))

	statusCode, body, err := fasthttp.Get(nil, url)
	if err != nil || statusCode != fasthttp.StatusOK {
		return "", fmt.Errorf("Sorry, i could not communicate with OpenLibrary API")
	}

	var olResponse openlibraryResponse
	if err := json.Unmarshal(body, &olResponse); err != nil {
		return "", fmt.Errorf("Sorry, i could not get the book information from OpenLibrary")
	}

	if len(olResponse.Docs) == 0 {
		return "", fmt.Errorf("Sorry, i could not find the book you requested on OpenLibrary")
	}

	return fmt.Sprintf("https://openlibrary.org%s", olResponse.Docs[0].Key), nil

}
