package command

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/gurparit/marbles/util"
)

// OxfordDictionaryURL base Oxford Dictionary API URL
const OxfordDictionaryURL = "https://od-api.oxforddictionaries.com/api/v1/entries/en/%s"

// OxfordResponse the default response format
const OxfordResponse = "%s - %s"

const oxfordStandardResponse = "Oxford Dict.: no results found."

// OxfordDictionaryCommand dictionary command implementation
type OxfordDictionaryCommand struct {
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

func (oxford OxfordResult) hasEtyEntry() bool {
	isValid := (len(oxford.Results) > 0 && len(oxford.Results[0].LexicalEntries) > 0 && len(oxford.Results[0].LexicalEntries[0].Entries) > 0 && len(oxford.Results[0].LexicalEntries[0].Entries[0].Etymologies) > 0)

	return isValid
}

func (oxford OxfordResult) hasDefinitionEntry() bool {
	isValid := (len(oxford.Results) > 0 && len(oxford.Results[0].LexicalEntries) > 0 && len(oxford.Results[0].LexicalEntries[0].Entries) > 0 && len(oxford.Results[0].LexicalEntries[0].Entries[0].Senses) > 0 && len(oxford.Results[0].LexicalEntries[0].Entries[0].Senses[0].Definitions) > 0)

	return isValid
}

func (oxford OxfordResult) getEty() string {
	return oxford.Results[0].LexicalEntries[0].Entries[0].Etymologies[0]
}

func (oxford OxfordResult) getDefinition() string {
	return oxford.Results[0].LexicalEntries[0].Entries[0].Senses[0].Definitions[0]
}

func (oxford OxfordDictionaryCommand) getTargetURL(searchString string) string {
	return fmt.Sprintf(OxfordDictionaryURL, searchString)
}

func (oxford OxfordDictionaryCommand) search(searchString string) (OxfordResult, error) {
	var err error

	httpCommand := HTTPCommand{}
	httpCommand.Headers = make(map[string]string)
	httpCommand.Headers["Accept"] = "application/json"
	httpCommand.Headers["app_id"] = util.Config.OxfordID
	httpCommand.Headers["app_key"] = util.Config.OxfordKey

	queryString := url.QueryEscape(searchString)
	targetURL := oxford.getTargetURL(queryString)
	body, err := httpCommand.JSONResult(targetURL)

	var result OxfordResult
	err = json.Unmarshal(body, &result)

	return result, err
}

// Execute OxfordDictionaryCommand implementation
func (oxford OxfordDictionaryCommand) Execute(respond func(string), message string) {
	messages := strings.SplitN(message, " ", 2)
	searchString := messages[1]

	result, err := oxford.search(searchString)
	if util.IsError(err) {
		respond("Oxford Dict.: time to upskill that spelling game.")
		return
	}

	resultCount := len(result.Results)

	if resultCount > 0 {
		var definition string

		if oxford.Etymology && result.hasEtyEntry() {
			definition = result.getEty()
		} else if result.hasDefinitionEntry() {
			definition = result.getDefinition()
		} else {
			definition = oxfordStandardResponse
		}

		message := fmt.Sprintf(OxfordResponse, searchString, definition)

		respond(message)
	} else {
		respond(oxfordStandardResponse)
	}
}
