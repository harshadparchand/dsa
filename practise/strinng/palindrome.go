package strinng

func IsPalindrome(input string) bool {
	start, end := 0, len(input)-1
	for start <= end {
		if input[start] == input[end] {
			start++
			end--
		} else {
			return false
		}
	}
	return true
}

func IsPalindromeRecur(s string, start, end int) bool {
	if start > end {
		return true
	}
	if s[start] == s[end] {
		start++
		end--
		return IsPalindromeRecur(s, start, end)
	} else {
		return false
	}
}
