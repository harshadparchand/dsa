package bst

import "github.com/shellagilehub/practise/structs"

var BSTToArrayIndex = 0

// BSTToArray : Reset Index to 0
func BSTToArray(root *structs.Node, array []int) []int {
	if root == nil {
		return array
	}
	BSTToArray(root.Left, array)
	array[BSTToArrayIndex] = root.Data
	BSTToArrayIndex++
	BSTToArray(root.Right, array)
	return array
}
