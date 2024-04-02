// Package power_of_three
//
// Given an integer n, return true if it is a power of three. Otherwise, return false.
//
// An integer n is a power of three, if there exists an integer x such that n == 3x.
//
// Follow up: Could you solve it without loops/recursion?

package power_of_three

func recursive(n int) bool {
	// 3^0
	if n == 1 {
		return true
	}
	if n < 3 {
		return false
	}
	//3^1
	if n == 3 {
		return true
	}

	if n%3 != 0 {
		return false
	}

	return recursive(n / 3)
}

func iterative(n int) bool {
	if n <= 0 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}

	//3^1
	return n == 1
}

// math
//
// Found it out there.
//
// The trick here is that the max n is (2^31) - 1 = 2147483647
//
// The last (3^x) less than that is (3^19) = 1162261467
//
// If we divide this number by any (3^x), where x<=19, we won't get any reminder.
func math(n int) bool {
	return n > 0 && 1_162_261_467%n == 0
}
