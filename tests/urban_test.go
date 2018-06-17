package testing

import (
	"testing"

	"github.com/gurparit/gobot/command"
)

func TestUrbanDictionarySuccess(test *testing.T) {
	u := command.Urban{}

	u.Execute(func(response string) {
		if response != "tetrible - Portmanteau of [Tetris] and terrible, for when things just don't fit." {
			test.Errorf("UD Command Failed")
		}
	}, "tetrible")
}
