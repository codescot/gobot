package command

import (
	"time"
)

// Time server time command
type Time struct{}

// Execute Time implementation
func (Time) Execute(r Response, message string) {
	r(time.Now().Format(time.RFC850))
}
