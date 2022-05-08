package queue

import "github.com/shellagilehub/practise/structs"

//https://www.geeksforgeeks.org/check-if-a-given-binary-tree-is-complete-tree-or-not/

func IsBST(root *structs.Node) bool {
	if root == nil {
		return false
	}
	queue := structs.QueueNode{}
	queue = queue.Enqueue(root)
	hasSeenNull := false
	for !queue.IsEmpty() {
		node := queue.Peek()
		queue = queue.Dequeue()
		if !hasSeenNull && node == nil {
			hasSeenNull = true
		} else if hasSeenNull && node != nil {
			return false
		}
		if node != nil {
			queue = queue.Enqueue(node.Left)
			queue = queue.Enqueue(node.Right)
		}
	}

	return true
}
