package strings

import (
	"bytes"
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	shortestStrLen := math.MaxInt64
	for i := 0; i < len(strs); i++ {
		str := strs[i]
		if len(str) == 0 {
			// if any of the strings is empty, return empty string as there is no common prefix
			return ""
		}
		// Find shortest string among all
		if len(str) < shortestStrLen {
			shortestStrLen = len(str)
		}
	}
	prefix := bytes.NewBufferString("")
	firstString := strs[0]
	for pos := 0; pos < shortestStrLen; pos++ {
		for i := 1; i < len(strs); i++ {
			if firstString[pos] != strs[i][pos] {
				return prefix.String()
			}
		}
		prefix.WriteString(string(firstString[pos]))
	}

	return prefix.String()
}

func main() {
	a := []string{"abcde", "abc", "abcd"}
	b := []string{}
	c := []string{"", "abc"}
	d := []string{"single string"}
	fmt.Println(longestCommonPrefix(a))
	fmt.Println(longestCommonPrefix(b))
	fmt.Println(longestCommonPrefix(c))
	fmt.Println(longestCommonPrefix(d))
}
