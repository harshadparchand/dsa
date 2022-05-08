package tree

import (
	"github.com/shellagilehub/practise/structs"
	"math"
)

func TreeHeightRecursive(root *structs.Node) int {
	if root == nil {
		return 0
	}
	return 1 + int(math.Max(float64(TreeHeightRecursive(root.Left)), float64(TreeHeightRecursive(root.Right))))
}

func TreeHeightIterative(root *structs.Node) int {
	if root == nil {
		return 0
	}
	q := structs.QueueNode{}
	q = q.Enqueue(root)
	height := 0
	for !q.IsEmpty() {
		loop := q.Size()
		for i := 0; i < loop; i++ {
			node := q.Peek()
			q = q.Dequeue()
			if node.Left != nil {
				q = q.Enqueue(node.Left)
			}
			if node.Right != nil {
				q = q.Enqueue(node.Right)
			}
		}
		height++
	}
	return height
}
