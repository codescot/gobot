package twitch

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gurparit/go-common/httpc"
	"github.com/gurparit/twitchbot/conf"
	"github.com/gurparit/twitchbot/core"
)

type authentication struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	ExpiresIn    int      `json:"expires_in"`
	Scope        []string `json:"scope"`
	TokenType    string   `json:"token_type"`
}

const oAuthBaseURL = "https://id.twitch.tv/oauth2/authorize?client_id=%s&redirect_uri=http://localhost:8080/oauth2&response_type=code&scope=chat:read%%20chat:edit&state=1234"

var auth = authentication{}

func callbackReceived() core.CallbackHandler {
	return func(state, code string) (int, error) {
		// Get authentication token from auth code
		tokenQuery := httpc.HTTP{
			TargetURL: "https://id.twitch.tv/oauth2/token",
			Method:    http.MethodPost,
			Form: map[string]string{
				"client_id":     conf.ENV.TwitchClientID,
				"client_secret": conf.ENV.TwitchClientSecret,
				"code":          code,
				"grant_type":    "authorization_code",
				"redirect_uri":  "http://localhost:8080/oauth2",
			},
		}

		err := tokenQuery.JSON(&auth)
		if err != nil {
			return 403, err
		}

		joinChat()

		return 200, nil
	}
}

func joinChat() {
	bot := core.Bot{}

	server := conf.ENV.TwitchURL
	username := conf.ENV.Username
	password := "oauth:" + auth.AccessToken
	channel := conf.ENV.TwitchChannelID

	durationInSeconds := time.Duration(auth.ExpiresIn) * time.Second
	expiration := time.Now().Add(durationInSeconds).Format(time.RFC3339)

	fmt.Println("[Twitch] Scope: " + strings.Join(auth.Scope, " "))
	fmt.Println("[Twitch] Expiry: " + expiration)

	go bot.Start(server, username, password, channel)
}

// Go start the Twitch Bot application
func Go() {
	web := core.Web{}

	targetURL := fmt.Sprintf(oAuthBaseURL, conf.ENV.TwitchClientID)

	web.OpenBrowser(targetURL)
	web.Start(callbackReceived())
}
