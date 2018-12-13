package conf

// ENV contains the preloaded environment variables
var ENV Environment

// Environment contains a key/value of environment variable.
type Environment struct {
	TwitchURL          string `env:"TWITCH_URL"`
	TwitchClientID     string `env:"TWITCH_CLIENT_ID"`
	TwitchClientSecret string `env:"TWITCH_CLIENT_SECRET"`

	Username        string `env:"TWITCH_USERNAME"`
	TwitchChannelID string `env:"TWITCH_CHANNEL_ID"`
	TwitchScope     string `env:"TWITCH_SCOPE"`

	GoogleKey      string `env:"GOOGLE_API_KEY"`
	GoogleSearchID string `env:"GOOGLE_APP_ID"`

	GiphyKey string `env:"GIPHY_API_KEY"`

	OxfordAppID string `env:"OXFORD_APP_ID"`
	OxfordKey   string `env:"OXFORD_API_KEY"`
}
