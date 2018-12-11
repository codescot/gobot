package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gurparit/go-common/env"
	"github.com/gurparit/twitchbot/command"
)

func TestYoutubeRequest(t *testing.T) {
	sampleYoutube := `
{
	"items": [
		{
			"snippet": {
				"title": "Ne-Yo - GOOD MAN"
			},
			"id": {
				"videoId": "abc123"
			}
		}
	]
}
`

	testHttp := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(sampleYoutube))
	}))

	defer testHttp.Close()

	os.Setenv("YOUTUBE_SEARCH_URL", testHttp.URL)
	env.Read(&command.OS)

	youtube := command.Youtube{}
	youtube.Execute(func(response string) {
		if response != "Ne-Yo - GOOD MAN - http://www.youtube.com/watch?v=abc123" {
			t.Log(response)
			t.Fail()
		}
	}, "Ne-Yo - GOOD MAN")
}
