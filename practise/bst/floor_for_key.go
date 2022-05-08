package bst

import (
	"github.com/shellagilehub/practise/structs"
	"math"
)

func FloorForKey(root *structs.Node, key int) int {
	if root == nil {
		return math.MaxInt32
	}
	if root.Data == key {
		return key
	}
	//if root.Data is greater than key then the floor is in right subtree
	if root.Data > key {
		return FloorForKey(root.Left, key)
	}
	// else either the left subtree or root has the floor value
	floor := FloorForKey(root.Right, key)

	if floor <= key {
		return floor
	}
	return root.Data
}
