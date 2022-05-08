package bst

import "github.com/shellagilehub/practise/structs"

func DeleteLeafNodes(root *structs.Node) *structs.Node {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return nil
	}
	root.Left = DeleteLeafNodes(root.Left)
	root.Right = DeleteLeafNodes(root.Right)
	return root
}
