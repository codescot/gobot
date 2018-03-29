package util

import (
	"encoding/json"
	"io/ioutil"
)

// Config jenna.conf file.
type Config struct {
	SlackToken string   `json:"slack_token"`
	Channels   []string `json:"channels"`
	Debug      bool     `json:"debug"`

	GoogleAPI string `json:"google_api"`
	GoogleCX  string `json:"google_cx"`

	OxfordID  string `json:"oxford_app_id"`
	OxfordKey string `json:"oxford_app_key"`
}

// Marbles the configuration
var Marbles = loadConfig()

func loadConfig() Config {
	config := Config{}

	contents, error := ioutil.ReadFile("./marbles.conf")
	if IsError(error) {
		return config
	}

	json.Unmarshal(contents, &config)

	return config
}
