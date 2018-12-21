package command

import (
	"fmt"
	"net/url"

	"net/http"

	"github.com/gurparit/go-common/httpc"
)

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

// Execute run command
func (Giphy) Execute(resp Response, event MessageEvent) {
	targetURL := httpc.FormatURL(
		"https://api.giphy.com/v1/gifs/search?api_key=%s&q=%s&limit=1&lang=en",
		event.Keys.GiphyKey,
		url.QueryEscape(event.Message),
	)

	request := httpc.HTTP{
		TargetURL: targetURL,
		Method:    http.MethodGet,
	}

	var result GiphyResult
	if err := request.JSON(&result); err != nil {
		resp(fmt.Sprintf("[gif] %s", err.Error()))
		return
	}

	resultCount := len(result.Data)

	if resultCount > 0 {
		value := result.Data[0]
		result := fmt.Sprintf(GiphyResponse, value.Title, value.Images.Original.URL)

		resp(result)
	} else {
		resp("[gif] no results found")
	}
}
