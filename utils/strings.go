package utils

import "strings"

// IsStringEmpty See if a String if empty, like ""
func IsStringEmpty(s string) bool {
	if strings.Compare(s, "") == 0 {
		return true
	}

	return false
}
