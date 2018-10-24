package strings

import (
	"fmt"
)

// Challenge:
// Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
// Note: For the purpose of this problem, we define empty string as valid palindrome.

// Solution:
// Most straightforward solution is to loop the string from both sides character by character and compare them
// For this particular challenge, loop should skip non alpha-numeric characters when looping
// Comparison of symbols should be case insensitive
// Time complexity O(n)
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	start := -1
	end := len(s)

	for {
		start = nextAlphanumeric(s, start)
		end = prevAlphanumeric(s, end)

		// Finished scanning the string
		if start >= end {
			break
		}

		if toUpperCase(s[start]) != toUpperCase(s[end]) {
			return false
		}
	}
	return true
}

func nextAlphanumeric(str string, start int) int {
	for {
		start++
		if start == (len(str)-1) || isAlphanumeric(str[start]) {
			break
		}
	}
	return start
}

func prevAlphanumeric(str string, start int) int {
	for {
		start--
		if start == 0 || isAlphanumeric(str[start]) {
			break
		}
	}
	return start
}

func isAlphanumeric(char byte) bool {
	if (char >= '0' && char <= '9') || (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
		return true
	}

	return false
}

func toUpperCase(char byte) byte {
	if char >= 'a' && char <= 'z' {
		return 'A' + char - 'a'
	}

	return char
}

func main() {
	fmt.Println(isPalindrome(".."))
	fmt.Println(isPalindrome(". ."))
	fmt.Println(isPalindrome(",."))
	fmt.Println(isPalindrome("Superman"))
	fmt.Println(isPalindrome(",a."))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("Never odd or even."))
}
