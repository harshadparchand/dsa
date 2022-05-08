package stack

import (
	"github.com/shellagilehub/practise/structs"
	"strconv"
)

func ReverseNumber(num int) int {
	val := strconv.Itoa(num)
	stack := structs.StackChar{}
	for _, value := range val {
		stack = stack.Push(value)
	}

	var output string
	for !stack.IsEmpty() {
		val := stack.Peek()
		stack = stack.Pop()
		output += string(val)
	}

	num, _ = strconv.Atoi(output)
	return num
}
