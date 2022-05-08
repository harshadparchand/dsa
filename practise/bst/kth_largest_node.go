package bst

import (
	"fmt"
	"github.com/shellagilehub/practise/structs"
)

var KthLargestIndex = 0

func KthLargestNode(root *structs.Node, k int) {
	if root == nil {
		return
	}

	KthLargestNode(root.Right, k)
	KthLargestIndex++
	if k == KthLargestIndex {
		fmt.Println(root.Data)
		return
	}
	KthLargestNode(root.Left, k)
	return
}
