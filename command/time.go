package command

import (
	"time"
)

// TimeCommand server time command
type TimeCommand struct{}

// Execute TimeCommand implementation
func (timeCommand TimeCommand) Execute(respond func(string), message string) {
	respond(time.Now().Format(time.RFC850))
}
