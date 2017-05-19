package main

import (
	"math/rand"
	"net/http"
	"strings"

	"net/url"

	irc "github.com/thoj/go-ircevent"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
)

// UrbanDictURL Urban Dictionary base URL
const UrbanDictURL = "http://www.urbandictionary.com/define.php?term="

// UDCommand Urban Dictionary command
type UDCommand struct{}

func (ud UDCommand) execute(ircobj *irc.Connection, event *irc.Event) {
	var err error

	messageChannel := event.Arguments[0]
	messages := strings.SplitN(event.Message(), " ", 1)
	searchString := messages[1]

	targetURL := UrbanDictURL + url.QueryEscape(searchString)
	response, err := http.Get(targetURL)
	defer response.Body.Close()

	root, err := html.Parse(response.Body)
	handleError(err)

	meanings := scrape.FindAll(root, scrape.ByClass("meaning"))
	numberOfMeanings := len(meanings)
	if numberOfMeanings == 0 {
		return
	}

	randomMeaning := rand.Intn(numberOfMeanings)
	meaning := scrape.Text(meanings[randomMeaning])

	ircobj.Privmsg(messageChannel, searchString+": "+meaning)
}
