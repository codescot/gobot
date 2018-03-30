package command

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/gurparit/marbles/util"
)

// SpotifyURL the base Spotify search URL
const SpotifyURL = "https://api.spotify.com/v1/search?type=%s&limit=1&q=%s"

// SpotifyCommand spotify implementation
type SpotifyCommand struct{}

// SpotifyResult spotify result implementation
type SpotifyResult struct {
	Albums  SpotifyInnerResult `json:"albums"`
	Artists SpotifyInnerResult `json:"artists"`
	Tracks  SpotifyInnerResult `json:"tracks"`
}

// SpotifyInnerResult inner part of the spotify results
type SpotifyInnerResult struct {
	Items []struct {
		Name         string `json:"name"`
		ExternalURLS struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
	} `json:"items"`
}

func (spotify SpotifyInnerResult) length() int {
	return len(spotify.Items)
}

func (spotify SpotifyCommand) getTargetURL(searchType string, searchString string) string {
	return fmt.Sprintf(SpotifyURL, searchType, searchString)
}

func (spotify SpotifyCommand) search(searchType string, searchString string) (SpotifyResult, error) {
	var err error

	httpCommand := HTTPCommand{}
	queryString := url.QueryEscape(searchString)
	targetURL := spotify.getTargetURL(searchType, queryString)
	body, err := httpCommand.JSONResult(targetURL)

	var result SpotifyResult
	err = json.Unmarshal(body, &result)

	return result, err
}

// Execute spotify search implementation
func (spotify SpotifyCommand) Execute(respond func(string), message string) {
	searchString := ""
	messages := strings.SplitN(message, " ", 3)
	searchType := messages[1]
	if len(messages) > 2 {
		searchString = messages[2]
	}

	switch searchType {
	case "album":
	case "artist":
	case "track":
		break
	default:
		messages = strings.SplitN(message, " ", 2)
		searchType = "track"
		searchString = messages[1]
	}

	result, err := spotify.search(searchType, searchString)
	if util.IsError(err) {
		respond("Spotify: (search error).")
		return
	}

	var spotifyResults SpotifyInnerResult

	switch searchType {
	case "album":
		spotifyResults = result.Albums
		break
	case "artist":
		spotifyResults = result.Artists
		break
	case "track":
		spotifyResults = result.Tracks
		break
	}

	resultCount := spotifyResults.length()

	if resultCount > 0 {
		value := spotifyResults.Items[0]
		respond(value.Name + " - " + value.ExternalURLS.Spotify)
	} else {
		respond("Spotify: No results found.")
	}
}
