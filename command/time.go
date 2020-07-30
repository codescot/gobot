package command

import (
	"time"
)

// Time default time command
type Time struct{}

func (Time) CanExecute(MessageEvent) bool {
	return true
}

// Execute run command
func (Time) Execute(resp Response, event MessageEvent) {
	resp(time.Now().Format(time.RFC850))
}
