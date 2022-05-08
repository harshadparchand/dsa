package queue

import "github.com/shellagilehub/practise/structs"

var enqueueStack = structs.StackInt{}
var dequeueStack = structs.StackInt{}

func Enqueue(data int) {
	for !dequeueStack.IsEmpty() {
		val := dequeueStack.Peek()
		dequeueStack = dequeueStack.Pop()
		enqueueStack = enqueueStack.Push(val)
	}
	enqueueStack = enqueueStack.Push(data)
}

func Dequeue() int {
	for !enqueueStack.IsEmpty() {
		data := enqueueStack.Peek()
		enqueueStack = enqueueStack.Pop()
		dequeueStack = dequeueStack.Push(data)
	}
	if !dequeueStack.IsEmpty() {
		data := dequeueStack.Peek()
		dequeueStack = dequeueStack.Pop()
		return data
	}
	return -1
}
