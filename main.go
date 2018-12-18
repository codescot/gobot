package main

import (
	"github.com/gurparit/go-common/env"
	"github.com/gurparit/go-irc-bot/conf"
	"github.com/gurparit/go-irc-bot/twitch"
)

func main() {
	env.Read(&conf.ENV)

	twitch.Go()
}
