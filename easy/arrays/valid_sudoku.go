package arrays

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	hash := make(map[int]byte)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			item := board[i][j]

			// Calc Id of the board
			if item == byte('.') {
				continue
			}

			boardID := ((i/3)*3+(j/3)+1)*10 + int(item)
			columnID := (j+1)*100 + int(item)
			rowID := (i+1)*1000 + int(item)

			if _, ok := hash[boardID]; ok {
				return false
			}

			if _, ok := hash[columnID]; ok {
				return false
			}

			if _, ok := hash[rowID]; ok {
				return false
			}
			hash[boardID] = item
			hash[columnID] = item
			hash[rowID] = item
		}
	}

	return true
}

func main() {
	valid := [][]byte{
		{5, 3, byte('.'), byte('.'), 7, byte('.'), byte('.'), byte('.'), byte('.')},
		{6, byte('.'), byte('.'), 1, 9, 5, byte('.'), byte('.'), byte('.')},
		{byte('.'), 9, 8, byte('.'), byte('.'), byte('.'), byte('.'), 6, byte('.')},
		{8, byte('.'), byte('.'), byte('.'), 6, byte('.'), byte('.'), byte('.'), 3},
		{4, byte('.'), byte('.'), 8, byte('.'), 3, byte('.'), byte('.'), 1},
		{7, byte('.'), byte('.'), byte('.'), 2, byte('.'), byte('.'), byte('.'), 6},
		{byte('.'), 6, byte('.'), byte('.'), byte('.'), byte('.'), 2, 8, byte('.')},
		{byte('.'), byte('.'), byte('.'), 4, 1, 9, byte('.'), byte('.'), 5},
		{byte('.'), byte('.'), byte('.'), byte('.'), 8, byte('.'), byte('.'), 7, 9},
	}

	invalid := [][]byte{
		{8, 3, byte('.'), byte('.'), 7, byte('.'), byte('.'), byte('.'), byte('.')},
		{6, byte('.'), byte('.'), 1, 9, 5, byte('.'), byte('.'), byte('.')},
		{byte('.'), 9, 8, byte('.'), byte('.'), byte('.'), byte('.'), 6, byte('.')},
		{8, byte('.'), byte('.'), byte('.'), 6, byte('.'), byte('.'), byte('.'), 3},
		{4, byte('.'), byte('.'), 8, byte('.'), 3, byte('.'), byte('.'), 1},
		{7, byte('.'), byte('.'), byte('.'), 2, byte('.'), byte('.'), byte('.'), 6},
		{byte('.'), 6, byte('.'), byte('.'), byte('.'), byte('.'), 2, 8, byte('.')},
		{byte('.'), byte('.'), byte('.'), 4, 1, 9, byte('.'), byte('.'), 5},
		{byte('.'), byte('.'), byte('.'), byte('.'), 8, byte('.'), byte('.'), 7, 9},
	}

	case2 := [][]byte{
		{byte('.'), byte('.'), 4, byte('.'), byte('.'), byte('.'), 6, 3, byte('.')},
		{byte('.'), byte('.'), byte('.'), byte('.'), byte('.'), byte('.'), byte('.'), byte('.'), byte('.')},
		{5, byte('.'), byte('.'), byte('.'), byte('.'), byte('.'), byte('.'), 9, byte('.')},
		{byte('.'), byte('.'), byte('.'), 5, 6, byte('.'), byte('.'), byte('.'), byte('.')},
		{4, byte('.'), 3, byte('.'), byte('.'), byte('.'), byte('.'), byte('.'), 1},
		{byte('.'), byte('.'), byte('.'), 7, byte('.'), byte('.'), byte('.'), byte('.'), byte('.')},
		{byte('.'), byte('.'), byte('.'), 5, byte('.'), byte('.'), byte('.'), byte('.'), byte('.')},
		{byte('.'), byte('.'), byte('.'), byte('.'), byte('.'), byte('.'), byte('.'), byte('.'), byte('.')},
		{byte('.'), byte('.'), byte('.'), byte('.'), byte('.'), byte('.'), byte('.'), byte('.'), byte('.')},
	}
	fmt.Println(isValidSudoku(valid))
	fmt.Println(isValidSudoku(invalid))
	fmt.Println(isValidSudoku(case2))
}
