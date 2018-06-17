package command

import (
	"os"
	"fmt"

	"net/url"

	"github.com/gurparit/gobot/env"
	"github.com/gurparit/gobot/httpc"
	"net/http"
)

// OxfordDictionaryURL base Oxford Dictionary API URL
const OxfordDictionaryURL = "https://od-api.oxforddictionaries.com/api/v1/entries/en/%s"

// OxfordResponse the default response format
const OxfordResponse = "%s - %s"

const oxfordNoResults = "[oxford] no results found"

// Oxford dictionary command implementation
type Oxford struct {
	Etymology bool
}

// OxfordResult oxford dictionary result
type OxfordResult struct {
	Results []struct {
		LexicalEntries []struct {
			Entries []struct {
				Etymologies []string `json:"etymologies"`
				Senses      []struct {
					Definitions []string `json:"definitions"`
				} `json:"senses"`
			} `json:"entries"`
		} `json:"lexicalEntries"`
	} `json:"results"`
}

var oxfordAppID = os.Getenv(env.OxfordAppID)
var oxfordApiKey = os.Getenv(env.OxfordApiKey)

func (ox OxfordResult) hasEtyEntry() bool {
	isValid := len(ox.Results) > 0 &&
		len(ox.Results[0].LexicalEntries) > 0 &&
		len(ox.Results[0].LexicalEntries[0].Entries) > 0 &&
		len(ox.Results[0].LexicalEntries[0].Entries[0].Etymologies) > 0

	return isValid
}

func (ox OxfordResult) hasDefinitionEntry() bool {
	isValid := len(ox.Results) > 0 &&
		len(ox.Results[0].LexicalEntries) > 0 &&
		len(ox.Results[0].LexicalEntries[0].Entries) > 0 &&
		len(ox.Results[0].LexicalEntries[0].Entries[0].Senses) > 0 &&
		len(ox.Results[0].LexicalEntries[0].Entries[0].Senses[0].Definitions) > 0

	return isValid
}

func (ox OxfordResult) getEty() string {
	return ox.Results[0].LexicalEntries[0].Entries[0].Etymologies[0]
}

func (ox OxfordResult) getDefinition() string {
	return ox.Results[0].LexicalEntries[0].Entries[0].Senses[0].Definitions[0]
}

func (ox Oxford) search(searchString string) (OxfordResult, error) {
	targetURL := fmt.Sprintf(
		OxfordDictionaryURL,
		url.QueryEscape(searchString),
	)

	headers := map[string]string{
		"Accept": "application/json",
		"app_id": oxfordAppID,
		"app_key": oxfordApiKey,
	}

	request := httpc.HTTP{
		TargetURL: targetURL,
		Method: http.MethodGet,
		Headers: headers,
	}

	var result OxfordResult
	err := request.JSON(&result)

	return result, err
}

// Execute Oxford implementation
func (ox Oxford) Execute(r Response, query string) {
	result, err := ox.search(query)
	if err != nil {
		r(fmt.Sprintf("[ox] %s", err.Error()))
		return
	}

	resultCount := len(result.Results)

	if resultCount > 0 {
		var definition string

		if ox.Etymology && result.hasEtyEntry() {
			definition = result.getEty()
		} else if result.hasDefinitionEntry() {
			definition = result.getDefinition()
		} else {
			definition = oxfordNoResults
		}

		message := fmt.Sprintf(OxfordResponse, query, definition)

		r(message)
	} else {
		r(oxfordNoResults)
	}
}
