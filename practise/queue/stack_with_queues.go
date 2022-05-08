package queue

import "github.com/shellagilehub/practise/structs"

var q1 = structs.QueueInt{}
var q2 = structs.QueueInt{}

//O(n)

func Push(data int) {
	q2 = q2.Enqueue(data)
	for !q1.IsEmpty() {
		val := q1.Peek()
		q1 = q1.Dequeue()
		q2 = q2.Enqueue(val)
	}
	temp := q2
	q2 = q1
	q1 = temp
}

//O(1)

func Pop() int {

	if q1.IsEmpty() {
		return -1
	}
	data := q1.Peek()
	q1 = q1.Dequeue()
	return data
}
