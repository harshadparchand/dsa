package tree

import (
	"github.com/shellagilehub/practise/structs"
)

func CreateTree() *structs.Node {
	return &structs.Node{
		Data: 80,
		Left: &structs.Node{
			Data: 40,
			Left: &structs.Node{
				Data:  20,
				Left:  nil,
				Right: nil,
			},
			Right: &structs.Node{
				Data:  50,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &structs.Node{
			Data: 160,
			Left: &structs.Node{
				Data:  100,
				Left:  nil,
				Right: nil,
			},
			Right: &structs.Node{
				Data:  320,
				Left:  nil,
				Right: nil,
			},
		},
	}
}

func CountNodes(root *structs.Node) int {
	if root == nil {
		return 0
	}

	return 1 + CountNodes(root.Left) + CountNodes(root.Right)
}
