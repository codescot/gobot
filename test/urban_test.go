package test

import (
	"testing"

	"os"

	"net/http"
	"net/http/httptest"

	"github.com/gurparit/gobot/command"
	"github.com/gurparit/gobot/env"
)

func TestUrbanDictionarySuccess(t *testing.T) {
	sampleUrban := `
{
	"list": [
		{
			"definition": "Portmanteau of [Tetris] and terrible, for when things just don't fit."
		}
	]
}
`

	testHttp := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(sampleUrban))
	}))

	defer testHttp.Close()

	os.Setenv("URBAN_URL", testHttp.URL)
	env.OS = env.LoadConfig()

	urban := command.Urban{}
	urban.Execute(func(response string) {
		if response != "tetrible - Portmanteau of [Tetris] and terrible, for when things just don't fit." {
			t.Log(response)
			t.Fail()
		}
	}, "tetrible")
}
