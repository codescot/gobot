package command

import (
	"time"
)

// Time default time command
type Time struct{}

// Execute run command
func (Time) Execute(r Response, message string) {
	r(time.Now().Format(time.RFC850))
}
