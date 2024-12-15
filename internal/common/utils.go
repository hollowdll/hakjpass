package common

import (
	"strings"
)

func ContainsCharFromCharset(input string, charset string) bool {
	for _, char := range input {
		if strings.ContainsRune(charset, char) {
			return true
		}
	}
	return false
}
