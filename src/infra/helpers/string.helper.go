package helpers

import "strings"

func IsEmpty(data string) bool {
	if len(strings.TrimSpace(data)) > 0 {
		return false
	}
	return true
}
