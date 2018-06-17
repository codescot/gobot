package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gurparit/gobot/command"
	"github.com/gurparit/gobot/env"
)

func TestGoogleRequest(t *testing.T) {
	sampleGoogle := `
{
	"items": [
		{
			"title": "Netflix - Watch TV Shows Online, Watch Movies Online",
			"link": "https://www.netflix.com/"
		}
	]
}
`

	testHttp := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(sampleGoogle))
	}))

	defer testHttp.Close()

	os.Setenv("GOOGLE_SEARCH_URL", testHttp.URL)
	env.OS = env.LoadConfig()

	google := command.Google{}
	google.Execute(func(response string) {
		if response != "Netflix - Watch TV Shows Online, Watch Movies Online - https://www.netflix.com/" {
			t.Log(response)
			t.Fail()
		}
	}, "Netflix")
}
