package main

import (
	"fmt"
	"strings"

	"github.com/gurparit/marbles/command"
	"github.com/gurparit/marbles/util"
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

func ircStart() {
	config := util.Marbles
	username := config.IRCUsername

	ircobj := irc.IRC(username, username)
	ircobj.Password = config.IRCPassword

	ircobj.UseTLS = config.UseTLS
	ircobj.Debug = config.Debug

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
