package command

import (
	"encoding/json"
	"fmt"
	"strings"

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

func (youtube YoutubeCommand) search(searchString string) (YoutubeResult, error) {
	var err error

	httpCommand := HTTPCommand{}
	queryString := url.QueryEscape(searchString)
	targetURL := httpCommand.GetTargetURL(YoutubeURL, util.Config.GoogleAPI, queryString)
	body, err := httpCommand.GetJSONResult(targetURL)

	var result YoutubeResult
	err = json.Unmarshal(body, &result)

	return result, err
}

// Execute YoutubeCommand implementation
func (youtube YoutubeCommand) Execute(respond func(string), message string) {
	messages := strings.SplitN(message, " ", 2)
	searchString := messages[1]

	result, err := youtube.search(searchString)
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
		respond("Youtube: No results found.")
	}
}
