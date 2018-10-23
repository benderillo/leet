package strings

import (
	"fmt"
)

func reverse(x int) int {
	if (x < 10) && (x > -10) {
		return x
	}

	cutX := int32(x)
	if x != int(cutX) {
		fmt.Println("X is too large")
		return 0
	}

	var result int32
	var newValue int32
	var end int32
	for x != 0 {
		end = int32(x % 10)
		newValue = result*10 + end

		//overflow guard. Undo previous operation and see if it gives same value as it was.
		if (newValue-end)/10 != result {
			return 0
		}

		result = newValue
		x = x / 10
	}

	return int(result)
}

func main() {
	fmt.Println(reverse(-123))
	fmt.Println(reverse(1000000009))
}
