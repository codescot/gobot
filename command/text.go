package command

// Text basic command wrapper
type Text struct {
	Name string `json:"name"`
	Mod  bool   `json:"mod"`
	Sub  bool   `json:"sub"`
	Text string `json:"text"`
}

// Execute run command
func (b Text) Execute(resp Response, event MessageEvent) {
	if b.Mod && !event.IsModerator {
		return
	}

	if b.Sub && !event.IsSubscriber {
		return
	}

	resp(event.Format(b.Text))
}
