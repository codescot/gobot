package filter

import (
	"strings"
)

// BadWords bad words filter
type BadWords struct {
	BadWords []string
}

func (BadWords) ShouldApply(sub, mod bool) bool {
	return true
}

// Apply filter logic
func (bw BadWords) Apply(message string) int {
	for _, word := range bw.BadWords {
		if strings.Contains(message, word) {
			return Delete
		}
	}

	return Ignore
}
