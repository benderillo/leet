package arrays

import (
	"fmt"
)

func rotate_copy(matrix [][]int) {
	kopy := [][]int{}
	for i := range matrix {
		row := make([]int, len(matrix[i]))
		kopy = append(kopy, row)
	}

	n := len(matrix) - 1

	for i, row := range matrix {
		for j := range row {
			kopy[j][n-i] = matrix[i][j]
		}
	}

	for i, row := range kopy {
		copy(matrix[i], row)
	}
}

func rotate_inplace(matrix [][]int) {
	if len(matrix) == 1 {
		return
	}

	n := len(matrix)

	// Each element will have to move, so total of n*n transitions required.
	// This nested loop will go ring by ring, from outer to inner.
	swaps := 0
	for i := 0; swaps < n*n && i < n-1; i++ {
		for j := i; swaps < n*n && j < n-i-1; j++ {
			prev := matrix[i][j]
			cur_i := i
			cur_j := j

			// This loop concept I borrowed from rotating 1d array by circular shifts
			// The idea is to pick one element, place it to new position, then the one that was there before shall
			// be moved to its new position too.
			// This shall continue until we come back to the starting position, then we break the loop
			for {
				next_i := cur_j
				next_j := n - 1 - cur_i

				tmp := matrix[next_i][next_j]
				matrix[next_i][next_j] = prev
				prev = tmp
				swaps++
				cur_i = next_i
				cur_j = next_j

				if i == cur_i && j == cur_j {
					break
				}
			}
		}
	}
}

func main() {
	a := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	rotate_inplace(a)
	fmt.Println(a)

}
