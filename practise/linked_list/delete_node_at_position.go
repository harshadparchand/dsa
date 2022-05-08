package linked_list

import "github.com/shellagilehub/practise/structs"

func DeleteNodeAtPosition(head *structs.LinkedListNode, position int) *structs.LinkedListNode {
	if head == nil {
		return head
	}

	//if the key is in head
	if position == 0 {
		head = head.Next
		return head
	}

	current := head
	//iterate until the key is found
	counter := 0
	for current != nil && counter != position-1 {
		current = current.Next
	}

	//if key is not found
	if current == nil || current.Next == nil {
		return head
	}

	//delete the key
	current.Next = current.Next.Next
	return head
}
