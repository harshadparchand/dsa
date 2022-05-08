package main

import (
	"fmt"
	"github.com/shellagilehub/practise/array"
	"github.com/shellagilehub/practise/bst"
	"github.com/shellagilehub/practise/heap"
	"github.com/shellagilehub/practise/linked_list"
	"github.com/shellagilehub/practise/queue"
	"github.com/shellagilehub/practise/recursion"
	stack2 "github.com/shellagilehub/practise/stack"
	"github.com/shellagilehub/practise/structs"
	"github.com/shellagilehub/practise/traversal"
	"github.com/shellagilehub/practise/tree"
)

func main() {
	// Taking input from user
	//var input int
	//fmt.Scanln(&input)
	//
	//scanner := bufio.NewScanner(os.Stdin)
	//if scanner.Scan() {
	//	opinions := scanner.Text()
	//	fmt.Println(opinions)
	//}

	head := bst.InsertRecursive(nil, 10)
	head = bst.InsertRecursive(head, 20)
	head = bst.InsertRecursive(head, 30)
	head = bst.InsertRecursive(head, 40)
	head = bst.InsertRecursive(head, 70)
	head = bst.InsertRecursive(head, 60)
	head = bst.InsertRecursive(head, 80)
	head = bst.InsertRecursive(head, 80)
	head = bst.InsertRecursive(head, 55)
	traversal.InOrderRecursiveTraversal(head)
	fmt.Println("-------------")
	fmt.Println(array.IsValidSudoku([][]int{
		{5, 1, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}))
	fmt.Println(array.DecodeString("2[a2[ac]]2[bc]"))
	fmt.Println(array.LongestConsecutive([]int{100, 4, 200, 1, 2, 3}))
	fmt.Println(array.MaxDistinctNum([]int{1, 2, 2, 2}, 1))
	heap.PrintNodesGreaterThanK(heap.CreateHeap(), 80)
	fmt.Println()

	queue.PushQueue(1)
	queue.PushQueue(2)
	queue.PushQueue(3)
	queue.PushQueue(4)
	fmt.Println(queue.PopQueue())
	fmt.Println(queue.PopQueue())
	//queue.PushQueue(5)
	fmt.Println(queue.PopQueue())
	fmt.Println(queue.PopQueue())
	fmt.Println(queue.PopQueue())
	fmt.Println(queue.PopQueue())
	fmt.Println(queue.PopQueue())
	//queue.PushQueue(6)
	fmt.Println(queue.PopQueue())
	fmt.Println(queue.PopQueue())
	fmt.Println(queue.CheckSorted([]int{7, 1, 6, 2, 4, 3, 5}))

	fmt.Println(tree.TreeHeightIterative(tree.CreateTree()))
	queue.MaximumLevelSumBST(tree.CreateTree())
	merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3)
	stack := structs.StackInt{}
	stack = stack.Push(4)
	stack = stack.Push(3)
	stack = stack.Push(2)
	stack = stack.Push(1)
	fmt.Println(stack)
	fmt.Println(ReverseStack(stack))
	//printNumber(8)
	fmt.Println(kthGrammarS(7, 15))
	recursion.GenerateStringsWithUpperCase("a1b2c", "")
	recursion.GenerateBalancedParenthesis("(((", ")))", "")
	fmt.Println(stack2.ReverseNumber(12345))
	stack = structs.StackInt{}
	stack = stack.Push(-3)
	stack = stack.Push(-2)
	stack = stack.Push(9)
	stack = stack.Push(8)

	fmt.Println(stack2.PairwiseConsecutive(stack))
	fmt.Println(stack2.SimplifyEquation("a-((b-c)-d)"))
	linkedList := &structs.LinkedListNode{
		Data: 1,
		Next: &structs.LinkedListNode{
			Data: 2,
			Next: &structs.LinkedListNode{
				Data: 3,
				Next: nil,
			},
		},
	}
	node := linked_list.NthNodeInLinkedListRecursion(linkedList, 3)

	fmt.Println(node.Data)
}

func Gen(NBit, ones, zeros int, op string) {
	if NBit <= 0 {
		fmt.Println(op)
		return
	}

	op1 := op + "1"
	ones++
	NBit--
	Gen(NBit, ones, zeros, op1)
	if ones > zeros+1 {
		op2 := op + "0"
		zeros++
		NBit--
		Gen(NBit, ones, zeros, op2)
	}
	return
}

func generateStrings(input, op string) {
	if input == "" {
		fmt.Println(op)
		return
	}
	opInc := op + string(input[0])
	input = input[1:len(input)]
	generateStrings(input, opInc)
	generateStrings(input, op)
}

func balancedStringSplit(s string) int {
	count := 0
	LCount := 0
	RCount := 0
	for _, c := range s {
		if string(c) == "L" {
			LCount++
		} else {
			RCount++
		}
		if LCount == RCount {
			count++
			LCount, RCount = 0, 0
		}
	}
	return count
}

func ReverseStack(stack structs.StackInt) structs.StackInt {
	if stack.IsEmpty() {
		return stack
	}
	top := stack.Peek()
	stack = stack.Pop()
	ReverseStack(stack)
	InsertAtTheBottom(stack, top)
	return stack
}

func InsertAtTheBottom(stack structs.StackInt, x int) structs.StackInt {
	if stack.IsEmpty() {
		stack = stack.Push(x)
		return stack
	}
	val := stack.Peek()
	stack = stack.Pop()
	InsertAtTheBottom(stack, x)
	stack = stack.Push(val)
	return stack
}

//Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//Output: [1,2,2,3,5,6]

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	i, j := m-1, n-1
	index := m + n - 1
	for i >= 0 && j >= 0 && index >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[index] = nums1[i]
			nums1[i] = 0
			i--
		} else {
			nums1[index] = nums2[j]
			j--
		}
		index--
	}
	for j >= 0 {
		nums1[index] = nums2[j]
		index--
		j--
	}
	return
}

func kthGrammar(n int, k int) int {
	if k > pow(2, n-1) {
		return -1
	}
	if n == 1 && k == 1 {
		return 0
	}

	mid := pow(2, n-2)
	if k <= mid {
		return kthGrammar(n-1, k)
	} else {
		if kthGrammar(n-1, k-mid) == 0 {
			return 1
		} else {
			return 0
		}
	}
}
func kthGrammarS(n int, k int) int {
	if n == 1 {
		return 0
	}
	parent := kthGrammarS(n-1, k/2+k%2)
	kIsOdd := k % 2
	if parent == 1 {
		if kIsOdd == 1 {
			return 1
		} else {
			return 0
		}
	} else {
		if kIsOdd == 1 {
			return 0
		} else {
			return 1
		}
	}
}

func pow(base, power int) int {
	if power == 0 {
		return 1
	}

	return base * pow(base, power-1)
}
