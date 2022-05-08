package heap

import "github.com/shellagilehub/practise/structs"

func CreateHeap() *structs.Node {
	root := &structs.Node{
		Data: 2,
		Left: &structs.Node{
			Data: 3,
			Left: &structs.Node{
				Data: 5,
				Left: &structs.Node{
					Data: 6,
					Left: nil,
				},
				Right: &structs.Node{
					Data: 150,
					Left: nil,
				},
			},
			Right: &structs.Node{
				Data: 4,
				Left: &structs.Node{
					Data: 77,
					Left: nil,
				},
				Right: &structs.Node{
					Data: 120,
					Left: nil,
				},
			},
		},
		Right: &structs.Node{
			Data: 15,
			Left: &structs.Node{
				Data: 45,
				Left: nil,
			},
			Right: &structs.Node{
				Data: 80,
				Left: nil,
			},
		},
	}
	return root
}
