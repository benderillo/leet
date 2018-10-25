package main

import "fmt"

//Challenge:
// Implement strStr().
// Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

// Clarification:
// What should we return when needle is an empty string? This is a great question to ask during an interview.
// For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

// Solution:
// Just go through the haystack symbol by symbol and try to compare with the needle
// Time complexity O(n) because at most n comparisons will be made
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(haystack) == 0 {
		return -1
	}

	if len(needle) > len(haystack) {
		return -1
	}

	if len(haystack) == len(needle) {
		if haystack == needle {
			return 0
		}
		return -1
	}

	for i := 0; i < (len(haystack) - len(needle) + 1); i++ {

		j := 0
		for j < len(needle) {
			if haystack[i+j] != needle[j] {
				break
			}
			j++
		}

		if j == len(needle) {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("helllo", "ll"))
	fmt.Println(strStr("hello", "he"))
	fmt.Println(strStr("hello", "o"))
	fmt.Println(strStr("hello", "lo"))
	fmt.Println(strStr("", "d"))
	fmt.Println(strStr("", ""))
	fmt.Println(strStr("he", "hee"))
	fmt.Println(strStr("test", "test"))
	fmt.Println(strStr("ping", "pong"))
	fmt.Println(strStr("aaaa", "bba"))
}
