package queue

import "github.com/shellagilehub/practise/structs"

func ReverserQueueUsingStack(queue structs.QueueNode) structs.QueueNode {
	if queue == nil {
		return queue
	}
	stack := structs.StackNode{}
	for !queue.IsEmpty() {
		node := queue.Peek()
		queue = queue.Dequeue()
		stack = stack.Push(node)
	}

	for !stack.IsEmpty() {
		node := stack.Peek()
		queue = queue.Enqueue(node)
		stack = stack.Pop()
	}
	return queue
}
