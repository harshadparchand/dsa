package stack

import "github.com/shellagilehub/practise/structs"

func SimplifyEquation(s string) string {
	n := len(s)

	res := make([]rune, n)
	index, i := 0, 0

	st := structs.StackInt{}
	st = st.Push(0)

	for i < n {
		if s[i] == '+' {
			if st.Peek() == 1 {
				res[index] = '-'
				index++
			}
			if st.Peek() == 0 {
				res[index] = '+'
				index++
			}
		} else if s[i] == '-' {
			if st.Peek() == 1 {
				res[index] = '+'
				index++
			}
			if st.Peek() == 0 {
				res[index] = '-'
				index++
			}
		} else if s[i] == '(' && i > 0 {
			if s[i-1] == '-' {
				x := st.Peek()
				if x == 1 {
					st = st.Push(0)
				} else {
					st = st.Push(1)
				}
			} else if s[i-1] == '+' {
				x := st.Peek()
				st = st.Push(x)
			}
		} else if s[i] == ')' && i > 0 {
			st = st.Pop()
		} else {
			res[index] = rune(s[i])
			index++
		}
		i++
	}
	return string(res)
}
