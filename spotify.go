package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	irc "github.com/thoj/go-ircevent"
)

// SpotifyURL the base Spotify search URL
const SpotifyURL = "https://api.spotify.com/v1/search?type=%s&limit=1&q=%s"

// SpotifyCommand spotify implementation
type SpotifyCommand struct{}

// SpotifyResult spotify result implementation
type SpotifyResult struct {
	Result []struct {
		Items []struct {
			Name         string `json:"name"`
			ExternalURLS []struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
		} `json:"items"`
	} `json:"albums" json:"artists" json:"tracks"`
}

func (spotify SpotifyCommand) getTargetURL(searchString string) string {
	return fmt.Sprintf(GoogleURL, config.GoogleAPI, config.GoogleCX, searchString)
}

func (spotify SpotifyCommand) search(searchString string) (SpotifyResult, error) {
	var err error

	httpCommand := HTTPCommand{}
	queryString := url.QueryEscape(searchString)
	targetURL := spotify.getTargetURL(queryString)
	body, err := httpCommand.JSONResult(targetURL)

	var result SpotifyResult
	err = json.Unmarshal(body, &result)

	return result, err
}

// Execute spotify search implementation
func (spotify SpotifyCommand) Execute(ircobj *irc.Connection, event *irc.Event) {
	searchString := ""
	sender := event.Nick
	messageChannel := event.Arguments[0]

	messages := strings.SplitN(event.Message(), " ", 3)
	searchType := messages[1]
	if len(messages) > 1 {
		searchString = messages[2]
	}

	switch searchType {
	case "album":
	case "artist":
	case "track":
		break
	default:
		messages = strings.SplitN(event.Message(), " ", 2)
		searchString = messages[1]
		break
	}

	result, err := spotify.search(searchString)
	if IsError(err) {
		ircobj.Privmsg(messageChannel, sender+": (search error).")
		return
	}

	resultCount := len(result.Result)

	if resultCount > 0 {
		value := result.Result[0].Items[0]
		ircobj.Privmsg(messageChannel, value.Name+" - "+value.ExternalURLS[0].Spotify)
	} else {
		ircobj.Privmsg(messageChannel, sender+": No results found.")
	}
}
