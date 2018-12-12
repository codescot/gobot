package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/gurparit/go-common/env"
	"github.com/gurparit/go-common/httpc"
	irc "github.com/gurparit/go-ircevent"
	"github.com/gurparit/twitchbot/command"
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

func botStart() {
	username := command.ENV.Username
	channelID := command.ENV.TwitchChannelID

	password := "oauth:" + twitchAuth.AccessToken

	ircobj := irc.IRC(username, username)
	ircobj.UseTLS = true
	ircobj.Debug = false
	ircobj.Password = password

	ircobj.AddCallback("001", func(event *irc.Event) {
		ircobj.Join(channelID)
	})

	ircobj.AddCallback("PRIVMSG", func(event *irc.Event) {
		go func(event *irc.Event) {
			fmt.Printf("[TwitchBot] %s: %s\n", event.Nick, event.Message())
			message := event.Message()
			if strings.HasPrefix(message, "!") {
				go run(func(response string) {
					ircobj.Privmsg(event.Arguments[0], response)
				}, message)
			}
		}(event)
	})

	fmt.Println("[TwitchBot] Connect to IRC")
	ircobj.Connect(command.ENV.TwitchURL)
	ircobj.Loop()
}

var server *http.Server

// TwitchAuth holds all the properties from the authorised twitch response
type TwitchAuth struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	ExpiresIn    int      `json:"expires_in"`
	Scope        []string `json:"scope"`
	TokenType    string   `json:"token_type"`
}

var twitchAuth = TwitchAuth{}

func twitchOauthStart() {
	targetURL := fmt.Sprintf("https://id.twitch.tv/oauth2/authorize?client_id=%s&redirect_uri=http://localhost:8080/oauth2&response_type=code&scope=chat:read%%20chat:edit&state=1234", command.ENV.TwitchClientID)

	openbrowser(targetURL)
	callbackListen()
}

func openbrowser(url string) {
	fmt.Println("[TwitchBot] Open Browser on " + runtime.GOOS)
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func twitchCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[TwitchBot] OAuth2 Callback Received")

	values := r.URL.Query()

	state := values.Get("state")
	code := values.Get("code")

	if state == "1234" {
		tokenQuery := httpc.HTTP{
			TargetURL: "https://id.twitch.tv/oauth2/token",
			Method:    http.MethodPost,
			Form: map[string]string{
				"client_id":     command.ENV.TwitchClientID,
				"client_secret": command.ENV.TwitchClientSecret,
				"code":          code,
				"grant_type":    "authorization_code",
				"redirect_uri":  "http://localhost:8080/oauth2",
			},
		}

		fmt.Println("[TwitchBot] Requesting Access Token")
		err := tokenQuery.JSON(&twitchAuth)
		if err != nil {
			w.WriteHeader(403)
			log.Fatal(err)
		}

		fmt.Println("[TwitchBot] Authenticated")
		fmt.Println("[TwitchBot] Scope: " + strings.Join(twitchAuth.Scope, " "))
		fmt.Println("[TwitchBot] Expiry: " + time.Now().Add(time.Duration(twitchAuth.ExpiresIn)*time.Second).Format(time.RFC3339))
		fmt.Println("[TwitchBot] Starting Bot")

		w.Write([]byte("Login Successful, you may now close this tab."))

		go botStart()
	} else {
		w.WriteHeader(500)
	}
}

func callbackListen() {
	fmt.Println("[TwitchBot] Start Callback Listener on http://localhost:8080/oauth2")

	http.HandleFunc("/oauth2", twitchCallback)
	server = &http.Server{Addr: ":8080"}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal()
	}
}

func main() {
	fmt.Println("[TwitchBot] Start: Reading Environment Variables")
	env.Read(&command.ENV)

	fmt.Println("[TwitchBot] Mapping Commands")
	mapCommands()

	fmt.Println("[TwitchBot] Starting Twitch OAuth")
	twitchOauthStart()
}
