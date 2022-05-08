package bst

import "github.com/shellagilehub/practise/structs"

func SortedArrayToBalancedBST(array []int, start, end int) *structs.Node {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	node := structs.Node{Data: array[mid]}
	//Form the left subtree
	node.Left = SortedArrayToBalancedBST(array, start, mid-1)
	//Form the right subtree
	node.Right = SortedArrayToBalancedBST(array, mid+1, end)
	return &node
}
