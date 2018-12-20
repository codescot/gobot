package main

import (
	"github.com/gurparit/go-irc-bot/bot"
)

func main() {
	irc := bot.Default("irc.example.com", "username", "password")
	irc.Start()
}
