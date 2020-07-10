package command

import (
	"fmt"
	"net/url"

	"net/http"

	"github.com/codescot/go-common/env"
	"github.com/codescot/go-common/httputil"
)

const googleResponse = "%s - %s"

// Google google search command
type Google struct{}

// GoogleResult : sample response {"items":[{"title":"Netflix - Watch TV Shows Online, Watch Movies Online","link":"https://www.netflix.com/"}]}
type GoogleResult struct {
	Items []struct {
		Title string `json:"title"`
		Link  string `json:"link"`
	} `json:"items"`
}

// Execute run command
func (g Google) Execute(resp Response, event MessageEvent) {
	searchID := env.Optional("GOOGLE_API_SEARCH_ID", "")
	key := env.Optional("GOOGLE_API_KEY", "")
	if key == "" || searchID == "" {
		return
	}

	targetURL := httputil.FormatURL(
		"https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&num=1&fields=items(title,link)&prettyPrint=false&q=%s",
		key,
		searchID,
		url.QueryEscape(event.Message),
	)

	request := httputil.HTTP{
		TargetURL: targetURL,
		Method:    http.MethodGet,
	}

	var result GoogleResult
	if err := request.JSON(&result); err != nil {
		resp(fmt.Sprintf("[google] %s", err.Error()))
		return
	}

	resultCount := len(result.Items)

	if resultCount > 0 {
		value := result.Items[0]
		result := fmt.Sprintf(googleResponse, value.Title, value.Link)

		resp(result)
	} else {
		resp("[google] no results found")
	}
}
