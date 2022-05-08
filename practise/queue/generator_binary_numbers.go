package queue

import (
	"fmt"
	"github.com/shellagilehub/practise/structs"
)

func GenerateBinaryNumbers(numbers int) {
	queue := structs.QueueString{}
	queue = queue.Enqueue("1")
	numbers--
	for numbers > 0 {
		s1 := queue.Peek()
		queue = queue.Dequeue()
		fmt.Println(s1)
		if numbers > 0 {
			queue = queue.Enqueue(s1 + "0")
			numbers--
		}
		if numbers > 0 {
			queue = queue.Enqueue(s1 + "1")
			numbers--
		}

	}
	for !queue.IsEmpty() {
		num := queue.Peek()
		queue = queue.Dequeue()
		fmt.Println(num)
	}
}
