package stack

import "github.com/shellagilehub/practise/structs"

/*
https://www.geeksforgeeks.org/check-if-stack-elements-are-pairwise-consecutive/
stack = structs.StackInt{}
	stack = stack.Push(1)
	stack = stack.Push(2)
	stack = stack.Push(8)
	stack = stack.Push(6)
*/

func PairwiseConsecutive(stack structs.StackInt) bool {
	if stack.IsEmpty() {
		return true
	}

	for !stack.IsEmpty() {
		first := stack.Peek()
		stack = stack.Pop()
		if stack.IsEmpty() {
			return false
		}
		second := stack.Peek()
		stack = stack.Pop()
		if first-second == 1 && first-second == -1 {
			return false
		}
	}
	return true
}
