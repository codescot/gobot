package command

import (
	"fmt"
	"net/url"

	"net/http"

	"github.com/gurparit/go-common/httpc"
)

const YoutubeVideoURL = "%s - http://www.youtube.com/watch?v=%s"

type Youtube struct{}

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

func (Youtube) Execute(r Response, query string) {
	targetURL := FormatURL(
		OS.YoutubeURL,
		OS.GoogleKey,
		url.QueryEscape(query),
	)

	request := httpc.HTTP{
		TargetURL: targetURL,
		Method:    http.MethodGet,
	}

	var result YoutubeResult
	if err := request.JSON(&result); err != nil {
		r(fmt.Sprintf("[youtube] %s", err.Error()))
		return
	}

	resultCount := len(result.Items)

	if resultCount > 0 {
		value := result.Items[0]
		message := fmt.Sprintf(YoutubeVideoURL, value.Snippet.Title, value.ID.VideoID)

		r(message)
	} else {
		r("[youtube] no results found")
	}
}
