package command

import (
	"strings"
)

// Shoutout the shoutout command
type Shoutout struct{}

// Execute run command
func (Shoutout) Execute(r Response, query string) {
	user := query

	if strings.HasPrefix(user, "@") {
		user = user[1:]
	}

	r("https://twitch.tv/" + user)
}
