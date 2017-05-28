package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

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

func (oxford OxfordDictionaryCommand) getTargetURL(searchString string) string {
	return fmt.Sprintf(OxfordDictionaryURL, searchString)
}

func (oxford OxfordDictionaryCommand) search(searchString string) (OxfordResult, error) {
	var err error

	httpCommand := HTTPCommand{}
	httpCommand.Headers = make(map[string]string)
	httpCommand.Headers["Accept"] = "application/json"
	httpCommand.Headers["app_id"] = config.OxfordID
	httpCommand.Headers["app_key"] = config.OxfordKey

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
	if IsError(err) {
		ircobj.Privmsg(messageChannel, sender+": (search error).")
		return
	}

	resultCount := len(result.Results)

	if resultCount > 0 {
		var definition string
		value := result.Results[0]

		if oxford.Etymology {
			definition = value.LexicalEntries[0].Entries[0].Etymologies[0]
		} else {
			definition = value.LexicalEntries[0].Entries[0].Senses[0].Definitions[0]
		}

		message := fmt.Sprintf(OxfordResponse, searchString, definition)

		ircobj.Privmsg(messageChannel, message)
	} else {
		ircobj.Privmsg(messageChannel, sender+": No results found.")
	}
}
