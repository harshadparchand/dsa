package queue

import "github.com/shellagilehub/practise/structs"

func NumberOfSiblingsOfBSTNode(root *structs.Node, k int) int {
	if root == nil {
		return 0
	}
	queue := structs.QueueNode{}
	queue = queue.Enqueue(root)
	for !queue.IsEmpty() {
		node := queue.Peek()
		queue = queue.Dequeue()
		if node.Left != nil && node.Left.Data == k && node.Right != nil {
			return 1
		}
		if node.Right != nil && node.Right.Data == k && node.Left != nil {
			return 1
		}

		if node.Left != nil {
			queue = queue.Enqueue(node.Left)
		}
		if node.Right != nil {
			queue = queue.Enqueue(node.Right)
		}
	}
	return 0
}
