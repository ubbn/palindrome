package palindrome

import (
	"errors"
	"strings"
)

// Truth of whether given string is palindrome
// i.e. a string that is read same from its back and front
func IsPalindrome(str string) (bool, error) {
	trimmed := strings.Trim(str, " ")
	if trimmed == "" {
		return false, errors.New("empty input")
	}

	length := len(trimmed)
	for i := 0; i < length/2; i++ {
		if trimmed[i] != trimmed[len(trimmed)-i-1] {
			return false, nil
		}
	}
	return true, nil
}
