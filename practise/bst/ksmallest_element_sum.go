package bst

import "github.com/shellagilehub/practise/structs"

var KSmallestElementSumIndex = 0
var KSmallestElementSumCumulativeSum = 0

func KSmallestElementSum(root *structs.Node, k int) {
	if root == nil {
		return
	}
	if KSmallestElementSumIndex < k {
		KSmallestElementSum(root.Left, k)
	}
	KSmallestElementSumIndex++
	if KSmallestElementSumIndex <= k {
		KSmallestElementSumCumulativeSum += root.Data
	}
	if KSmallestElementSumIndex < k {
		KSmallestElementSum(root.Right, k)
	}
}
