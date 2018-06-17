package testing

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gurparit/gobot/command"
)

var youtubeSampleJSON = `
{
	"items": [
		{
			"snippet": {
				"title": "Hello, World."
			},
			"id": {
				"videoId": "abcdef123"
			}
		}
	]
}
`

func TestYoutubeSuccess(test *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, youtubeSampleJSON)
	}))

	var result command.YoutubeResult

	err := command.JSON(func() string {
		return testServer.URL
	}, &result)

	if err != nil {
		test.Errorf(err.Error())
	}

	if result.Items[0].Snippet.Title != "Hello, World." {
		test.Errorf("Title didn't match expected. Actual: %s", result.Items[0].Snippet.Title)
	}

	if result.Items[0].ID.VideoID != "abcdef123" {
		test.Errorf("Link didn't match expected. Actual: %s", result.Items[0].ID.VideoID)
	}
}
