package bst

import "github.com/shellagilehub/practise/structs"

func InOrderSuccessorIterative(root *structs.Node, n *structs.Node) *structs.Node {
	inOrderSuccessor := root
	if n.Right != nil {
		for inOrderSuccessor != nil {
			inOrderSuccessor = inOrderSuccessor.Left
		}
		return inOrderSuccessor
	}

	for root != nil {
		if root.Data > n.Data {
			inOrderSuccessor = root
			root = root.Left
		} else if root.Data < n.Data {
			root = root.Right
		} else {
			break
		}
	}
	return inOrderSuccessor
}

var InOrderSuccessor *structs.Node

func InOrderSuccessorRecursive(root *structs.Node, n *structs.Node) {
	if root == nil {
		return
	}

	InOrderSuccessorIterative(root.Left, n)
	if root.Data > n.Data && InOrderSuccessor == nil {
		InOrderSuccessor = root
		return
	}
	InOrderSuccessorIterative(root.Right, n)
}
