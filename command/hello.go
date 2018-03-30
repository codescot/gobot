package command

// HelloCommand hello, world command
type HelloCommand struct{}

// Execute HelloCommand implementation
func (hello HelloCommand) Execute(respond func(string), message string) {
	respond("Hello, Go!")
}
