package strings

import (
	"fmt"
)

func reverseString(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	bytes := []byte(s)

	for start, end := 0, len(bytes)-1; start < end; {
		tmp := bytes[start]
		bytes[start] = bytes[end]
		bytes[end] = tmp
		start++
		end--
	}

	return string(bytes)
}

func main() {
	fmt.Println(reverseString("A man, a plan, a canal: Panama"))
}
