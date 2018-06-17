package test

import (
	"testing"

	"os"

	"net/http"
	"net/http/httptest"

	"github.com/gurparit/gobot/command"
	"github.com/gurparit/gobot/env"
)

func TestGiphySuccess(t *testing.T) {
	sampleGiphy := `
{
	"data": [
		{
			"id": "abc123",
			"url": "http://img.example.com/meme.gif",
			"type": "gif",
			"title": "funny meme gif",
			"images": {
				"original": {
					"url": "http://img.example.com/original/meme.gif"
				}
			}
		}
	]
}
`

	testHttp := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(sampleGiphy))
	}))

	os.Setenv("GIPHY_URL", testHttp.URL)
	env.OS = env.LoadConfig()

	gif := command.Giphy{}

	gif.Execute(func(response string) {
		if response != "funny meme gif - http://img.example.com/original/meme.gif" {
			t.Log(response)
			t.Fail()
		}
	}, "meme")
}
