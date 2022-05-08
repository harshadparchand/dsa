package stack

import (
	"fmt"
	"github.com/shellagilehub/practise/structs"
)

func ZigZagBSTTraversal(root *structs.Node) {
	if root == nil {
		return
	}

	s1 := structs.StackNode{}
	s2 := structs.StackNode{}
	s1 = s1.Push(root)

	for !s1.IsEmpty() || !s2.IsEmpty() {
		for !s1.IsEmpty() {
			node := s1.Peek()
			fmt.Println(node.Data)
			s1 = s1.Pop()
			if node.Left != nil {
				s2 = s2.Push(node.Left)
			}
			if node.Right != nil {
				s2 = s2.Push(node.Right)
			}
		}
		for !s2.IsEmpty() {
			node := s2.Peek()
			fmt.Println(node.Data)
			s2 = s2.Pop()
			if node.Right != nil {
				s1 = s1.Push(node.Right)
			}
			if node.Left != nil {
				s1 = s1.Push(node.Left)
			}
		}
	}
}
