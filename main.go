package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/codescot/go-common/array"
	"github.com/codescot/go-common/fileio"

	irc "github.com/codescot/go-ircevent"
	"github.com/codescot/gobot/command"
	"github.com/codescot/gobot/filter"
)

// Bot irc bot object
type Bot struct {
	IRC      IRC `json:"bot"`
	Commands map[string]command.Text

	Team     []string
	Blocked  []string
	BadWords []string

	conn *irc.Connection
}

// IRC irc configuration
type IRC struct {
	Server   string
	Username string
	Password string
	TLS      bool
	Debug    bool

	Channels        []string
	RequestCaps     []string
	AcknowledgeCaps []string
}

var functions = make(map[string]command.Command)

func addCommand(key string, cmd command.Command) {
	if t, ok := cmd.(command.Text); ok {
		bot.Commands[key] = t
	}
	functions[key] = cmd
}

func disableCommand(key string) {
	delete(functions, key)
}

func deleteCommand(key string) {
	disableCommand(key)
	delete(bot.Commands, key)
}

func enableCommand(key string) {
	functions[key] = bot.Commands[key]
}

func recovery() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

func (*Bot) processCommand(response command.Response, event command.MessageEvent) {
	defer recovery()

	params := strings.SplitN(event.Message, " ", 2)
	action := params[0][1:] // [1:] removes the !
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
		if len(bot.IRC.RequestCaps) > 0 {
			bot.startCaps()
		}

		for _, channel := range channels {
			bot.conn.Join(channel)
		}
	}
}

func (bot *Bot) onMessageEvent(event *irc.Event) {
	channel := event.Arguments[0]
	user := event.Nick
	message := event.Message()
	tags := event.Tags

	// get message tags and assign appropriately
	messageID := event.Tags["id"]
	badges := strings.Split(event.Tags["badges"], ",")
	isModerator := array.DeepContains(badges, "moderator") || array.DeepContains(badges, "broadcaster")
	isSubscriber := array.DeepContains(badges, "subscriber")

	messageEvent := command.MessageEvent{
		MessageID:    messageID,
		Channel:      channel,
		Username:     user,
		Message:      message,
		IsModerator:  isModerator,
		IsSubscriber: isSubscriber,
		Tags:         tags,
	}

	deleteHandler := bot.DeleteHandler(channel, messageID)
	banHandler := bot.BanHandler(channel, user)
	if bot.moderate(deleteHandler, banHandler, user, message) {
		return
	}

	if strings.HasPrefix(message, "!") {
		responseHandler := bot.ResponseHandler(channel)
		go bot.processCommand(responseHandler, messageEvent)
	}
}

func (bot *Bot) moderate(delete filter.DeleteHandler, ban filter.BanHandler, username string, message string) bool {
	fs := []filter.Filter{
		filter.Domain{},
		filter.Usernames{
			Blocked:  bot.Blocked,
			Username: username,
		},
		filter.BadWords{
			BadWords: bot.BadWords,
		},
	}

	for _, f := range fs {
		switch f.Apply(message) {
		case filter.Delete:
			delete()
			return true
		case filter.Ban:
			ban()
			return true
		}
	}

	return false
}

// DeleteHandler a handle for deleting the last message if filters match
func (bot *Bot) DeleteHandler(channel string, messageID string) filter.DeleteHandler {
	return func() {
		bot.conn.Privmsg(channel, fmt.Sprintf("/delete %s", messageID))
	}
}

// BanHandler a handle for banning a user
func (bot *Bot) BanHandler(channel string, username string) filter.BanHandler {
	return func() {
		bot.conn.Privmsg(channel, fmt.Sprintf("/ban %s", username))
	}
}

// ResponseHandler a handle for executing a command
func (bot *Bot) ResponseHandler(channel string) command.Response {
	return func(response string) {
		bot.conn.Privmsg(channel, response)
	}
}

const (
	// Welcome welcome
	Welcome = "001"
	// PrivateMessage private message
	PrivateMessage = "PRIVMSG"
	// JOIN message
	JOIN = "JOIN"
	// Cap cap
	Cap = "CAP"
	// CapLS cap ls
	CapLS = "CAP LS"
	// CapEnd cap end
	CapEnd = "CAP END"
	// LS ls
	LS = "LS"
	// ACK ack
	ACK = "ACK"
)

func (bot *Bot) startCaps() {
	bot.conn.SendRaw(CapLS)
}

func (bot *Bot) onCapEvent(event *irc.Event) {
	if event.Arguments[1] == LS {
		bot.nextCap()
	}

	if event.Arguments[1] == ACK {
		capCount := len(bot.IRC.RequestCaps)
		switch {
		case capCount == 0:
			bot.endCap()
		case capCount > 0:
			bot.nextCap()
		}
	}
}

func (bot *Bot) nextCap() {
	cap, newCaps := bot.IRC.RequestCaps[0], bot.IRC.RequestCaps[1:]

	bot.IRC.RequestCaps = newCaps
	bot.IRC.AcknowledgeCaps = append(bot.IRC.AcknowledgeCaps, cap)

	bot.conn.SendRawf("CAP REQ :%s", cap)
}

func (bot *Bot) endCap() {
	bot.conn.SendRaw(CapEnd)

	time.Sleep(2 * time.Second)
	bot.conn.Privmsg("#GeeScot", "Hello, World.\n")

	bot.IRC.RequestCaps = bot.IRC.AcknowledgeCaps
	bot.IRC.AcknowledgeCaps = []string{}
}

func initCommands(bot Bot) {
	addCommand("time", command.Time{})
	addCommand("uptime", command.Uptime{})
	addCommand("so", command.Shoutout{
		Team: bot.Team,
	})

	for key, c := range bot.Commands {
		addCommand(key, c)
	}
}

var bot Bot

func startBot() {
	bot = Bot{}
	fileio.ReadJSON("config.json", &bot)

	username := bot.IRC.Username

	conn := irc.IRC(username, username)
	conn.UseTLS = bot.IRC.TLS
	conn.Debug = bot.IRC.Debug
	conn.Password = bot.IRC.Password

	bot.conn = conn

	conn.AddCallback(Welcome, bot.onWelcomeEvent(bot.IRC.Channels))
	conn.AddCallback(PrivateMessage, bot.onMessageEvent)
	conn.AddCallback(Cap, bot.onCapEvent)

	initCommands(bot)

	fmt.Println("connecting...")
	conn.Connect(bot.IRC.Server)

	conn.Loop()
}

type messageRequest struct {
	Channel string
	Message string
}

func sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var m messageRequest
	json.NewDecoder(r.Body).Decode(&m)

	bot.conn.Privmsg(m.Channel, m.Message)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "ok",
	})
}

func main() {
	startBot()
}
