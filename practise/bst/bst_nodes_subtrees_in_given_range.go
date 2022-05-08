package bst

import (
	"github.com/shellagilehub/practise/structs"
)

var BSTSubtreesInGivenRangeCount = 0

func BSTSubtreesInGivenRange(root *structs.Node, min, max int) bool {
	if root == nil {
		return true
	}

	left := BSTSubtreesInGivenRange(root.Left, min, max)
	right := BSTSubtreesInGivenRange(root.Right, min, max)

	if left && right && root.Data >= min && root.Data <= max {
		BSTSubtreesInGivenRangeCount++
		return true
	}

	return false
}
