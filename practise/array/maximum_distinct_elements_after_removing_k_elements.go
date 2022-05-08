package array

//https://www.geeksforgeeks.org/maximum-distinct-elements-removing-k-elements/

func MaxDistinctNum(arr []int, k int) int {
	set := make(map[int]bool)

	for _, value := range arr {
		_, isPresent := set[value]
		if isPresent && k > 0 {
			k--
		} else {
			set[value] = true
		}
	}
	if k > 0 {
		return len(set) - k
	}
	return len(set)
}
