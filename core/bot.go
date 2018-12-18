package core

import (
	"fmt"
	"strings"

	"github.com/gurparit/go-irc-bot/command"
	irc "github.com/gurparit/go-ircevent"
)

// Bot default bot object
type Bot struct{}

var functions = make(map[string]command.Command)

func mapCommands() {
	functions["!go"] = command.Hello{}
	functions["!so"] = command.Shoutout{}
	functions["!time"] = command.Time{}
	functions["!g"] = command.Google{}
	functions["!ud"] = command.Urban{}
	functions["!echo"] = command.Echo{}
	functions["!yt"] = command.Youtube{}
	functions["!gif"] = command.Giphy{}
	functions["!define"] = command.Oxford{}
	functions["!ety"] = command.Oxford{Etymology: true}
}

func recovery() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

func (*Bot) onNewMessage(bot func(string), message string) {
	defer recovery()

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

// Start bot start
func (bot *Bot) Start(server, username, password, channel string) {
	ircobj := irc.IRC(username, username)
	ircobj.UseTLS = true
	ircobj.Debug = false
	ircobj.Password = password

	ircobj.AddCallback("001", func(event *irc.Event) {
		ircobj.Join(channel)
	})

	ircobj.AddCallback("PRIVMSG", func(event *irc.Event) {
		go func(event *irc.Event) {
			fmt.Printf("%s: %s\n", event.Nick, event.Message())
			message := event.Message()
			if strings.HasPrefix(message, "!") {
				go bot.onNewMessage(func(response string) {
					ircobj.Privmsg(event.Arguments[0], response)
				}, message)
			}
		}(event)
	})

	mapCommands()
	ircobj.Connect(server)
	ircobj.Loop()
}
