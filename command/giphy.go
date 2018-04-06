package command

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"strings"

	"github.com/gurparit/slackbot/util"
)

// GiphyURL base url for API call
const GiphyURL = "https://api.giphy.com/v1/gifs/search?api_key=%s&q=%s&limit=5&lang=en"

// GiphyResponse base response for Giphy Search result
const GiphyResponse = "%s - %s"

// GiphyCommand Giphy API
type GiphyCommand struct{}

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

func (giphy GiphyCommand) search(searchString string) (GiphyResult, error) {
	var err error

	queryString := url.QueryEscape(searchString)
	targetURL := fmt.Sprintf(GiphyURL, util.Config.GiphyAPI, queryString)

	httpCommand := HTTPCommand{URL: targetURL}
	body, err := httpCommand.Result()

	var result GiphyResult
	err = json.Unmarshal(body, &result)

	return result, err
}

// Execute GiphyCommand implementation
func (giphy GiphyCommand) Execute(respond func(string), message string) {
	messages := strings.SplitN(message, " ", 2)
	searchString := messages[1]

	result, err := giphy.search(searchString)
	if util.IsError(err) {
		respond("Giphy: (search error).")
		return
	}

	resultCount := len(result.Data)

	if resultCount > 0 {
		randomGiphy := rand.Intn(resultCount)

		value := result.Data[randomGiphy]
		result := fmt.Sprintf(GiphyResponse, value.Title, value.Images.Original.URL)

		respond(result)
	} else {
		respond("Giphy: no results found.")
	}
}
