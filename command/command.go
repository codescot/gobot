package command

type Response func(string)

type Command interface {
	Execute(Response, string)
}
