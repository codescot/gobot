package command

import (
	"fmt"
	"math/rand"

	"net/url"

	"github.com/gurparit/slackbot/util"
)

// UrbanDictURL Urban Dictionary base URL
const UrbanDictURL = "https://api.urbandictionary.com/v0/define?term=%s"

// UrbanResponse base response for Google Search result
const UrbanResponse = "%s - %s"

// UDCommand Urban Dictionary command
type UDCommand struct{}

// UrbanResult : sample response {unknown}
type UrbanResult struct {
	List []struct {
		Definition string `json:"definition"`
	} `json:"list"`
}

// Execute GoogleCommand implementation
func (ud UDCommand) Execute(respond func(string), query string) {
	var result UrbanResult

	err := JSON(func() string {
		queryString := url.QueryEscape(query)
		targetURL := fmt.Sprintf(UrbanDictURL, queryString)

		return targetURL
	}, &result)

	if util.IsError(err) {
		respond("Google: (search error).")
		return
	}

	resultCount := len(result.List)
	if resultCount > 0 {
		randomDefinition := rand.Intn(resultCount)
		meaning := result.List[randomDefinition]

		result := fmt.Sprintf(UrbanResponse, query, meaning.Definition)
		respond(result)
	} else {
		respond("UD: no results found.")
		return
	}
}
