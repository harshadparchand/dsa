package recursion

import "fmt"

func GenerateStringsWithSpaces(input string) {
	op := string(input[0])
	input = input[1:len(input)]
	recurGenerateStringsWithSpaces(input, op)
}

func recurGenerateStringsWithSpaces(input, op string) {
	if input == "" {
		fmt.Println(op)
		return
	}
	op1 := op + string(input[0])
	op2 := op + "_" + string(input[0])
	input = input[1:len(input)]
	recurGenerateStringsWithSpaces(input, op1)
	recurGenerateStringsWithSpaces(input, op2)
	return
}
