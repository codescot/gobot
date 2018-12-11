package test

import (
	"testing"
	"time"

	"github.com/gurparit/twitchbot/command"
)

func TestTimeSuccess(t *testing.T) {
	tc := command.Time{}

	tc.Execute(func(response string) {
		if response != time.Now().Format(time.RFC850) {
			t.Log(response)
			t.Fail()
		}
	}, "")
}
