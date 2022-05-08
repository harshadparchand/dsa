package hashmap

//Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
//https://leetcode.com/problems/top-k-frequent-elements/

func TopKFrequent(nums []int, k int) []int {
	hashmapValueToFrequency := make(map[int]int)
	hashmapFrequencyToValue := make(map[int][]int)
	//Storing nums -> frequencies
	for _, value := range nums {
		hashmapValueToFrequency[value] = hashmapValueToFrequency[value] + 1
	}
	//Storing frequencies -> nums
	for value, frequency := range hashmapValueToFrequency {
		hashmapFrequencyToValue[frequency] = append(hashmapFrequencyToValue[frequency], value)
	}
	var res []int
	for i := len(nums); i > 0; i-- {
		values := hashmapFrequencyToValue[i]
		for _, value := range values {
			if k == 0 {
				break
			}
			res = append(res, value)
			k--
		}
	}
	return res
}
