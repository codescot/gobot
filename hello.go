package main

import irc "github.com/thoj/go-ircevent"

// HelloCommand hello, world command
type HelloCommand struct{}

func (hello HelloCommand) execute(ircobj *irc.Connection, event *irc.Event) {
	sender := event.Nick
	messageChannel := event.Arguments[0]

	ircobj.Privmsg(messageChannel, sender+": Hello, Go!")
}
