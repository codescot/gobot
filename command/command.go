package command

import (
	"github.com/gurparit/go-ircbot/conf"
)

// Response a response wrapper
type Response func(string)

// MessageEvent sends the relevant params to the command handling the message event.
type MessageEvent struct {
	Channel  string
	Username string
	Message  string

	Keys *conf.Keys
}

// Command basic command interface
type Command interface {
	Execute(Response, MessageEvent)
}
