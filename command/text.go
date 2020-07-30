package command

// Text basic command wrapper
type Text struct {
	Perm string `json:"perm"`
	Name string `json:"name"`
	Text string `json:"text"`
}

func (b Text) CanExecute(event MessageEvent) bool {
	return HasPerm(b.Perm, event.IsSub, event.IsMod)
}

// Execute run command
func (b Text) Execute(resp Response, event MessageEvent) {
	resp(event.Format(b.Text))
}
