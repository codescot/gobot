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
	functions["!gif"] = command.GiphyCommand{}
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

	parameters := strings.SplitN(message, " ", 2)
	action := parameters[0]
	query := ""

	if len(parameters) > 1 {
		query = parameters[1]
	}

	if command, ok := functions[action]; ok {
		command.Execute(bot, query)
	}
}

func botStart() {
	config := util.Config
	username := config.Username

	client := slack.New(config.SlackToken)
	client.SetDebug(config.Debug)

	bot := slack.New(config.BotUserToken)
	bot.SetDebug(config.Debug)

	rtm := bot.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch event := msg.Data.(type) {
		case *slack.MessageEvent:
			if event.Msg.Username != username {
				go run(func(response string) {
					bot.PostMessage(event.Msg.Channel, response, slack.NewPostMessageParameters())
				}, event.Msg.Text)
			}

			break
		case *slack.ReactionAddedEvent:
			slackMsg, err := client.GetChannelReplies(event.Item.Channel, event.Item.Timestamp)
			if err != nil {
				fmt.Println(err)
			}

			if len(slackMsg) > 0 && slackMsg[0].Username == username {
				for _, reaction := range slackMsg[0].Msg.Reactions {
					if reaction.Name == "-1" && reaction.Count > 2 {
						rtm.DeleteMessage(event.Item.Channel, event.Item.Timestamp)
						break
					}
				}
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
