package command

import (
	"fmt"

	"net/http"
	"net/url"

	"github.com/gurparit/go-common/httpc"
)

const oxfordResponse = "%s - %s"

const oxfordNoResults = "[oxford] no results found"

// Oxford oxford dictionary search command
type Oxford struct {
	Etymology bool
}

// OxfordResult container for oxford dictionary response
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

func (ox Oxford) search(event MessageEvent) (OxfordResult, error) {
	targetURL := httpc.FormatURL(
		"https://od-api.oxforddictionaries.com/api/v1/entries/en/%s",
		url.QueryEscape(event.Message),
	)

	headers := map[string]string{
		"Accept":  "application/json",
		"app_id":  event.Config.OxfordAppID,
		"app_key": event.Config.OxfordKey,
	}

	request := httpc.HTTP{
		TargetURL: targetURL,
		Method:    http.MethodGet,
		Headers:   headers,
	}

	var result OxfordResult
	err := request.JSON(&result)

	return result, err
}

// Execute run command
func (ox Oxford) Execute(resp Response, event MessageEvent) {
	result, err := ox.search(event)
	if err != nil {
		resp(fmt.Sprintf("[ox] %s", err.Error()))
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

		message := fmt.Sprintf(oxfordResponse, event.Message, definition)

		resp(message)
	} else {
		resp(oxfordNoResults)
	}
}
