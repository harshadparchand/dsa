package linked_list

import "github.com/shellagilehub/practise/structs"

func LengthOfLinkedListIterative(head *structs.LinkedListNode) int {
	if head == nil {
		return 0
	}

	length := 0
	for head != nil {
		head = head.Next
		length++
	}
	return length
}

func LengthOfLinkedListRecursive(head *structs.LinkedListNode) int {
	if head == nil {
		return 0
	}

	return 1 + LengthOfLinkedListIterative(head.Next)
}
