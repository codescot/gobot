package main

import (
	"encoding/json"
	"fmt"
	"strings"

	irc "github.com/thoj/go-ircevent"
)

// GoogleURL base URL for Google Search
const GoogleURL = "https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&num=1&fields=items(title,link)&prettyPrint=false&q=%s"

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
	return fmt.Sprintf(GoogleURL, config.GoogleAPI, config.GoogleCX, searchString)
}

func (google GoogleCommand) search(searchString string) GoogleResult {
	var err error

	httpCommand := HTTPCommand{}
	targetURL := google.getTargetURL(searchString)
	body, err := httpCommand.getJSONResult(targetURL)

	var result GoogleResult
	err = json.Unmarshal(body, &result)
	handleError(err)

	return result
}

func (google GoogleCommand) execute(ircobj *irc.Connection, event *irc.Event) {
	sender := event.Nick
	messageChannel := event.Arguments[0]

	messages := strings.SplitN(event.Message(), " ", 1)
	searchString := messages[1]

	result := google.search(searchString)
	resultCount := len(result.Items)

	if resultCount > 0 {
		value := result.Items[0]
		ircobj.Privmsg(messageChannel, value.Title+" - "+value.Link)
	} else {
		ircobj.Privmsg(messageChannel, sender+": No results found.")
	}
}
