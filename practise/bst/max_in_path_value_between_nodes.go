package bst

import (
	"github.com/shellagilehub/practise/structs"
	"math"
)

func MaxValueBetweenNodes(root *structs.Node, key int) int32 {
	if root == nil {
		return 0
	}
	res := math.MinInt32
	for root != nil {
		if root.Data > res {
			res = root.Data
		}
		if root.Data > key {
			root = root.Left
		} else if root.Data < key {
			root = root.Right
		} else {
			break
		}
	}
	return int32(math.Max(float64(res), float64(key)))
}
