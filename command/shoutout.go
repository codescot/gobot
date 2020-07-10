package command

import (
	"fmt"
	"strings"

	"github.com/codescot/go-common/array"
	"github.com/codescot/go-common/httputil"
)

// Shoutout the shoutout command
type Shoutout struct {
	Team []string
}

// Execute run command
func (so Shoutout) Execute(resp Response, event MessageEvent) {
	if !event.IsModerator {
		return
	}

	tokens := strings.SplitN(event.Message, " ", 2)
	user := tokens[0]
	if user[0] == '@' {
		user = user[1:]
	}

	body := ""
	if len(so.Team) > 0 && array.Contains(so.Team, strings.ToLower(user)) {
		body = fmt.Sprintf("/me ▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬ Fellow Nook member [[ %s ]] https://twitch.tv/%s bleedPurple", user, user)
	} else {
		body = fmt.Sprintf("/me ▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬ Go follow the wonderful [[ %s ]] https://twitch.tv/%s bleedPurple", user, user)
	}

	req := httputil.HTTP{
		TargetURL: fmt.Sprintf("https://decapi.me/twitch/game/%s", user),
		Method:    "GET",
	}

	if r, err := req.String(); err == nil {
		game := fmt.Sprintf(" ▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬ Last Seen Playing: \"%s\" ▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬", r)
		body += game
	}

	resp(event.Format(body))
}
