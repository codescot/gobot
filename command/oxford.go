package command

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/gurparit/marbles/util"
	irc "github.com/thoj/go-ircevent"
)

// OxfordDictionaryURL base Oxford Dictionary API URL
const OxfordDictionaryURL = "https://od-api.oxforddictionaries.com/api/v1/entries/en/%s"

// OxfordResponse the default response format
const OxfordResponse = "%s - %s"

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

func (result OxfordResult) hasEtyEntry() bool {
	isValid := len(result.Results) > 0
		&& len(result.Results.LexicalEntries) > 0
		&& len(result.Results.LexicalEntries.Entries) > 0
		&& len(result.Results.LexicalEntries.Entries.Etymologies) > 0
	
	return isValid
}

func (result OxfordResult) hasDictionaryEntry() bool {
	isValid := len(result.Results) > 0
		&& len(result.Results.LexicalEntries) > 0
		&& len(result.Results.LexicalEntries.Entries) > 0
		&& len(result.Results.LexicalEntries.Entries.Senses) > 0
		&& len(result.Results.LexicalEntries.Entries.Senses.Definitions) > 0
	
	return isValid
}

func (oxford OxfordDictionaryCommand) getTargetURL(searchString string) string {
	return fmt.Sprintf(OxfordDictionaryURL, searchString)
}

func (oxford OxfordDictionaryCommand) search(searchString string) (OxfordResult, error) {
	var err error

	httpCommand := HTTPCommand{}
	httpCommand.Headers = make(map[string]string)
	httpCommand.Headers["Accept"] = "application/json"
	httpCommand.Headers["app_id"] = util.Marbles.OxfordID
	httpCommand.Headers["app_key"] = util.Marbles.OxfordKey

	queryString := url.QueryEscape(searchString)
	targetURL := oxford.getTargetURL(queryString)
	body, err := httpCommand.JSONResult(targetURL)

	var result OxfordResult
	err = json.Unmarshal(body, &result)

	return result, err
}

// Execute OxfordDictionaryCommand implementation
func (oxford OxfordDictionaryCommand) Execute(ircobj *irc.Connection, event *irc.Event) {
	sender := event.Nick
	messageChannel := event.Arguments[0]

	messages := strings.SplitN(event.Message(), " ", 2)
	searchString := messages[1]

	result, err := oxford.search(searchString)
	if util.IsError(err) {
		ircobj.Privmsg(messageChannel, sender+": learn to spell, noob!")
		return
	}

	resultCount := len(result.Results)

	if resultCount > 0 {
		var definition string

		if oxford.Etymology && result.hasEtyEntry() {
			definition = result.Results[0].LexicalEntries[0].Entries[0].Etymologies[0]
		} else if result.hasDictionaryEntry() {
			definition = result.Results[0].LexicalEntries[0].Entries[0].Senses[0].Definitions[0]
		} else {
			definition = "no results found."
		}

		message := fmt.Sprintf(OxfordResponse, searchString, definition)

		ircobj.Privmsg(messageChannel, message)
	} else {
		ircobj.Privmsg(messageChannel, sender+": no results found.")
	}
}
