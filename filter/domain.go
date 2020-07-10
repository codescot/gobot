package filter

import (
	"fmt"
	"regexp"
)

// Domain domain
type Domain struct{}

// Apply filter logic
func (Domain) Apply(message string) int {
	matched, err := regexp.MatchString(`([a-zA-Z0-9]+[.]+[a-zA-Z]+)`, message)
	if err != nil {
		fmt.Println(err)
		return Ignore
	}

	if matched {
		return Delete
	}

	return Ignore
}
