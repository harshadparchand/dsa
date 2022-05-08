package bst

import "github.com/shellagilehub/practise/structs"

func removeOutOfRange(root *structs.Node, min, max int) *structs.Node {
	if root == nil {
		return nil
	}

	root.Left = removeOutOfRange(root.Left, min, max)
	root.Right = removeOutOfRange(root.Right, min, max)

	if root.Data < min {
		return root.Right
	}

	if root.Data > max {
		return root.Left
	}

	return root
}
