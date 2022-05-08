package bst

import "github.com/shellagilehub/practise/structs"

func GetCount(root *structs.Node, min, max int) int {
	if root == nil {
		return 0
	}
	if root.Data >= min && root.Data <= max {
		return 1 + GetCount(root.Left, min, max) + GetCount(root.Right, min, max)
	} else if root.Data < min {
		return GetCount(root.Right, min, max)
	} else {
		return GetCount(root.Left, min, max)
	}
}
