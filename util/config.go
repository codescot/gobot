package util

import (
	"encoding/json"
	"io/ioutil"
)

// ConfigFile slackbot.conf file.
type ConfigFile struct {
	Username   string   `json:"username"`
	SlackToken string   `json:"slack_token"`
	Channels   []string `json:"channels"`
	Debug      bool     `json:"debug"`

	GoogleAPI string `json:"google_api"`
	GoogleCX  string `json:"google_cx"`

	GiphyAPI string `json:"giphy_api"`

	OxfordID  string `json:"oxford_app_id"`
	OxfordKey string `json:"oxford_app_key"`
}

// Config the configuration
var Config = loadConfig()

func loadConfig() ConfigFile {
	config := ConfigFile{}

	contents, error := ioutil.ReadFile("./slackbot.conf")
	if IsError(error) {
		return config
	}

	json.Unmarshal(contents, &config)

	return config
}
