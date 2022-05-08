package bst

import "github.com/shellagilehub/practise/structs"

func CeilForKey(root *structs.Node, key int) int {
	if root == nil {
		return -1
	}
	if root.Data == key {
		return key
	}
	//if root.Data is greater than key then the ceil is in left subtree
	if root.Data < key {
		return CeilForKey(root.Right, key)
	}
	// else either the left subtree or root has the ceil value
	ceil := CeilForKey(root.Left, key)

	//if ceil > key then return root.Data
	if ceil >= key {
		return ceil
	}
	return root.Data
}
