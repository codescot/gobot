package command

import (
	"fmt"
	"net/url"

	"net/http"

	"github.com/gurparit/go-common/httpc"
)

const GoogleResponse = "%s - %s"

type Google struct{}

// GoogleResult : sample response {"items":[{"title":"Netflix - Watch TV Shows Online, Watch Movies Online","link":"https://www.netflix.com/"}]}
type GoogleResult struct {
	Items []struct {
		Title string `json:"title"`
		Link  string `json:"link"`
	} `json:"items"`
}

// Execute Google implementation
func (Google) Execute(r Response, query string) {
	targetURL := httpc.FormatURL(
		OS.GoogleURL,
		OS.GoogleKey,
		OS.GoogleSearchID,
		url.QueryEscape(query),
	)

	request := httpc.HTTP{
		TargetURL: targetURL,
		Method:    http.MethodGet,
	}

	var result GoogleResult
	if err := request.JSON(&result); err != nil {
		r(fmt.Sprintf("[google] %s", err.Error()))
		return
	}

	resultCount := len(result.Items)

	if resultCount > 0 {
		value := result.Items[0]
		result := fmt.Sprintf(GoogleResponse, value.Title, value.Link)

		r(result)
	} else {
		r("[google] no results found")
	}
}
