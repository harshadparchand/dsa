package stack

import "github.com/shellagilehub/practise/structs"

//https://www.geeksforgeeks.org/find-expression-duplicate-parenthesis-not/
/*
Below expressions have duplicate parenthesis -
((a+b)+((c+d)))
The subexpression "c+d" is surrounded by two
pairs of brackets.

(((a+(b)))+(c+d))
The subexpression "a+(b)" is surrounded by two
pairs of brackets.

(((a+(b))+c+d))
The whole expression is surrounded by two
pairs of brackets.

((a+(b))+(c+d))
(b) and ((a+(b)) is surrounded by two
pairs of brackets.

Below expressions don't have any duplicate parenthesis -
((a+b)+(c+d))
No subsexpression is surrounded by duplicate
brackets.
*/

func HasDuplicateParenthesis(input string) bool {
	if len(input) == 0 {
		return false
	}

	stack := structs.StackString{}

	for _, ch := range input {
		if string(ch) == "(" || string(ch) == "+" {
			stack = stack.Push(string(ch))
		} else if string(ch) == ")" {
			val1 := stack.Peek()
			stack = stack.Pop()
			if val1 != "+" {
				return true
			}
			val2 := stack.Peek()
			stack = stack.Pop()
			if val2 != "(" {
				return true
			}
		} else {
			continue
		}
	}

	return false
}
