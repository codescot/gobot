package command

// Echo the Echo class
type Echo struct{}

// Execute run command
func (Echo) Execute(r Response, query string) {
	r(query)
}
