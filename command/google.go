package command

import (
	"encoding/json"
	"fmt"
	"strings"

	"net/url"

	"github.com/gurparit/marbles/util"
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

func (google GoogleCommand) getTargetURL(searchString string) string {
	return fmt.Sprintf(GoogleURL, util.Marbles.GoogleAPI, util.Marbles.GoogleCX, searchString)
}

func (google GoogleCommand) search(searchString string) (GoogleResult, error) {
	var err error

	httpCommand := HTTPCommand{}
	queryString := url.QueryEscape(searchString)
	targetURL := google.getTargetURL(queryString)
	body, err := httpCommand.JSONResult(targetURL)

	var result GoogleResult
	err = json.Unmarshal(body, &result)

	return result, err
}

// Execute GoogleCommand implementation
func (google GoogleCommand) Execute(respond func(string), message string) {
	messages := strings.SplitN(message, " ", 2)
	searchString := messages[1]

	result, err := google.search(searchString)
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
		respond("No results found.")
	}
}
