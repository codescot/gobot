package main

import (
	"fmt"
	"strings"

	"github.com/thoj/go-ircevent"
)

var functions = make(map[string]Command)

// IsError generic error handling
func IsError(appError error) bool {
	if appError == nil {
		return false
	}

	fmt.Println(appError.Error())
	return true
}

func mapCommands() {
	functions["!go"] = HelloCommand{}
	functions["!time"] = TimeCommand{}
	functions["!g"] = GoogleCommand{}
	functions["!ud"] = UDCommand{}
	functions["!echo"] = EchoCommand{}
	functions["!spotify"] = SpotifyCommand{}
	functions["!yt"] = YoutubeCommand{}
}

func run(ircobj *irc.Connection, event *irc.Event) {
	message := event.Message()
	parameters := strings.Split(message, " ")
	action := parameters[0]

	if command, ok := functions[action]; ok {
		command.Execute(ircobj, event)
	}
}

func ircStart() {
	username := config.IRCUsername

	ircobj := irc.IRC(username, username)
	ircobj.Password = config.IRCPassword

	ircobj.UseTLS = true
	ircobj.Debug = false

	ircobj.AddCallback("001", func(e *irc.Event) {
		for _, channel := range config.IRCChannels {
			ircobj.Join(channel)
		}
	})
	ircobj.AddCallback("PRIVMSG", func(event *irc.Event) {
		go run(ircobj, event)
	})

	ircobj.Connect(config.IRCServer)
	ircobj.Nick(username)
	ircobj.Loop()
}

func main() {
	mapCommands()
	ircStart()
}
