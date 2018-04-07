package command

import (
	"fmt"
	"math/rand"
	"net/http"

	"net/url"

	"github.com/gurparit/slackbot/util"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
)

// UrbanDictURL Urban Dictionary base URL
const UrbanDictURL = "http://www.urbandictionary.com/define.php?term=%s"

// UDCommand Urban Dictionary command
type UDCommand struct{}

// Execute UDCommand implementation
func (ud UDCommand) Execute(respond func(string), query string) {
	var err error

	queryString := url.QueryEscape(query)
	targetURL := fmt.Sprintf(UrbanDictURL, queryString)

	response, err := http.Get(targetURL)

	if err != nil {
		respond("UD: " + err.Error())
		return
	}

	defer response.Body.Close()

	root, err := html.Parse(response.Body)
	if util.IsError(err) {
		respond("UD: " + err.Error())
		return
	}

	meanings := scrape.FindAll(root, scrape.ByClass("meaning"))
	numberOfMeanings := len(meanings)
	if numberOfMeanings == 0 {
		respond("UD: no results found.")
		return
	}

	randomMeaning := rand.Intn(numberOfMeanings)
	meaning := scrape.Text(meanings[randomMeaning])

	respond(query + ": " + meaning)
}
