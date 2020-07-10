package command

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/codescot/go-common/env"
	"github.com/codescot/go-common/httputil"
)

const youtubeVideoURL = "%s - http://www.youtube.com/watch?v=%s"

// Youtube youtube search command
type Youtube struct{}

// YoutubeResult container for youtube search result
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

// Execute run command
func (y Youtube) Execute(resp Response, event MessageEvent) {
	key := env.Optional("GOOGLE_API_KEY", "")
	if key == "" {
		return
	}

	targetURL := httputil.FormatURL("https://www.googleapis.com/youtube/v3/search?part=snippet&key=%s&maxResults=1&type=video&q=%s", key, url.QueryEscape(event.Message))
	request := httputil.HTTP{
		TargetURL: targetURL,
		Method:    http.MethodGet,
	}

	var result YoutubeResult
	if err := request.JSON(&result); err != nil {
		resp(fmt.Sprintf("[youtube] %s", err.Error()))
		return
	}

	resultCount := len(result.Items)

	if resultCount > 0 {
		value := result.Items[0]
		message := fmt.Sprintf(youtubeVideoURL, value.Snippet.Title, value.ID.VideoID)

		resp(message)
	} else {
		resp("[youtube] no results found")
	}
}
