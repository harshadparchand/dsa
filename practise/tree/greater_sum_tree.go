package tree

import "github.com/shellagilehub/practise/structs"

var sum = 0

func GreaterSumTree(root *structs.Node) {
	if root == nil {
		return
	}

	GreaterSumTree(root.Right)
	sum += root.Data
	root.Data = sum - root.Data
	GreaterSumTree(root.Left)
}
