package command

import (
	"time"
)

type Time struct{}

func (Time) Execute(r Response, message string) {
	r(time.Now().Format(time.RFC850))
}
