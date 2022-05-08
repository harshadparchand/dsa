package hashmap

func ContainsDuplicates(nums []int) bool {
	hashmap := make(map[int]bool)
	for _, key := range nums {
		if hashmap[key] == true {
			return true
		}
		hashmap[key] = true
	}
	return false
}
