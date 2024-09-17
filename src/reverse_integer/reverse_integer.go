// Package reverse_integer
//
// Given a signed 32-bit integer x, return x with its digits reversed.
// If reversing x causes the value to go outside the signed 32-bit integer range [-2^31, 2^31 - 1], then return 0.
//
// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
//
// Constraints:
//
// -2^31 <= x <= 2^31 - 1
package reverse_integer

import (
	"github.com/konorlevich/leetcode/src/common"
)

var (
	min = -common.Pow(2, 31)
	max = common.Pow(2, 31) - 1
)

// reverse
func reverse(x int) (y int) {
	if x <= min || x >= max {
		return 0
	}

	for x != 0 {
		y = y*10 + x%10
		x /= 10
		// Remember?
		//> Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
		if y > max || y <= min {
			return 0
		}
	}

	return
}
