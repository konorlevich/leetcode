// Package find_numbers_with_even_number_of_digits
//
// Given an array `nums` of integers, return how many of them contain an even number of digits.
//
// Constraints:
//
// 1 <= nums.length <= 500
//
// 1 <= nums[i] <= 10^5
package find_numbers_with_even_number_of_digits

import (
	"strconv"
)

func calculateNumbers(nums []int) (found int) {
	for _, n := range nums {
		c := 0
		for n > 0 {
			c++
			n /= 10
		}
		if c%2 == 0 {
			found++
		}
	}
	return
}

func convertNumbers(nums []int) (found int) {
	for _, n := range nums {
		if len(strconv.Itoa(n))%2 == 0 {
			found++
		}
	}
	return
}

func compareNumbers(nums []int) (found int) {
	for _, n := range nums {
		if n < 10 || (n >= 100 && n < 1000) || (n >= 10000 && n < 100000) {
			continue
		}
		found++
	}
	return
}
