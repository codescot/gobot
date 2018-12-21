package conf

// Keys contains a key/value of environment variable.
type Keys struct {
	GoogleSearchID string `env:"GOOGLE_APP_ID"`
	GoogleKey      string `env:"GOOGLE_API_KEY"`

	OxfordAppID string `env:"OXFORD_APP_ID"`
	OxfordKey   string `env:"OXFORD_API_KEY"`

	GiphyKey string `env:"GIPHY_API_KEY"`
}
