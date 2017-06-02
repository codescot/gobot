package command

import (
	"strings"

	irc "github.com/thoj/go-ircevent"
)

// EchoCommand the Echo class
type EchoCommand struct{}

// Execute Echo implementation
func (echo EchoCommand) Execute(ircobj *irc.Connection, event *irc.Event) {
	messages := strings.SplitN(event.Message(), " ", 3)
	messageChannel := messages[1]
	text := messages[2]

	ircobj.Privmsg(messageChannel, text)
}
