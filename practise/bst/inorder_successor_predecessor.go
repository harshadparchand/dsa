package bst

import (
	"github.com/shellagilehub/practise/structs"
)

var InorderSuc *structs.Node
var InorderPre *structs.Node

func InOrderSuccessorPredecessorRecursive(root *structs.Node, n *structs.Node) {
	if root == nil {
		return
	}

	InOrderSuccessorPredecessorRecursive(root.Left, n)
	//Get Predecessor
	if root.Data < n.Data {
		InorderPre = root
	} else if root.Data > n.Data { //Get Successor
		if InorderSuc == nil || (InorderSuc != nil && InorderSuc.Data > root.Data) { // Because of traversal way which take left subtree first and then the root node
			InorderSuc = root
		}
	}
	InOrderSuccessorPredecessorRecursive(root.Right, n)

}

func InOrderSuccessorPredecessorIterative(root *structs.Node, n *structs.Node) {
	if root == nil {
		return
	}

	stack := structs.StackNode{}
	for root != nil || !stack.IsEmpty() {
		if root != nil {
			stack = stack.Push(root)
			root = root.Left
		} else {
			root = stack.Peek()
			stack = stack.Pop()
			if root.Data < n.Data {
				InorderPre = root
			} else if root.Data > n.Data { //Get Successor
				if InorderSuc == nil || (InorderSuc != nil && InorderSuc.Data > root.Data) { // Because of traversal way which take left subtree first and then the root node
					InorderSuc = root
				}
			}
			root = root.Right
		}
	}
}
