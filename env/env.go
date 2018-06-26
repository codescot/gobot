package env

import (
	"os"
	"reflect"
)

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

var OS Environment

func OpenConfig() Environment {
	env := Environment{}

	reflectValue := reflect.ValueOf(&env)
	reflectElem := reflectValue.Elem()
	reflectType := reflectElem.Type()

	numOfFields := reflectType.NumField()
	for i := 0; i < numOfFields; i++ {
		structField := reflectType.Field(i)
		if key, ok := structField.Tag.Lookup("env"); ok {
			fieldValue := reflectElem.FieldByName(structField.Name)
			if fieldValue.CanAddr() && fieldValue.CanSet() {
				value := os.Getenv(key)
				fieldValue.SetString(value)
			}
		}
	}

	return env
}
