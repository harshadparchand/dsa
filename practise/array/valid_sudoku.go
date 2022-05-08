package array

import "strconv"

//Testing
//fmt.Println(array.IsValidSudoku([][]int{
//		{5, 1, 0, 0, 7, 0, 0, 0, 0},
//		{6, 0, 0, 1, 9, 5, 0, 0, 0},
//		{0, 9, 8, 0, 0, 0, 0, 6, 0},
//		{8, 0, 0, 0, 6, 0, 0, 0, 3},
//		{4, 0, 0, 8, 0, 3, 0, 0, 1},
//		{7, 0, 0, 0, 2, 0, 0, 0, 6},
//		{0, 6, 0, 0, 0, 0, 2, 8, 0},
//		{0, 0, 0, 4, 1, 9, 0, 0, 5},
//		{0, 0, 0, 0, 8, 0, 0, 7, 9},
//	}))

func IsValidSudoku(board [][]int) bool {
	myMap := make(map[string]bool)

	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if board[row][column] != 0 {
				_, isPresent1 := myMap[strconv.Itoa(board[row][column])+" found in row "+strconv.Itoa(row)]
				_, isPresent2 := myMap[strconv.Itoa(board[row][column])+" found in column "+strconv.Itoa(column)]
				_, isPresent3 := myMap[strconv.Itoa(board[row][column])+" found in box "+strconv.Itoa(row/3)+"-"+strconv.Itoa(column/3)]
				if isPresent1 || isPresent2 || isPresent3 {
					return false
				} else {
					myMap[strconv.Itoa(board[row][column])+" found in row "+strconv.Itoa(row)] = true
					myMap[strconv.Itoa(board[row][column])+" found in column "+strconv.Itoa(column)] = true
					myMap[strconv.Itoa(board[row][column])+" found in box "+strconv.Itoa(row/3)+"-"+strconv.Itoa(column/3)] = true
				}
			}
		}
	}
	return true
}
