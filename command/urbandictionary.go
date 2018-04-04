package command

import (
	"math/rand"
	"net/http"
	"strings"

	"net/url"

	"github.com/gurparit/slackbot/util"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
)

// UrbanDictURL Urban Dictionary base URL
const UrbanDictURL = "http://www.urbandictionary.com/define.php?term="

// UDCommand Urban Dictionary command
type UDCommand struct{}

// Execute UDCommand implementation
func (ud UDCommand) Execute(respond func(string), message string) {
	var err error
	messages := strings.SplitN(message, " ", 2)
	searchString := messages[1]

	targetURL := UrbanDictURL + url.QueryEscape(searchString)
	response, err := http.Get(targetURL)
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

	respond(searchString + ": " + meaning)
}
