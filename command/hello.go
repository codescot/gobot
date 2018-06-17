package command

// Hello hello, world command
type Hello struct{}

// Execute Hello implementation
func (Hello) Execute(r Response, message string) {
	r("Hello, Go!")
}
