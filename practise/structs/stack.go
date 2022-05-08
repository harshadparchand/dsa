package structs

//Node

type StackNode []*Node

func (s StackNode) Push(v *Node) StackNode {
	return append(s, v)
}

func (s StackNode) Pop() StackNode {
	if len(s) == 0 {
		return nil
	}
	return s[:len(s)-1]
}

func (s StackNode) IsEmpty() bool {
	if len(s) == 0 {
		return true
	}
	return false
}

func (s StackNode) Peek() *Node {
	return s[len(s)-1]
}

func (s StackNode) Size() int {
	return len(s)
}

//int

type StackChar []rune

func (s StackChar) Push(v rune) StackChar {
	return append(s, v)
}

func (s StackChar) Pop() StackChar {
	if len(s) == 0 {
		return nil
	}
	return s[:len(s)-1]
}

func (s StackChar) IsEmpty() bool {
	if len(s) == 0 {
		return true
	}
	return false
}

func (s StackChar) Peek() rune {
	return s[len(s)-1]
}

func (s StackChar) Size() int {
	return len(s)
}

//int

type StackInt []int

func (s StackInt) Push(v int) StackInt {
	return append(s, v)
}

func (s StackInt) Pop() StackInt {
	if len(s) == 0 {
		return nil
	}
	return s[:len(s)-1]
}

func (s StackInt) IsEmpty() bool {
	if len(s) == 0 {
		return true
	}
	return false
}

func (s StackInt) Peek() int {
	return s[len(s)-1]
}

func (s StackInt) Size() int {
	return len(s)
}

//String

type StackString []string

func (s StackString) Push(v string) StackString {
	return append(s, v)
}

func (s StackString) Pop() StackString {
	if len(s) == 0 {
		return nil
	}
	return s[:len(s)-1]
}

func (s StackString) IsEmpty() bool {
	if len(s) == 0 {
		return true
	}
	return false
}

func (s StackString) Peek() string {
	return s[len(s)-1]
}

func (s StackString) Size() int {
	return len(s)
}
