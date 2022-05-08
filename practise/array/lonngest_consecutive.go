package array

//fmt.Println(LongestConsecutive([]int{100, 4, 200, 1, 2 , 3}))
//https://leetcode.com/problems/longest-consecutive-sequence/

func LongestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, value := range nums {
		set[value] = true
	}
	longest := 0
	for i := 0; i < len(nums); i++ {
		current := nums[i]
		partial := 1
		//Checking if it is start of sequence
		_, isPresent := set[current-1]
		if !isPresent {
			for true {
				if set[current+1] {
					partial++
					current++
					continue
				}
				break
			}
			if longest < partial {
				longest = partial
			}
		}
	}
	return longest
}
