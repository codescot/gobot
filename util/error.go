package util

import "fmt"

// IsError generic error handling
func IsError(appError error) bool {
	if appError == nil {
		return false
	}

	fmt.Println(appError.Error())
	return true
}
