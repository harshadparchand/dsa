package linked_list

import (
	"github.com/shellagilehub/practise/structs"
)

func NthNodeInLinkedListIteration(head *structs.LinkedListNode, position int) *structs.LinkedListNode {
	if head == nil {
		return head
	}

	//if the key is in head
	if position == 0 {
		return head
	}

	current := head
	//iterate until the key is found
	counter := 0
	for current != nil && counter != position {
		current = current.Next
		counter++
	}

	return current
}

func NthNodeInLinkedListRecursion(head *structs.LinkedListNode, position int) *structs.LinkedListNode {
	if head == nil {
		return head
	}

	if position == 0 {
		return head
	}
	position--
	return NthNodeInLinkedListRecursion(head.Next, position)
}
