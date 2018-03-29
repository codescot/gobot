package main

import (
	"fmt"
	"strings"

	"github.com/gurparit/marbles/command"
	"github.com/gurparit/marbles/util"
	"github.com/nlopes/slack"
	"github.com/thoj/go-ircevent"
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

func run(ircobj *irc.Connection, event *irc.Event) {
	defer CatchErrors()

	message := event.Message()
	parameters := strings.Split(message, " ")
	action := parameters[0]

	if command, ok := functions[action]; ok {
		command.Execute(ircobj, event)
	}
}

func botStart() {
	config := util.Marbles

	api := slack.New(config.SlackToken)
	api.SetDebug(config.Debug)

	var channels []*slack.Channel
	channelsList := config.Channels
	for _, channel := range channelsList {
		fmt.Printf("Connecting to channel %s\n", channel)
		connectedChannel, err := api.JoinChannel(channel)
		if err == nil {
			channels = append(channels, connectedChannel)
		} else {
			fmt.Printf(err.Error())
		}
	}

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch event := msg.Data.(type) {
		case *slack.MessageEvent:
			fmt.Printf("Message: %v\n", event.Msg.Text)

			//go run(rtm, event)
			//rtm.SendMessage(rtm.NewOutgoingMessage("Hello, World.", "channel_id"))
			break
		default:
			fmt.Printf("Unhandled event type: %v\n", event)
		}
	}
}

func main() {
	mapCommands()
	botStart()
}
