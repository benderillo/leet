package arrays

import (
	"fmt"
)

func plusOne(digits []int) []int {
	carry := 0
	result := append([]int(nil), digits...)
	pos := len(result) - 1
	for {
		result[pos] = (digits[pos] + 1) % 10
		carry = (digits[pos] + 1) / 10

		if carry == 0 || pos == 0 {
			break
		}
		pos--
	}
	if carry > 0 {
		result = append([]int{carry}, result...)
	}
	return result
}

func main() {
	a := []int{9, 9, 9}
	fmt.Println(plusOne(a))
}
