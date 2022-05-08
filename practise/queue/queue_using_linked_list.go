package queue

import "github.com/shellagilehub/practise/structs"

var front *structs.LinkedListNode
var rear *structs.LinkedListNode

//O(1)

func EnqueueQueue(data int) {
	node := &structs.LinkedListNode{
		Data: data,
		Next: nil,
	}
	//If queue is empty
	if front == nil && rear == nil {
		front = node
		rear = node
		return
	}
	// Add the new node at the end of queue and change rear
	rear.Next = node
	rear = rear.Next
	return
}

//O(1)

func DequeueQueue() int {
	if rear == nil {
		return -1
	}
	temp := front
	front = front.Next
	if front == nil {
		rear = nil
	}
	return temp.Data
}
