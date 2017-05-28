package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"net/url"

	irc "github.com/thoj/go-ircevent"
)

// YoutubeURL base URL for Youtube Search
const YoutubeURL = "https://www.googleapis.com/youtube/v3/search?part=snippet&key=%s&maxResults=1&type=video&q=%s"

// YoutubeVideoURL base URL for Youtube Videos
const YoutubeVideoURL = "%s - http://www.youtube.com/watch?v=%s"

// YoutubeCommand the Youtube class
type YoutubeCommand struct{}

// YoutubeResult : sample response
type YoutubeResult struct {
	Items []struct {
		Snippet struct {
			Title string `json:"title"`
		} `json:"snippet"`
		ID struct {
			VideoID string `json:"videoId"`
		} `json:"id"`
	} `json:"items"`
}

func (youtube YoutubeCommand) getTargetURL(searchString string) string {
	return fmt.Sprintf(YoutubeURL, config.GoogleAPI, searchString)
}

func (youtube YoutubeCommand) search(searchString string) (YoutubeResult, error) {
	var err error

	httpCommand := HTTPCommand{}
	queryString := url.QueryEscape(searchString)
	targetURL := youtube.getTargetURL(queryString)
	body, err := httpCommand.JSONResult(targetURL)

	var result YoutubeResult
	err = json.Unmarshal(body, &result)

	return result, err
}

// Execute YoutubeCommand implementation
func (youtube YoutubeCommand) Execute(ircobj *irc.Connection, event *irc.Event) {
	sender := event.Nick
	messageChannel := event.Arguments[0]

	messages := strings.SplitN(event.Message(), " ", 2)
	searchString := messages[1]

	result, err := youtube.search(searchString)
	if IsError(err) {
		ircobj.Privmsg(messageChannel, sender+": (search error).")
		return
	}

	resultCount := len(result.Items)

	if resultCount > 0 {
		value := result.Items[0]
		message := fmt.Sprintf(YoutubeVideoURL, value.Snippet.Title, value.ID.VideoID)

		ircobj.Privmsg(messageChannel, message)
	} else {
		ircobj.Privmsg(messageChannel, sender+": No results found.")
	}
}
