package command

import (
	"encoding/json"
	"fmt"

	"net/url"

	"github.com/gurparit/slackbot/util"
)

// GoogleURL base URL for Google Search
const GoogleURL = "https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&num=1&fields=items(title,link)&prettyPrint=false&q=%s"

// GoogleResponse base response for Google Search result
const GoogleResponse = "%s - %s"

// GoogleCommand the Google class
type GoogleCommand struct{}

// GoogleResult : sample response {"items":[{"title":"Netflix - Watch TV Shows Online, Watch Movies Online","link":"https://www.netflix.com/"}]}
type GoogleResult struct {
	Items []struct {
		Title string `json:"title"`
		Link  string `json:"link"`
	} `json:"items"`
}

func (google GoogleCommand) search(searchString string) (GoogleResult, error) {
	var err error

	queryString := url.QueryEscape(searchString)
	targetURL := fmt.Sprintf(GoogleURL, util.Config.GoogleAPI, util.Config.GoogleCX, queryString)

	httpCommand := HTTPCommand{URL: targetURL}
	body, err := httpCommand.Result()

	var result GoogleResult
	err = json.Unmarshal(body, &result)

	return result, err
}

// Execute GoogleCommand implementation
func (google GoogleCommand) Execute(respond func(string), query string) {
	result, err := google.search(query)
	if util.IsError(err) {
		respond("Google: (search error).")
		return
	}

	resultCount := len(result.Items)

	if resultCount > 0 {
		value := result.Items[0]
		result := fmt.Sprintf(GoogleResponse, value.Title, value.Link)

		respond(result)
	} else {
		respond("Google: no results found.")
	}
}
