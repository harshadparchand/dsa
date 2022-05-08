package queue

import "github.com/shellagilehub/practise/structs"

func NumberOfCousinsOfBSTNode(root *structs.Node, k int) int {
	if root == nil {
		return -1
	}
	queue := structs.QueueNode{}
	queue = queue.Enqueue(root)
	found := false
	for !queue.IsEmpty() && !found {
		size := queue.Size()
		for i := 0; i < size; i++ {
			node := queue.Peek()
			queue = queue.Dequeue()
			//If the element is found in child elements the
			if (node.Left != nil && node.Left.Data == k) || (node.Right != nil && node.Right.Data == k) {
				found = true
			} else {
				if node.Left != nil {
					queue = queue.Enqueue(node.Left)
				}
				if node.Right != nil {
					queue = queue.Enqueue(node.Right)
				}
			}
		}
	}
	if queue.IsEmpty() {
		return -1
	}
	return queue.Size()
}
