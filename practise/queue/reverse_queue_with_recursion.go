package queue

import "github.com/shellagilehub/practise/structs"

//https://www.geeksforgeeks.org/reversing-queue-using-recursion/

func ReverseQueueUsingRecursion(queue structs.QueueNode) structs.QueueNode {
	if queue.IsEmpty() {
		return queue
	}
	node := queue.Peek()
	queue = queue.Dequeue()
	queue = ReverserQueueUsingStack(queue)
	queue = queue.Enqueue(node)
	return queue
}
