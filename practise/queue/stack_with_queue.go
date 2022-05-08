package queue

import "github.com/shellagilehub/practise/structs"

var queue = structs.QueueInt{}

func PushQueue(data int) {
	size := queue.Size()
	queue = queue.Enqueue(data)
	for size != 0 {
		val := queue.Peek()
		queue = queue.Dequeue()
		queue = queue.Enqueue(val)
		size--
	}
}

func PopQueue() int {
	if queue.Size() == 0 {
		return -1
	}
	val := queue.Peek()
	queue = queue.Dequeue()
	return val
}
