package bst

import (
	"fmt"
	"github.com/shellagilehub/practise/structs"
)

func PrintBSTInGivenRange(root *structs.Node, min, max int) {
	if root == nil {
		return
	}

	if min < root.Data {
		PrintBSTInGivenRange(root.Left, min, max)
	}
	if root.Data >= min && root.Data <= max {
		fmt.Println(root.Data)
	}
	if max > root.Data {
		PrintBSTInGivenRange(root.Right, min, max)
	}
}
