package command

// EchoCommand the Echo class
type EchoCommand struct{}

// Execute Echo implementation
func (echo EchoCommand) Execute(respond func(string), query string) {
	respond(query)
}
