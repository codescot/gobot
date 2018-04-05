package testing

import (
	"testing"

	"github.com/gurparit/slackbot/command"
)

func TestHelloGoSuccess(test *testing.T) {
	helloCommand := command.HelloCommand{}

	helloCommand.Execute(func(response string) {
		if response != "Hello, Go!" {
			test.Errorf("Hello, Go! Command Failed")
		}
	}, "")
}
