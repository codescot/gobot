package command

// Echo the Echo class
type Echo struct{}

// Execute Echo implementation
func (Echo) Execute(r Response, query string) {
	r(query)
}
