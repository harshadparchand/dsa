package hashmap

import (
	"sort"
)

func GroupAnagrams(strs []string) [][]string {
	hashmap := make(map[string][]string)
	//iterating over strs
	for _, value := range strs {
		//creating character array from str
		runeSlice := []rune(value)
		//sorting character array
		sort.Slice(runeSlice, func(i, j int) bool {
			return runeSlice[i] < runeSlice[j]
		})
		//checking if the sorted character array is present or not
		if _, isPresent := hashmap[string(runeSlice)]; !isPresent {
			//if character array is not present then insert it as is
			hashmap[string(runeSlice)] = []string{value}
		} else {
			//if character array is present then append the valuesand then stored it back
			hashmap[string(runeSlice)] = append(hashmap[string(runeSlice)], value)
		}
	}
	var res [][]string
	for _, value := range hashmap {
		res = append(res, value)
	}
	return res
}
