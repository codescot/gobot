package command

import (
	"github.com/gurparit/gobot/conf"
)

var OS conf.Environment

type Response func(string)

type Command interface {
	Execute(Response, string)
}
