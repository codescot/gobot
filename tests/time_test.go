package testing

import (
	"testing"
	"time"

	"github.com/gurparit/gobot/command"
)

func TestTimeSuccess(test *testing.T) {
	t := command.Time{}

	t.Execute(func(response string) {
		if response != time.Now().Format(time.RFC850) {
			test.Errorf("Time Command Failed")
		}
	}, "")
}
