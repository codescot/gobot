package command

// Response a response wrapper
type Response func(string)

// Command basic command interface
type Command interface {
	Execute(Response, string)
}
