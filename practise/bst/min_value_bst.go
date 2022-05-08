package bst

import "github.com/shellagilehub/practise/structs"

func MinValueInBSTIterative(root *structs.Node) *structs.Node {
	if root == nil {
		return nil
	}

	for root.Left != nil {
		root = root.Left
	}
	return root
}

func MinValueInBSTRecursive(root *structs.Node) *structs.Node {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root
	}
	return MinValueInBSTRecursive(root.Left)
}
