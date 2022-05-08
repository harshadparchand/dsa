package bst

import "github.com/shellagilehub/practise/structs"

func BSTSearchIterative(root *structs.Node, key int) bool {

	node := root
	for node != nil {
		if node.Data < key {
			node = node.Right
		} else if node.Data > key {
			node = node.Left
		} else {
			return true
		}
	}
	return false
}

func BSTSearchRecursive(root *structs.Node, key int) bool {

	if root == nil {
		return false
	}

	if root.Data > key {
		return BSTSearchRecursive(root.Left, key)
	} else if root.Data < key {
		return BSTSearchRecursive(root.Right, key)
	} else {
		return true
	}
}
