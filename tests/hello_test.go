package testing

import (
	"testing"

	"github.com/gurparit/gobot/command"
)

func TestHelloGoSuccess(test *testing.T) {
	h := command.Hello{}

	h.Execute(func(response string) {
		if response != "Hello, Go!" {
			test.Errorf("Hello, Go! Command Failed")
		}
	}, "")
}
