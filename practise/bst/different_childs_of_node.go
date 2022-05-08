package bst

import "github.com/shellagilehub/practise/structs"

func LeftmostChildOfNode(root *structs.Node) *structs.Node {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		return LeftmostChildOfNode(root.Left)
	}
	return root
}

func RightmostChildOfNode(root *structs.Node) *structs.Node {
	if root == nil {
		return nil
	}
	if root.Right != nil {
		return RightmostChildOfNode(root.Right)
	}
	return root
}

//Smallest node in left subtree
//Smallest node in right subtree
//Smallest node in right subtree
//Largest node in right subtree
