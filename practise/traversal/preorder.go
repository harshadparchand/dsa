package traversal

import (
	"fmt"
	"github.com/shellagilehub/practise/structs"
)

func PreOrderRecursiveTraversal(node *structs.Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Data)
	PreOrderRecursiveTraversal(node.Left)
	PreOrderRecursiveTraversal(node.Right)
}
