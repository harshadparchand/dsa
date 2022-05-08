package bst

import "github.com/shellagilehub/practise/structs"

func FindDistance3Pass(root *structs.Node, n1, n2 int) int {

	//Get LCA of the two nodes
	lca := LCAIterative(root, n1, n2)

	//Get distance of n1 from LCA
	d1 := findLevelRecursive(lca, n1)
	//Get distance of n2 from LCA
	d2 := findLevelRecursive(lca, n2)

	return d1 + d2

}

func FindDistanceSinglePass(root *structs.Node, n1, n2 int) int {

	if root == nil {
		return 0
	}

	if root.Data > n1 && root.Data > n2 {
		return FindDistanceSinglePass(root.Left, n1, n2)
	} else if root.Data < n1 && root.Data < n2 {
		return FindDistanceSinglePass(root.Right, n1, n2)
	} else {
		return findLevelRecursive(root, n1) + findLevelIterative(root, n2)
	}
	return 0
}

func findLevelRecursive(root *structs.Node, key int) int {
	if root == nil || root.Data == key {
		return 0
	}

	if root.Data > key {
		return 1 + findLevelRecursive(root.Left, key)
	} else {
		return 1 + findLevelRecursive(root.Right, key)
	}
}

func findLevelIterative(root *structs.Node, key int) int {
	if root == nil {
		return 0
	}
	level := 0
	for root != nil {
		if root.Data > key {
			root = root.Left
			level++
		} else if root.Data < key {
			root = root.Right
			level++
		} else {
			break
		}
	}

	return level
}
