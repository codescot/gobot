package command

import (
	"fmt"

	"net/url"

	"github.com/gurparit/slackbot/util"
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

// Execute YoutubeCommand implementation
func (youtube YoutubeCommand) Execute(respond func(string), query string) {
	var result YoutubeResult

	err := JSON(func() string {
		queryString := url.QueryEscape(query)
		targetURL := fmt.Sprintf(YoutubeURL, util.Config.GoogleAPI, queryString)

		return targetURL
	}, &result)

	if util.IsError(err) {
		respond("Youtube: (search error).")
		return
	}

	resultCount := len(result.Items)

	if resultCount > 0 {
		value := result.Items[0]
		message := fmt.Sprintf(YoutubeVideoURL, value.Snippet.Title, value.ID.VideoID)

		respond(message)
	} else {
		respond("Youtube: no results found.")
	}
}
