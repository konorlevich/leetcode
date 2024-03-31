// Package square_root
//
// Given a non-negative integer x, return the square root of x rounded down to the nearest integer.
// The returned integer should be non-negative as well.
//
// You must not use any built-in exponent function or operator.
//
// For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
package square_root

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	var left = 0
	var right = x

	for left < right {
		mid := left + ((right - left) / 2)
		ii := mid * mid
		if ii == x {
			return mid
		} else if ii < x {
			left = mid + 1
			continue
		}
		right = mid
	}
	return left - 1
}
