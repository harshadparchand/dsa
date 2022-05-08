package recursion

import (
	"fmt"
	"github.com/shellagilehub/practise/structs"
)

//https://leetcode.com/problems/find-the-winner-of-the-circular-game/

func KillWithArray(arr []int, k, index int) {
	if len(arr) == 1 {
		fmt.Println(arr[0])
		return
	}
	index = (index + k - 1) % len(arr)
	arr = append(arr[0:index], arr[index+1:]...)
	KillWithArray(arr, k, index)
	return
}

func Kill(n, k int) {
	queue := structs.QueueInt{}
	for i := 1; i <= n; i++ {
		queue = queue.Enqueue(i)
	}
	for queue.Size() != 1 {
		for i := 0; i < k; i++ {
			val := queue.Peek()
			queue = queue.Dequeue()
			if i != k-1 {
				queue = queue.Enqueue(val)
			}
			if queue.Size() == 1 {
				break
			}
		}
		if queue.Size() == 1 {
			break
		}
	}
	fmt.Println(queue.Peek())
}
