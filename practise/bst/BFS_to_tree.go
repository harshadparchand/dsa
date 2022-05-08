package bst

import "github.com/shellagilehub/practise/structs"

func BFSToTree(arr []int) *structs.Node {
	queue := structs.QueueNode{}
	root := &structs.Node{Data: arr[0]}
	queue = queue.Enqueue(root)
	i := 1
	for i < len(arr) {
		node := queue.Peek()
		queue = queue.Dequeue()
		node.Left = &structs.Node{Data: arr[i]}
		queue = queue.Enqueue(node.Left)
		i++
		if i < len(arr) {
			node.Right = &structs.Node{Data: arr[i]}
			queue = queue.Enqueue(node.Right)
			i++
		}
	}
	return root
}
