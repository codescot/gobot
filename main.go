package main

import (
	"github.com/gurparit/go-common/env"
	"github.com/gurparit/twitchbot/command"
	"github.com/gurparit/twitchbot/twitch"
)

func main() {
	env.Read(&command.ENV)

	twitch.Go()
}
