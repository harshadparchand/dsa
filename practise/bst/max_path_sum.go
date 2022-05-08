package bst

import (
	"github.com/shellagilehub/practise/structs"
	"math"
)

var MaxPathSumValue = math.MinInt32

func MaxPathSum(root *structs.Node) int {
	if root == nil {
		return 0
	}
	left := int(math.Max(float64(0), float64(MaxPathSum(root.Left))))
	right := int(math.Max(float64(0), float64(MaxPathSum(root.Right))))
	pathSumWithNode := root.Data + left + right
	if pathSumWithNode > MaxPathSumValue {
		MaxPathSumValue = pathSumWithNode
	}
	return root.Data + int(math.Max(float64(left), float64(right)))
}
