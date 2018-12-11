package command

import (
	"github.com/gurparit/twitchbot/conf"
)

// ENV contains the preloaded environment variables
var ENV conf.Environment

// Response a response wrapper
type Response func(string)

// Command basic command interface
type Command interface {
	Execute(Response, string)
}
