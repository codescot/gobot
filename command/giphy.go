package command

import (
	"fmt"
	"os"

	"net/url"

	"github.com/gurparit/gobot/env"
	"github.com/gurparit/gobot/httpc"
	"net/http"
)

// GiphyURL base url for API call
const GiphyURL = "https://api.giphy.com/v1/gifs/search?api_key=%s&q=%s&limit=1&lang=en"

// GiphyResponse base response for Giphy Search result
const GiphyResponse = "%s - %s"

// Giphy Giphy API
type Giphy struct{}

// GiphyResult JSON result struct for unmarshalling
type GiphyResult struct {
	Data []struct {
		ID     string `json:"id"`
		URL    string `json:"url"`
		Type   string `json:"type"`
		Title  string `json:"title"`
		Images struct {
			Original struct {
				URL string `json:"url"`
			} `json:"original"`
		} `json:"images"`
	} `json:"data"`
}

// Execute Giphy implementation
func (Giphy) Execute(r Response, query string) {
	targetURL := fmt.Sprintf(
		GiphyURL,
		os.Getenv(env.GiphyApiKey),
		url.QueryEscape(query),
	)

	request := httpc.HTTP{
		TargetURL: targetURL,
		Method: http.MethodGet,
	}

	var result GiphyResult
	if err := request.JSON(&result); err != nil {
		r(fmt.Sprintf("[gif] %s", err.Error()))
		return
	}

	resultCount := len(result.Data)

	if resultCount > 0 {
		value := result.Data[0]
		result := fmt.Sprintf(GiphyResponse, value.Title, value.Images.Original.URL)

		r(result)
	} else {
		r("[gif] no results found")
	}
}
