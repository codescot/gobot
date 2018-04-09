package testing

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gurparit/slackbot/command"
)

func TestGoogleSuccess(test *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"items":[{"title":"Netflix - Watch TV Shows Online, Watch Movies Online","link":"https://www.netflix.com/"}]}`)
	}))

	var result command.GoogleResult

	err := command.JSON(func() string {
		return testServer.URL
	}, &result)

	if err != nil {
		test.Errorf(err.Error())
	}

	if result.Items[0].Title != "Netflix - Watch TV Shows Online, Watch Movies Online" {
		test.Errorf("Title didn't match expected. Actual: %s", result.Items[0].Title)
	}

	if result.Items[0].Link != "https://www.netflix.com/" {
		test.Errorf("Link didn't match expected. Actual: %s", result.Items[0].Link)
	}
}
