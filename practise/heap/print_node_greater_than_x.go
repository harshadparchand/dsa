package heap

import (
	"fmt"
	"github.com/shellagilehub/practise/structs"
)

func PrintNodesGreaterThanK(root *structs.Node, k int) {
	if root == nil {
		return
	}
	if root.Data >= k {
		return
	}
	fmt.Print(root.Data)
	fmt.Print(" ")
	PrintNodesGreaterThanK(root.Left, k)
	PrintNodesGreaterThanK(root.Right, k)
	return
}
