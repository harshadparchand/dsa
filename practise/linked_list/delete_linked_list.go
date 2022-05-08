package linked_list

import "github.com/shellagilehub/practise/structs"

func DeleteLinkedList(head *structs.LinkedListNode) {
	if head == nil {
		return
	}
	DeleteLinkedList(head.Next)
	head = nil
	return
}
