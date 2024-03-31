// Package plus_one
//
// You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer.
// The digits are ordered from most significant to least significant in left-to-right order.
// The large integer does not contain any leading 0's.
//
// Increment the large integer by one and return the resulting array of digits.
package plus_one

func plusOne(digits []int) []int {
	lastIndex := len(digits) - 1
	if lastIndex < 0 {
		return []int{1}
	}
	if digits[lastIndex] == 9 {
		return append(plusOne(digits[:lastIndex]), 0)
	}
	digits[lastIndex]++
	return digits
}
