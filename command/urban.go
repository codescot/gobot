package command

import (
	"fmt"
	"math/rand"

	"net/http"
	"net/url"

	"github.com/gurparit/go-common/httpc"
)

const urbanResponse = "%s - %s"

// Urban urban dictionary command
type Urban struct{}

// UrbanResult urban dictionary json struct
type UrbanResult struct {
	List []struct {
		Definition string `json:"definition"`
	} `json:"list"`
}

// Execute run command
func (Urban) Execute(resp Response, event MessageEvent) {
	targetURL := httpc.FormatURL(
		"https://api.urbandictionary.com/v0/define?term=%s",
		url.QueryEscape(event.Message),
	)

	request := httpc.HTTP{
		TargetURL: targetURL,
		Method:    http.MethodGet,
	}

	var result UrbanResult
	if err := request.JSON(&result); err != nil {
		resp(fmt.Sprintf("[ud] %s", err.Error()))
		return
	}

	resultCount := len(result.List)
	if resultCount > 0 {
		randomDefinition := rand.Intn(resultCount)
		meaning := result.List[randomDefinition]

		result := fmt.Sprintf(urbanResponse, event.Message, meaning.Definition)
		resp(result)
	} else {
		resp("[ud] no results found")
		return
	}
}
