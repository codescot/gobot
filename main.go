package main

import (
	"github.com/gurparit/go-ircbot/bot"
)

func main() {
	irc := bot.Default("irc.example.com", "username", "password")
	irc.Start()
}
