package linked_list

import "github.com/shellagilehub/practise/structs"

func DeleteNode(head *structs.LinkedListNode, key int) *structs.LinkedListNode {
	if head == nil {
		return head
	}

	//if the key is in head
	if head.Data == key {
		head = head.Next
		return head
	}

	var prev *structs.LinkedListNode
	current := head
	//iterate until the key is found
	for current != nil {
		if current.Data != key {
			prev = current
			current = current.Next
		} else {
			break
		}
	}

	//if key is not found
	if current == nil {
		return head
	}

	//delete the key
	prev.Next = current.Next
	return head
}
