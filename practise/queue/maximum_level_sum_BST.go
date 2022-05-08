package queue

import (
	"github.com/shellagilehub/practise/structs"
	"math"
)

func MaximumLevelSumBST(root *structs.Node) int {
	if root == nil {
		return math.MinInt32
	}
	queue := structs.QueueNode{}
	queue = queue.Enqueue(root)
	res := math.MinInt32
	for !queue.IsEmpty() {
		levelNodeCount := queue.Size()
		levelSum := 0
		for i := 0; i < levelNodeCount; i++ {
			node := queue.Peek()
			queue = queue.Dequeue()
			levelSum += node.Data
			if node.Left != nil {
				queue = queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue = queue.Enqueue(node.Right)
			}

		}
		if res < levelSum {
			res = levelSum
		}
	}
	return res
}
