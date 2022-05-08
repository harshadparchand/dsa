package bst

import "github.com/shellagilehub/practise/structs"

func FindNearestValueLessThanEqualToN(root *structs.Node, n int) int {
	if root == nil {
		return -1
	}
	if root.Data == n {
		return n
	} else if root.Data < n { //If root.Data < n, traverse to right subtree
		k := FindNearestValueLessThanEqualToN(root.Right, n)
		if k == -1 {
			return root.Data
		} else {
			return k
		}
	} else if root.Data > n {
		FindNearestValueLessThanEqualToN(root.Left, n)
	}
	return -1
}
