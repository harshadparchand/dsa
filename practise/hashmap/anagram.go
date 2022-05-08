package hashmap

import "strings"

func IsAnagram(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}
	sChars := strings.Split(s, "")
	tChars := strings.Split(t, "")
	hashmap := make(map[string]int)
	for _, value := range sChars {
		hashmap[value] = hashmap[value] + 1
	}
	for _, value := range tChars {
		hashmap[value] = hashmap[value] - 1
	}
	for _, value := range hashmap {
		if value != 0 {
			return false
		}
	}
	return true
}
