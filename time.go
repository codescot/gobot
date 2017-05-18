package main

import (
	"time"

	irc "github.com/thoj/go-ircevent"
)

// TimeCommand server time command
type TimeCommand struct{}

func (timeCommand TimeCommand) execute(ircobj *irc.Connection, event *irc.Event) {
	messageChannel := event.Arguments[0]

	ircobj.Privmsg(messageChannel, time.Now().Format(time.RFC850))
}
