package testing

import (
	"testing"
	"time"

	"github.com/gurparit/slackbot/command"
)

func TestTimeSuccess(test *testing.T) {
	timeCommand := command.TimeCommand{}

	timeCommand.Execute(func(response string) {
		if response != time.Now().Format(time.RFC850) {
			test.Errorf("Time Command Failed")
		}
	}, "")
}
