package bst

import "github.com/shellagilehub/practise/structs"

//Convert a BST to a Binary Tree such that sum of all greater keys is added to every key

var sum = 0

func BSTToBinaryTreeSumOfAllGreaterKeysRecursive(root *structs.Node) {
	if root == nil {
		return
	}
	BSTToBinaryTreeSumOfAllGreaterKeysRecursive(root.Right)
	sum += root.Data
	root.Data = sum
	BSTToBinaryTreeSumOfAllGreaterKeysRecursive(root.Left)
}

func BSTToBinaryTreeSumOfAllGreaterKeysIterative(root *structs.Node) {
	if root == nil {
		return
	}
	node := root
	stack := structs.StackNode{}
	sumValue := 0
	for !stack.IsEmpty() || node != nil {
		/* push all nodes up to (and including) this
		 * subtree's maximum on the stack. */
		for node != nil {
			stack = stack.Push(node)
			node = node.Right
		}
		node = stack.Peek()
		stack = stack.Pop()
		sumValue += node.Data
		node.Data = sumValue
		/* all nodes with values between the current and
		 * its parent lie in the left subtree. */
		node = node.Left
	}
}
