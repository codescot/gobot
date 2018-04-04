package main

import (
	"fmt"
	"strings"

	"github.com/gurparit/slackbot/command"
	"github.com/gurparit/slackbot/util"
	"github.com/nlopes/slack"
)

var functions = make(map[string]command.Command)

func mapCommands() {
	functions["!go"] = command.HelloCommand{}
	functions["!time"] = command.TimeCommand{}
	functions["!g"] = command.GoogleCommand{}
	functions["!ud"] = command.UDCommand{}
	functions["!echo"] = command.EchoCommand{}
	functions["!spotify"] = command.SpotifyCommand{}
	functions["!yt"] = command.YoutubeCommand{}
	functions["!giphy"] = command.GiphyCommand{}
	functions["!define"] = command.OxfordDictionaryCommand{}

	ety := command.OxfordDictionaryCommand{}
	ety.Etymology = true
	functions["!ety"] = ety
}

// CatchErrors catch all errors and recover.
func CatchErrors() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

func run(bot func(string), message string) {
	defer CatchErrors()

	parameters := strings.Split(message, " ")
	action := parameters[0]

	if command, ok := functions[action]; ok {
		command.Execute(bot, message)
	}
}

func botStart() {
	config := util.Config
	username := config.Username

	api := slack.New(config.SlackToken)
	api.SetDebug(config.Debug)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch event := msg.Data.(type) {
		case *slack.MessageEvent:
			if event.Msg.Username != username {
				go run(func(response string) {
					api.PostMessage(event.Msg.Channel, response, slack.NewPostMessageParameters())
				}, event.Msg.Text)
			}

			break
		default:
			// do nothing.
		}
	}
}

func main() {
	mapCommands()
	botStart()
}
