package command

// Echo the Echo class
type Echo struct{}

// Execute run command
func (Echo) Execute(resp Response, event MessageEvent) {
	resp(event.Message)
}
