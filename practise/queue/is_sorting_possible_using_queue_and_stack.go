package queue

import "github.com/shellagilehub/practise/structs"

//https://www.geeksforgeeks.org/check-queue-can-sorted-another-queue-using-stack/

func CheckSorted(arr []int) bool {
	if len(arr) == 0 {
		return false
	}
	stack := structs.StackInt{}

	expectedNumber := 1
	for i := 0; i < len(arr); i++ {
		if arr[i] == expectedNumber {
			expectedNumber++
			continue
		} else {
			if stack.IsEmpty() {
				stack = stack.Push(arr[i])
			} else {
				val := stack.Peek()
				if val < arr[i] {
					return false
				} else {
					stack = stack.Pop()
					stack = stack.Push(arr[i])
				}
			}
		}
	}
	return true
}
