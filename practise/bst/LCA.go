package bst

import "github.com/shellagilehub/practise/structs"

func LCAIterative(root *structs.Node, value1, value2 int) *structs.Node {
	if root == nil {
		return nil
	}

	for root != nil {
		//if value1 and value2 < root.Data, then LCA lies in left
		if root.Data > value1 && root.Data > value2 {
			root = root.Left
		} else if root.Data < value1 && root.Data < value2 { //if value1 and value2 > root.Data, then LCA lies in right
			root = root.Right
		} else {
			break
		}
	}

	return root
}

func LCARecursive(root *structs.Node, value1, value2 int) *structs.Node {
	if root == nil {
		return nil
	}

	//if value1 and value2 < root.Data, then LCA lies in left
	if root.Data < value1 && root.Data < value2 {
		return LCARecursive(root.Right, value1, value2)
	}
	//if value1 and value2 > root.Data, then LCA lies in right
	if root.Data > value1 && root.Data > value2 {
		return LCARecursive(root.Left, value1, value2)
	}

	return root
}
