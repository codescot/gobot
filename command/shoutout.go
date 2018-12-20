package command

import (
	"fmt"
	"strings"
)

// Shoutout the shoutout command
type Shoutout struct{}

// Execute run command
func (Shoutout) Execute(resp Response, event MessageEvent) {
	user := event.Message

	if strings.HasPrefix(user, "@") {
		user = user[1:]
	}

	resp(fmt.Sprintf("Check out https://twitch.tv/%s", user))
}
