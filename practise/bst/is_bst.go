package bst

import "github.com/shellagilehub/practise/structs"

func IsBST(root *structs.Node, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Data < min || root.Data > max {
		return false
	}

	return IsBST(root.Left, min, root.Data-1) && IsBST(root.Right, root.Data+1, max)
}
