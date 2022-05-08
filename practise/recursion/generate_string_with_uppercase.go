package recursion

import (
	"fmt"
	"strings"
	"unicode"
)

func GenerateStringsWithUpperCase(input, op string) {
	if input == "" {
		fmt.Println(op)
		return
	}
	if unicode.IsDigit(rune(input[0])) {
		op = op + string(input[0])
		input = input[1:]
		GenerateStringsWithUpperCase(input, op)
	} else {
		op1 := op + strings.ToLower(string(input[0]))
		op2 := op + strings.ToUpper(string(input[0]))
		input = input[1:]
		GenerateStringsWithUpperCase(input, op1)
		GenerateStringsWithUpperCase(input, op2)
	}
	return
}
