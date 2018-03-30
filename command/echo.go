package command

import (
	"strings"
)

// EchoCommand the Echo class
type EchoCommand struct{}

// Execute Echo implementation
func (echo EchoCommand) Execute(respond func(string), message string) {
	messages := strings.SplitN(message, " ", 2)
	text := messages[1]

	respond(text)
}
