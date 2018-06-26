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

// Execute Giphy implementation
func (Giphy) Execute(r Response, query string) {
	targetURL := FormatURL(
		OS.GiphyURL,
		OS.GiphyKey,
		url.QueryEscape(query),
	)

	request := httpc.HTTP{
		TargetURL: targetURL,
		Method:    http.MethodGet,
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
