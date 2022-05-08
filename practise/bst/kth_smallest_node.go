package bst

import (
	"fmt"
	"github.com/shellagilehub/practise/structs"
)

var KthSmallestIndex = 0

func KthSmallestNode(root *structs.Node, k int) {
	if root == nil {
		return
	}

	KthSmallestNode(root.Left, k)
	KthLargestIndex++
	if k == KthLargestIndex {
		fmt.Println(root.Data)
		return
	}
	KthSmallestNode(root.Right, k)
	return
}
