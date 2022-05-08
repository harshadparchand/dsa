package structs

//Node

type QueueNode []*Node

func (q QueueNode) Enqueue(element *Node) QueueNode {
	return append(q, element)
}

func (q QueueNode) Dequeue() QueueNode {
	return q[1:]
}

func (q QueueNode) Peek() *Node {
	return q[0]
}

func (q QueueNode) IsEmpty() bool {
	if len(q) > 0 {
		return false
	}
	return true
}

func (q QueueNode) Size() int {
	return len(q)
}

//Integer

type QueueInt []int

func (q QueueInt) Enqueue(element int) QueueInt {
	return append(q, element)
}

func (q QueueInt) Dequeue() QueueInt {
	return q[1:]
}

func (q QueueInt) Peek() int {
	return q[0]
}

func (q QueueInt) IsEmpty() bool {
	if len(q) == 0 {
		return true
	}
	return false
}

func (q QueueInt) Size() int {
	return len(q)
}

//String

type QueueString []string

func (q QueueString) Enqueue(element string) QueueString {
	return append(q, element)
}

func (q QueueString) Dequeue() QueueString {
	return q[1:]
}

func (q QueueString) Peek() string {
	return q[0]
}

func (q QueueString) IsEmpty() bool {
	if len(q) == 0 {
		return true
	}
	return false
}

func (q QueueString) Size() int {
	return len(q)
}
