package command

import (
	"fmt"

	"github.com/codescot/go-common/httputil"
)

// Uptime the shoutout command
type Uptime struct{}

// Execute run command
func (Uptime) Execute(resp Response, event MessageEvent) {
	channel := event.Channel
	if channel[0] == '#' {
		channel = channel[1:]
	}

	req := httputil.HTTP{
		TargetURL: fmt.Sprintf("https://decapi.me/twitch/uptime/%s", channel),
		Method:    "GET",
	}

	uptime, _ := req.String()

	resp(uptime)
}
