package problems

//You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.
//Return true if you can reach the last index, or false otherwise.

func canJump(nums []int) bool {
	return Jumps(nums, 0)
}

func Jumps(nums []int, index int) bool {
	if index == len(nums) {
		return true
	}
	for i := 0; i < nums[index]; i++ {
		Jumps(nums, index)
	}
	return false
}
