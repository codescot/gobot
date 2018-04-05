package testing

import (
	"testing"

	"github.com/gurparit/slackbot/command"
)

func TestEchoSuccess(test *testing.T) {
	echoCommand := command.EchoCommand{}

	expectedResponse := "Hi, my name is echo."

	echoCommand.Execute(func(response string) {
		if response != expectedResponse {
			test.Errorf("Echo Command Failed: expecting %s but was %s", expectedResponse, response)
		}
	}, "!echo Hi, my name is echo.")
}
