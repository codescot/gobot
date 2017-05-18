package main

import (
	"encoding/json"
	"io/ioutil"
)

// Config jenna.conf file.
type Config struct {
	IRCServer   string   `json:"server"`
	IRCUsername string   `json:"username"`
	IRCPassword string   `json:"password"`
	IRCChannels []string `json:"channels"`

	GoogleAPI string `json:"google_api"`
	GoogleCX  string `json:"google_cx"`
}

var config = loadConfig()

func loadConfig() Config {
	config := Config{}

	contents, error := ioutil.ReadFile("./jenna.conf")
	handleError(error)

	json.Unmarshal(contents, &config)

	return config
}
