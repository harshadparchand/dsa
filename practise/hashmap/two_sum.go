package hashmap

func TwoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for currentIndex, value := range nums {
		requiredIndex, isPresent := hashMap[target-value]
		if isPresent {
			return []int{requiredIndex, currentIndex}
		}
		hashMap[value] = currentIndex
	}
	return []int{}
}
