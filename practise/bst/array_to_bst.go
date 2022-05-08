package bst

import "github.com/shellagilehub/practise/structs"

var ArrayToBSTIndex = 0

// ArrayToBST : Reset Index to 0
func ArrayToBST(root *structs.Node, array []int) {
	if root == nil {
		return
	}
	ArrayToBST(root.Left, array)
	root.Data = array[ArrayToBSTIndex]
	ArrayToBSTIndex++
	ArrayToBST(root.Right, array)
}
