package testing

import (
	"testing"

	"github.com/gurparit/gobot/command"
)

func TestEchoSuccess(test *testing.T) {
	echoCommand := command.Echo{}

	expectedResponse := "Hi, my name is echo."

	echoCommand.Execute(func(response string) {
		if response != expectedResponse {
			test.Errorf("Echo Command Failed: expecting %s but was %s", expectedResponse, response)
		}
	}, "Hi, my name is echo.")
}
