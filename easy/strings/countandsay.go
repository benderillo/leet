package strings

import (
	"fmt"
	"strconv"
)

// Challenge:
// 1 is read off as "one 1" or 11.
// 11 is read off as "two 1s" or 21.
// 21 is read off as "one 2, then one 1" or 1211.

// Given an integer n where 1 ≤ n ≤ 30, generate the nth term of the count-and-say sequence.
// Note: Each term of the sequence of integers will be represented as a string.

// Hint:
// The following are the terms from n=1 to n=10 of the count-and-say sequence:
//  1.     1
//  2.     11
//  3.     21
//  4.     1211
//  5.     111221
//  6.     312211
//  7.     13112221
//  8.     1113213211
//  9.     31131211131221
// 10.     13211311123113112211

// Solution:
// The idea is simple, we generate all terms from 1 to n.
// First two terms are initialized as “1” and “11”, and all other terms are generated using previous terms.
// To generate a term using previous term, we scan the previous term. While scanning a term, we simply keep track of count of all consecutive characters.
// For sequence of same characters, we append the count followed by character to generate the next term.

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	if n == 2 {
		return "11"
	}

	sequence := "11"
	for seq := 3; seq <= n; seq++ {
		newSequence := ""
		prevNum := sequence[0]
		count := 0
		for i := 0; i < len(sequence); i++ {
			if sequence[i] != prevNum && count > 0 {
				newSequence += strconv.Itoa(count)
				newSequence += strconv.Itoa(int(prevNum - '0'))
				prevNum = sequence[i]
				count = 1
			} else if sequence[i] == prevNum {
				count++
			}
		}

		if count > 0 {
			newSequence += strconv.Itoa(count)
			newSequence += strconv.Itoa(int(prevNum - '0'))
		}
		sequence = newSequence
	}

	return sequence
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i, countAndSay(i))
	}
}
