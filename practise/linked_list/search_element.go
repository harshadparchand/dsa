package linked_list

import "github.com/shellagilehub/practise/structs"

func SearchElementIterative(head *structs.LinkedListNode, key int) bool {
	if head == nil {
		return false
	}

	for head != nil {
		if head.Data == key {
			return true
		}
		head = head.Next
	}
	return false
}

func SearchElementRecursive(head *structs.LinkedListNode, key int) bool {
	if head == nil {
		return false
	}
	if head.Data == key {
		return true
	}
	return SearchElementIterative(head.Next, key)
}
