package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/gurparit/go-common/env"
	irc "github.com/gurparit/go-ircevent"
	"github.com/gurparit/twitchbot/command"
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

	if c, ok := functions[action]; ok {
		c.Execute(bot, query)
	}
}

func botStart(debug bool) {
	username := command.ENV.Username
	channelID := command.ENV.TwitchChannelID

	ircobj := irc.IRC(username, username)
	ircobj.Password = command.ENV.Password

	ircobj.UseTLS, _ = strconv.ParseBool(command.ENV.UseTLS)
	ircobj.Debug = debug

	ircobj.AddCallback("001", func(e *irc.Event) {
		ircobj.Join(channelID)
	})
	ircobj.AddCallback("PRIVMSG", func(event *irc.Event) {
		message := event.Message()
		if strings.HasPrefix(message, "!") {
			go run(func(response string) {
				ircobj.Privmsg(channelID, response)
			}, message)
		}
	})

	ircobj.Connect(command.ENV.TwitchURL)
	ircobj.Nick(username)
	ircobj.Loop()
}

func main() {
	debug := *flag.Bool("debug", false, "-debug=true")

	env.Read(&command.ENV)

	mapCommands()
	botStart(debug)
}
