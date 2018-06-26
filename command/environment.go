package command

type Environment struct {
	Slack string `env:"SLACK_USER_TOKEN"`
	Bot   string `env:"BOT_USER_TOKEN"`

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
