package traversal

import (
	"fmt"
	"github.com/shellagilehub/practise/structs"
)

func InOrderRecursiveTraversal(root *structs.Node) {
	if root == nil {
		return
	}
	InOrderRecursiveTraversal(root.Left)
	fmt.Println(root.Data)
	InOrderRecursiveTraversal(root.Right)
}

func InOrderIterative(root *structs.Node) {
	if root == nil {
		return
	}
	temp := root
	stack := structs.StackNode{}
	for temp != nil || !stack.IsEmpty() {
		if temp != nil {
			stack = stack.Push(temp)
			temp = temp.Left
		} else {
			temp = stack.Peek()
			stack = stack.Pop()
			fmt.Println(temp.Data)
			temp = temp.Right
		}
	}
}
