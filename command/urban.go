package command

import (
	"fmt"
	"math/rand"

	"net/http"
	"net/url"

	"github.com/gurparit/go-common/httpc"
)

const UrbanResponse = "%s - %s"

type Urban struct{}

type UrbanResult struct {
	List []struct {
		Definition string `json:"definition"`
	} `json:"list"`
}

func (Urban) Execute(r Response, query string) {
	targetURL := FormatURL(
		OS.UrbanURL,
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
