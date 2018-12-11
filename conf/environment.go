package conf

// Environment contains a key/value of environment variable.
type Environment struct {
	TwitchURL string `env:"TWITCHBOT_URL"`
	Username  string `env:"TWITCHBOT_USERNAME"`
	Password  string `env:"TWITCHBOT_PASSWORD"`
	UseTLS    string `env:"TWITCHBOT_USE_TLS"`

	TwitchChannelID string `env:"TWITCHBOT_CHANNEL_ID"`

	UrbanURL   string `env:"URBAN_URL"`
	YoutubeURL string `env:"YOUTUBE_SEARCH_URL"`

	GoogleURL      string `env:"GOOGLE_SEARCH_URL"`
	GoogleKey      string `env:"GOOGLE_API_KEY"`
	GoogleSearchID string `env:"GOOGLE_APP_ID"`

	GiphyURL string `env:"GIPHY_URL"`
	GiphyKey string `env:"GIPHY_API_KEY"`

	OxfordURL   string `env:"OXFORD_URL"`
	OxfordAppID string `env:"OXFORD_APP_ID"`
	OxfordKey   string `env:"OXFORD_API_KEY"`
}
