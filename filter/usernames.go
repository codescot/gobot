package filter

import (
	"fmt"
	"strings"
	"time"
)

// Usernames username filter
type Usernames struct {
	Blocked  []string
	Username string
}

// Apply filter logic
func (u Usernames) Apply(message string) int {
	if deepContains(u.Blocked, strings.ToLower(u.Username)) {
		fmt.Printf("[%v] bot enumeration attack %s\n", time.Now().Format(time.Stamp), u.Username)
		return Ban
	}

	return Ignore
}

func deepContains(haystack []string, needle string) bool {
	for _, item := range haystack {
		if strings.Contains(needle, item) {
			return true
		}
	}

	return false
}
