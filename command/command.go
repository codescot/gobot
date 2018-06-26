package command

import (
	"fmt"
	"strings"

	"github.com/gurparit/gobot/conf"
)

var OS conf.Environment

type Response func(string)

type Command interface {
	Execute(Response, string)
}

func FormatURL(baseURL string, params ...interface{}) string {
	if strings.Contains(baseURL, "%s") {
		return fmt.Sprintf(baseURL, params...)
	}

	return baseURL
}
