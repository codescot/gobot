package filter

import (
	"fmt"
	"regexp"

	"github.com/codescot/gobot/command"
)

// Domain domain
type Domain struct {
	Perm string
}

func (d Domain) ShouldApply(sub, mod bool) bool {
	return !command.HasPerm(d.Perm, sub, mod)
}

// Apply filter logic
func (Domain) Apply(message string) int {

	matched, err := regexp.MatchString(`([a-zA-Z0-9]+[.][a-zA-Z]+)`, message)
	if err != nil {
		fmt.Println(err)
		return Ignore
	}

	if matched {
		return Delete
	}

	return Ignore
}
