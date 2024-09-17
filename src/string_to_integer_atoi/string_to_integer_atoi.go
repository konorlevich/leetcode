// Package string_to_integer_atoi
//
// Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer.
//
// The algorithm for myAtoi(string s) is as follows:
//
// Whitespace: Ignore any leading whitespace (" ").
//
// Signedness: Determine the sign by checking if the next character is '-' or '+', assuming positivity is neither present.
//
// Conversion: Read the integer by skipping leading zeros until a non-digit character is encountered or the end of the string is reached. If no digits were read, then the result is 0.
//
// Rounding: If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then round the integer to remain in the range. Specifically, integers less than -231 should be rounded to -231, and integers greater than 231 - 1 should be rounded to 231 - 1.
//
// Return the integer as the final result.
//
// Constraints:
//
// 0 <= s.length <= 200
//
// s consists of English letters (lower-case and upper-case), digits (0-9), ' ', '+', '-', and '.'.

package string_to_integer_atoi

import (
	"github.com/konorlevich/leetcode/src/common"
	"github.com/konorlevich/leetcode/src/common/ascii"
)

var (
	min = -common.Pow(2, 31)
	max = common.Pow(2, 31) - 1
)

func myAtoi(s string) (n int) {
	numberStarted := false
	signRead := false
	negative := false
	for _, symbol := range s {
		if symbol > ascii.Digit9 || symbol < ascii.Space {
			return
		}
		if symbol == ascii.Space && (numberStarted || signRead) {
			break
		}
		if symbol == ascii.Period {
			break
		}
		if symbol == ascii.Minus || symbol == ascii.Plus {
			if signRead || numberStarted {
				break
			}
			signRead = true
			negative = symbol == ascii.Minus

		}
		if symbol >= ascii.Digit0 {
			numberStarted = true
			n *= 10
			if negative {
				n -= int(symbol) - ascii.Digit0
			} else {
				n += int(symbol) - ascii.Digit0
			}
			if n > max {
				return max
			}
			if n < min {
				return min
			}
		}
	}
	return
}
