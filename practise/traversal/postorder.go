package traversal

import (
	"fmt"
	"github.com/shellagilehub/practise/structs"
)

func PostOrderRecursiveTraversal(node *structs.Node) {
	if node == nil {
		return
	}
	PostOrderRecursiveTraversal(node.Left)
	PostOrderRecursiveTraversal(node.Right)
	fmt.Println(node.Data)
}
