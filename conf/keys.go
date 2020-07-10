package conf

// Keys contains a key/value of environment variable.
type Keys struct {
	GoogleSearchID string `env:"GOOGLE_APP_ID" json:"googleAppId"`
	GoogleKey      string `env:"GOOGLE_API_KEY" json:"googleApiKey"`

	OxfordAppID string `env:"OXFORD_APP_ID" json:"oxfordAppId"`
	OxfordKey   string `env:"OXFORD_API_KEY" json:"oxfordApiKey"`

	GiphyKey string `env:"GIPHY_API_KEY" json:"giphyApiKey"`
}
