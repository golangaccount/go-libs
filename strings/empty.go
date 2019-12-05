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

func RemoveEmpty(src []string) []string {
	result := make([]string, 0)
	for _, item := range src {
		if !IsEmptyOrWhite(item) {
			result = append(result, item)
		}
	}
	return result
}
