package strings

import (
	"strings"
)

var defaultString string

func IsEmpty(s string) bool {
	if s == defaultString {
		return true
	} else {
		return false
	}
}

func IsWhite(s string) bool {
	if len(s) != 0 && len(strings.TrimSpace(s)) == 0 {
		return true
	} else {
		return false
	}
}

func IsEmptyOrWhite(s string) bool {
	return IsEmpty(s) || IsWhite(s)
}
