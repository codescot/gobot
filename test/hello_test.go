package test

import (
	"testing"

	"github.com/gurparit/twitchbot/command"
)

func TestHelloGoSuccess(t *testing.T) {
	hello := command.Hello{}

	hello.Execute(func(response string) {
		if response != "Hello, Go!" {
			t.Log(response)
			t.Fail()
		}
	}, "")
}
