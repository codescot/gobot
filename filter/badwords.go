package filter

import (
	"strings"
)

// BadWords bad words filter
type BadWords struct {
	BadWords []string
}

// Apply filter logic
func (bw BadWords) Apply(message string) int {
	messageWords := strings.Split(" ", message)
	for _, word := range bw.BadWords {
		if itContains(messageWords, word) {
			return Delete
		}
	}

	return Ignore
}

func itContains(haystack []string, needle string) bool {
	for _, item := range haystack {
		if strings.Contains(strings.ToLower(needle), strings.ToLower(item)) {
			return true
		}
	}

	return false
}
