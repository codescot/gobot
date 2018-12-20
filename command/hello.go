package command

import (
	"fmt"
)

// Hello hello, world command
type Hello struct{}

// Execute Hello implementation
func (Hello) Execute(resp Response, event MessageEvent) {
	resp(fmt.Sprintf("Hello, %s!", event.Username))
}
