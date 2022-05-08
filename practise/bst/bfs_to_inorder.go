package bst

import "fmt"

func BSFToInorder(arr []int, start, end int) {
	if start > end {
		return
	}
	BSFToInorder(arr, (start*2)+1, end)
	fmt.Println(arr[start])
	BSFToInorder(arr, (start*2)+2, end)
}
