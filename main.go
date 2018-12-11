package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/gurparit/go-common/env"
	"github.com/gurparit/twitchbot/command"
	irc "gopkg.in/irc.v3"
)

var functions = make(map[string]command.Command)

func mapCommands() {
	functions["!go"] = command.Hello{}
	functions["!time"] = command.Time{}
	functions["!g"] = command.Google{}
	functions["!ud"] = command.Urban{}
	functions["!echo"] = command.Echo{}
	functions["!yt"] = command.Youtube{}
	functions["!gif"] = command.Giphy{}
	functions["!define"] = command.Oxford{}
	functions["!ety"] = command.Oxford{Etymology: true}
}

// CatchErrors catch all errors and recover.
func CatchErrors() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

func run(bot func(string), message string) {
	defer CatchErrors()

	params := strings.SplitN(message, " ", 2)
	action := params[0]
	query := ""

	if len(params) > 1 {
		query = params[1]
	}

	if c, ok := functions[action]; ok {
		c.Execute(bot, query)
	}
}

func botStart(debug bool) {
	username := command.ENV.Username
	password := command.ENV.Password
	channelID := command.ENV.TwitchChannelID

	fmt.Printf("PASS %s\n", password)
	fmt.Printf("USER %s\n", username)
	fmt.Printf("CHAN %s\n", channelID)

	fmt.Println("Dial Connection")
	conn, err := net.Dial("tcp", command.ENV.TwitchURL)
	if err != nil {
		log.Fatalln(err)
	}

	config := irc.ClientConfig{
		Nick: username,
		Pass: password,
		User: username,
		Handler: irc.HandlerFunc(func(c *irc.Client, m *irc.Message) {
			if m.Command == "001" {
				c.Write(fmt.Sprintf("JOIN %s", channelID))
			} else if m.Command == "PRIVMSG" && c.FromChannel(m) {
				message := m.Trailing()
				if strings.HasPrefix(message, "!") {
					go run(func(response string) {
						c.WriteMessage(&irc.Message{
							Command: "PRIVMSG",
							Params: []string{
								m.Params[0],
								response,
							},
						})
					}, message)
				}
			}
		}),
	}

	fmt.Println("New Client")
	// Create the client
	client := irc.NewClient(conn, config)

	fmt.Println("Run Client")
	err = client.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	debug := *flag.Bool("debug", false, "-debug=true")

	env.Read(&command.ENV)

	mapCommands()
	botStart(debug)
}
