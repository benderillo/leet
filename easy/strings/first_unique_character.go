package main

import (
	"fmt"
)

// This will build a map of occurrences for each character in the input string
// Then it will lookup num occurrences for each symbol in the input until find the first with only one occurrence
// So there are only two passes through n characters, hence time complexity is O(n)
func firstUniqChar(s string) int {
	hash := make(map[rune]int)
	// build a map
	for _, v := range s {
		hash[v]++
	}

	for i, v := range s {
		if hash[v] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("cc"))
	fmt.Println(firstUniqChar("loveleetcode"))
}
