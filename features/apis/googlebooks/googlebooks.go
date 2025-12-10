package googlebooks

//using the same http library as telego itself
import (
	"errors"
	"fmt"
	"net/url"

	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
)

type descriptionResponse struct {
	Items []struct {
		VolumeInfo struct {
			Description string `json:"description"`
		} `json:"volumeInfo"`
	} `json:"items"`
}

type downloadResponse struct {
	Items []struct {
		AccessInfo struct {
			Epub struct {
				AcsTokenLink string `json:"acsTokenLink"`
			} `json:"epub"`
			Pdf struct {
				AcsTokenLink string `json:"acsTokenLink"`
			} `json:"pdf"`
		} `json:"accessInfo"`
	} `json:"items"`
}

func GetBook(book string, author string, target string, key string) ([]byte, []string, []string, error) {
	//read https://developers.google.com/books/docs/v1/performance for tips on partical responses
	//network is also a bottleneck bigger than the cpu for my application, so let's use gzip' compression
	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=%s%s&key=%s&fields=kind,items(%s)",
		url.QueryEscape(book),
		url.QueryEscape(author),
		url.QueryEscape(key),
		url.QueryEscape(target))

	statusCode, body, err := fasthttp.Get(nil, url)
	if err != nil || statusCode != fasthttp.StatusOK {
		return nil, nil, nil, errors.New("Sorry, i could not communicate with google books api")
	}

	switch target {
	case "(accessInfo(epub(acsTokenLink),pdf(acsTokenLink)))":
		var downloadResponse downloadResponse
		if err := json.Unmarshal(body, &downloadResponse); err != nil {
			return nil, nil, nil, errors.New("Sorry, i could not get the download link for this book")
		}

		var pdfs []string
		var epubs []string

		for _, item := range downloadResponse.Items {
			if link := item.AccessInfo.Pdf.AcsTokenLink; link != "" {
				pdfs = append(pdfs, link)
			}
			if link := item.AccessInfo.Epub.AcsTokenLink; link != "" {
				epubs = append(epubs, link)
			}
		}

		return nil, pdfs, epubs, nil

	case "(volumeInfo(description))":
		var descriptionResponse descriptionResponse
		if err := json.Unmarshal(body, &descriptionResponse); err != nil {
			return nil, nil, nil, errors.New("Sorry, i could not get the description for this book")
		}

		//TODO: filter by language, locale, telegram's user ID in mongodb atlas will store it's locale option
		if len(descriptionResponse.Items) > 0 {
			return []byte(descriptionResponse.Items[0].VolumeInfo.Description), nil, nil, nil
		}
	}

	return nil, nil, nil, errors.New("Sorry, Something went wrong with the api target")
}
