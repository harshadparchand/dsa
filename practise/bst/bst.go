package bst

import (
	"github.com/shellagilehub/practise/structs"
)

func InsertRecursive(root *structs.Node, data int) *structs.Node {
	//you have reached the leaf node
	if root == nil {
		root = &structs.Node{Data: data}
		return root
	}

	//Do nothing if the data is already present int the BST
	if root.Data == data {
		return root
	}

	if root.Data > data {
		root.Left = InsertRecursive(root.Left, data)
	} else {
		root.Right = InsertRecursive(root.Right, data)
	}

	//return the node pointer
	return root
}

func InsertIterative(root *structs.Node, data int) *structs.Node {
	//if root is empty then create one and return it
	if root == nil {
		root = &structs.Node{Data: data}
		return root
	}

	//Iterate on the root node to get the node where the data should be inserted
	var prev *structs.Node
	temp := root
	for temp != nil {
		prev = temp
		if temp.Data > data {
			temp = temp.Left
		} else {
			temp = temp.Right
		}
	}

	if prev.Data > data {
		prev.Left = &structs.Node{Data: data}
	} else {
		prev.Right = &structs.Node{Data: data}
	}

	return root
}

func DeleteRecursive(root *structs.Node, data int) *structs.Node {
	if root == nil {
		return nil
	}

	if data < root.Data {
		root.Left = DeleteRecursive(root.Left, data)
	} else if data > root.Data {
		root.Right = DeleteRecursive(root.Right, data)
	} else {
		// if node has only one child or no child
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// if node has two children: Get the inorder successor (smallest in the right subtree)
		root.Data = SmallestNodeInRightSubtree(root.Right).Data
	}
	return root
}

func SmallestNodeInRightSubtree(root *structs.Node) *structs.Node {
	if root == nil {
		return nil
	}

	var parent *structs.Node
	for root.Left != nil {
		parent = root
		root = root.Left
	}
	parent.Left = nil
	return root
}

var SeqExistIndex = 0

func SeqExist(root *structs.Node, seq []int) {
	if root == nil {
		return
	}
	SeqExist(root.Left, seq)
	if SeqExistIndex < len(seq) && root.Data == seq[SeqExistIndex] {
		SeqExistIndex++
	}
	SeqExist(root.Right, seq)
}
