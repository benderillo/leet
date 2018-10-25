package strings

import (
	"fmt"
	"math"
)

// Challenge:
// Implement atoi which converts a string to an integer.

// The function first discards as many whitespace characters as necessary until the first non-whitespace character is found.
// Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

// The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.
// If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.
// If no valid conversion could be performed, a zero value is returned.

// Note:
// Only the space character ' ' is considered as whitespace character.
// Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1].
// If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.

//Solution:
// Time complexity is O(n)
func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}

	//Step1: trim leading space (only ' ')
	start := 0
	for start < len(str) && str[start] == ' ' {
		start++
	}

	str = str[start:]
	if len(str) == 0 {
		return 0
	}
	negative := false
	//Step2: test for leading +/-, creates boolean for it
	if str[0] == '-' {
		negative = true
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}

	//Step3: read digits until they last
	end := 0
	for i := 0; i < len(str); i++ {
		if isNumber(str[i]) {
			end++
		} else {
			break
		}
	}

	// no digits were found
	if end == 0 {
		return 0
	}

	str = str[:end]

	// Step4: convert and check for overflow
	var result int

	for cur := 0; cur < len(str); cur++ {
		digit := int(str[cur] - '0')
		result = result*10 + digit

		if result > math.MaxInt32 {
			result = math.MaxInt32
			if negative {
				result++
			}
			break
		}
	}

	if negative {
		return result * -1
	}

	return result
}

func isNumber(char byte) bool {
	return (char >= '0') && (char <= '9')
}

func main() {
	// fmt.Println(math.MaxInt32)
	// fmt.Println(math.MaxInt64)
	// fmt.Println(math.MinInt32)
	// fmt.Println(math.MinInt64)
	fmt.Println(myAtoi("123"))                           // valid
	fmt.Println(myAtoi("1234567 is a number!"))          // valid
	fmt.Println(myAtoi("     1234567 is a number too!")) //valid
	fmt.Println(myAtoi("+1234567"))                      //valid
	fmt.Println(myAtoi("-1234567"))                      //valid
	fmt.Println(myAtoi("Number 1234567 is not!"))        // invalid
	fmt.Println(myAtoi("1234567889897987"))              // too big
	fmt.Println(myAtoi("-1234567889897987"))             // too small
	fmt.Println(myAtoi(""))                              // invalid
	fmt.Println(myAtoi(" "))                             // invalid
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("2147483648"))
	fmt.Println(myAtoi("-2147483648"))
	fmt.Println(myAtoi("9223372036854775808"))
	fmt.Println(myAtoi("-9223372036854775808"))
	fmt.Println(myAtoi("-2147483649"))
	fmt.Println(myAtoi("-9223372036854775809"))
	fmt.Println(myAtoi("-6147483648"))
	fmt.Println(myAtoi("-2147483647"))
}
