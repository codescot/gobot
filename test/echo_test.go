package test

import (
	"testing"

	"github.com/gurparit/twitchbot/command"
)

func TestEchoSuccess(t *testing.T) {
	echo := command.Echo{}

	expectedResponse := "Hi, my name is echo."

	echo.Execute(func(response string) {
		if response != expectedResponse {
			t.Log(response)
			t.Fail()
		}
	}, "Hi, my name is echo.")
}
