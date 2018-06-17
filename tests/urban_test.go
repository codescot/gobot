package testing

import (
	"testing"

	"github.com/gurparit/gobot/command"
)

func TestUrbanDictionarySuccess(test *testing.T) {
	udCommand := command.Urban{}

	udCommand.Execute(func(response string) {
		if response != "tetrible - Portmanteau of [Tetris] and terrible, for when things just don't fit." {
			test.Errorf("UD Command Failed")
		}
	}, "tetrible")
}
