package array

import (
	"github.com/shellagilehub/practise/structs"
	"strconv"
	"unicode"
)

//https://leetcode.com/problems/decode-string/
//Testing
//fmt.Println(array.DecodeString("2[ac]2[bc]"))

func DecodeString(s string) string {
	chars := []rune(s)
	stack := structs.StackChar{}
	for i := 0; i < len(chars); i++ {
		//push chars to stack until we come across close parenthesis
		if string(chars[i]) != "]" {
			stack = stack.Push(chars[i])
		} else {
			var subStr string
			//pop content from stack until we come across open parenthesis
			for !stack.IsEmpty() && string(stack.Peek()) != "[" {
				subStr = string(stack.Peek()) + subStr
				stack = stack.Pop()
			}
			//pop open parenthesis
			stack = stack.Pop()
			numStr := ""
			//pop until you see numbers
			for !stack.IsEmpty() && unicode.IsNumber(stack.Peek()) {
				numStr = string(stack.Peek()) + numStr
				stack = stack.Pop()
			}
			str := ""
			k, _ := strconv.Atoi(numStr)
			//form string for number * string
			for i := 0; i < k; i++ {
				str += subStr
			}
			strNum := []rune(str)
			//push number*string back into the stack
			for _, char := range strNum {
				stack = stack.Push(char)
			}
		}
	}
	res := ""
	for !stack.IsEmpty() {
		res = string(stack.Peek()) + res
		stack = stack.Pop()
	}
	return res
}
