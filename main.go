package main

import (
	"github.com/gurparit/go-common/env"
	"github.com/gurparit/twitchbot/conf"
	"github.com/gurparit/twitchbot/twitch"
)

func main() {
	env.Read(&conf.ENV)

	twitch.Go()
}
