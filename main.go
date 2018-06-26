package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/gurparit/gobot/command"
	"github.com/gurparit/gobot/env"
	"github.com/nlopes/slack"
)

var functions = make(map[string]command.Command)

func mapCommands() {
	functions["!go"] = command.Hello{}
	functions["!time"] = command.Time{}
	functions["!g"] = command.Google{}
	functions["!ud"] = command.Urban{}
	functions["!echo"] = command.Echo{}
	functions["!yt"] = command.Youtube{}
	functions["!gif"] = command.Giphy{}
	functions["!define"] = command.Oxford{}
	functions["!ety"] = command.Oxford{Etymology: true}
}

// CatchErrors catch all errors and recover.
func CatchErrors() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

func run(bot func(string), message string) {
	defer CatchErrors()

	params := strings.SplitN(message, " ", 2)
	action := params[0]
	query := ""

	if len(params) > 1 {
		query = params[1]
	}

	if command, ok := functions[action]; ok {
		command.Execute(bot, query)
	}
}

func botStart(debug bool, username string) {
	client := slack.New(env.OS.Slack)
	client.SetDebug(debug)

	bot := slack.New(env.OS.Bot)
	bot.SetDebug(debug)

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
	debug := *flag.Bool("debug", false, "-debug=true")
	username := *flag.String("username", "gobot", "-username=gobot")

	env.OS = env.OpenConfig()

	mapCommands()
	botStart(debug, username)
}
