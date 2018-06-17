package command

import (
	"fmt"
	"math/rand"

	"net/http"
	"net/url"

	"github.com/gurparit/gobot/httpc"
)

// UrbanDictURL Urban Dictionary base URL
const UrbanDictURL = "https://api.urbandictionary.com/v0/define?term=%s"

// UrbanResponse base response for Google Search result
const UrbanResponse = "%s - %s"

// Urban Urban Dictionary command
type Urban struct{}

// UrbanResult : sample response {unknown}
type UrbanResult struct {
	List []struct {
		Definition string `json:"definition"`
	} `json:"list"`
}

// Execute Google implementation
func (Urban) Execute(r Response, query string) {
	targetURL := fmt.Sprintf(
		UrbanDictURL,
		url.QueryEscape(query),
	)

	request := httpc.HTTP{
		TargetURL: targetURL,
		Method:    http.MethodGet,
	}

	var result UrbanResult
	if err := request.JSON(&result); err != nil {
		r(fmt.Sprintf("[ud] %s", err.Error()))
		return
	}

	resultCount := len(result.List)
	if resultCount > 0 {
		randomDefinition := rand.Intn(resultCount)
		meaning := result.List[randomDefinition]

		result := fmt.Sprintf(UrbanResponse, query, meaning.Definition)
		r(result)
	} else {
		r("[ud] no results found")
		return
	}
}
