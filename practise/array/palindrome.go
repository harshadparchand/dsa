package array

import (
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	s = strings.ToUpper(s)
	l := 0
	r := len(s) - 1
	for l < r {
		if !unicode.IsLetter(rune(s[l])) && !unicode.IsNumber(rune(s[l])) {
			l++
		} else if !unicode.IsLetter(rune(s[r])) && !unicode.IsNumber(rune(s[r])) {
			r--
		} else if s[l] == s[r] {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}
