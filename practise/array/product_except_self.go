package array

//https://leetcode.com/problems/product-of-array-except-self/
//Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
//The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
//You must write an algorithm that runs in O(n) time and without using the division operation.

func ProductExceptSelf(nums []int) []int {
	prefix, postfix := 1, 1
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = prefix
		prefix *= nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = postfix * res[i]
		postfix *= nums[i]
	}
	return res
}
