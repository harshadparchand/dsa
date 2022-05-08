package recursion

import "fmt"

func GenerateBalancedParenthesis(open, close, output string) {
	if open == "" && close == "" {
		fmt.Println(output)
		return
	}
	if open != "" {
		op1 := output + string(open[0])
		GenerateBalancedParenthesis(open[1:], close, op1)
	}
	if len(close) > len(open) {
		op2 := output + string(close[0])
		GenerateBalancedParenthesis(open, close[1:], op2)
	}
	return
}
