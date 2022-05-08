package bst

import "github.com/shellagilehub/practise/structs"

//Dead End means, we are not able to insert any integer element after that node.
func DeadEndBST(root *structs.Node, min, max int) bool {
	if root == nil {
		return false
	}
	//If min == max, then the range has only one value and which is already occupied by a node
	if min == max {
		return true
	}
	return DeadEndBST(root.Left, min, root.Data-1) || DeadEndBST(root.Right, root.Data+1, max)
}
