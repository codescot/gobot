package bot

import (
	"fmt"
	"github.com/gurparit/go-common/env"
	"github.com/gurparit/go-ircbot/conf"
	"strings"

	"github.com/gurparit/go-ircbot/command"
	irc "github.com/gurparit/go-ircevent"
)

// Bot irc bot object
type Bot struct {
	keys    *conf.Keys
	config *Config
	conn   *irc.Connection
}

// Config irc bot config
type Config struct {
	Server   string
	Username string
	Password string
	UseTLS   bool
	Debug    bool

	Channels []string

	ExtendedCommands bool

	MessageListeners chan string
}

var functions = make(map[string]command.Command)

func defaultCommands() {
	addCommand("echo", command.Echo{})
	addCommand("time", command.Time{})
	addCommand("go", command.Hello{})
	addCommand("so", command.Shoutout{})
}

func extendedCommands() {
	addCommand("g", command.Google{})
	addCommand("ud", command.Urban{})
	addCommand("yt", command.Youtube{})
	addCommand("gif", command.Giphy{})
	addCommand("define", command.Oxford{})
	addCommand("ety", command.Oxford{Etymology: true})
}

func addCommand(key string, cmd command.Command) {
	functions["!"+key] = cmd
}

func recovery() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

func (*Bot) onNewMessage(response command.Response, event command.MessageEvent) {
	defer recovery()

	params := strings.SplitN(event.Message, " ", 2)
	action := params[0]
	query := ""

	if len(params) > 1 {
		query = params[1]
	}

	event.Message = query

	if c, ok := functions[action]; ok {
		c.Execute(response, event)
	}
}

func (bot *Bot) onWelcomeEvent(channels []string) func(*irc.Event) {
	return func(event *irc.Event) {
		for _, channel := range channels {
			bot.conn.Join(channel)
		}
	}
}

func (bot *Bot) onMessageEvent(event *irc.Event) {
	channel := event.Arguments[0]
	user := event.Nick
	message := event.Message()

	events := bot.config.MessageListeners
	if events != nil {
		events <- fmt.Sprintf("{ \"username\": \"%s\", \"message\": \"%s\" }", user, message)
	}

	if strings.HasPrefix(message, "!") {
		go bot.onNewMessage(bot.onResponseEvent(channel), command.MessageEvent{
			Channel:  channel,
			Username: user,
			Message:  message,
			Keys:     bot.keys,
		})
	}
}

func (bot *Bot) onResponseEvent(channel string) command.Response {
	return func(response string) {
		bot.conn.Privmsg(channel, response)
	}
}

// New returns a new IRCBot with supplied config
func New(cfg *Config) *Bot {
	return &Bot{config: cfg}
}

// Default returns a new IRCBot with default config
func Default(server, username, password string) *Bot {
	return New(&Config{
		Server:   server,
		Username: username,
		Password: password,
		UseTLS:   false,
		Debug:    true,
		Channels: []string{"#general"},

		ExtendedCommands: false,
	})
}

// DefaultTLS returns a new IRCBot with default config and TLS enabled
func DefaultTLS(server, username, password string) *Bot {
	return New(&Config{
		Server:   server,
		Username: username,
		Password: password,
		UseTLS:   true,
		Debug:    true,
		Channels: []string{"#general"},

		ExtendedCommands: false,
	})
}

const (
	// EventWelcome callback key for welcome event
	EventWelcome = "001"
	// EventPrivateMessage callback key for private message event
	EventPrivateMessage = "PRIVMSG"
)

// Start bot start
func (bot *Bot) Start() {
	username := bot.config.Username

	conn := irc.IRC(username, username)
	conn.UseTLS = bot.config.UseTLS
	conn.Debug = bot.config.Debug
	conn.Password = bot.config.Password

	bot.conn = conn

	conn.AddCallback(EventWelcome, bot.onWelcomeEvent(bot.config.Channels))
	conn.AddCallback(EventPrivateMessage, bot.onMessageEvent)

	defaultCommands()
	if bot.config.ExtendedCommands {
		keys := conf.Keys{}
		env.Read(&keys)

		bot.keys = &keys
		extendedCommands()
	}

	conn.Connect(bot.config.Server)
	conn.Loop()
}

//func main() {
//	irc := bot.Default("irc.example.com", "username", "password")
//	irc.Start()
//}
