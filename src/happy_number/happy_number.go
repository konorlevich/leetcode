// Package happy_number
//
// Write an algorithm to determine if a number n is happy.
//
// A happy number is a number defined by the following process:
//
// Starting with any positive integer, replace the number by the sum of the squares of its digits.
// Repeat the process until the number equals 1 (where it will stay),
// or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy.
//
// Return true if n is a happy number, and false if not.
package happy_number

import (
	"github.com/konorlevich/leetcode/src/common"
)

// simple
//
// We save each digit (except 0s) to a slice.
// Then we check if we left with a single 2, which is a dead end.
// Then we sum up the squared digits and repeat till we have 1.
func simple(n int) bool {
	for n != 1 {
		nums := make([]int, 0, 30)
		for n != 0 {
			reminder := n % 10
			if reminder != 0 {
				nums = append(nums, reminder)
			}
			n = n / 10
		}
		if len(nums) == 1 && nums[0] == 2 {
			return false
		}
		for _, num := range nums {
			n += common.Pow(num, 2)
		}
	}
	return true
}

func simpleWithMap(n int) bool {
	checked := map[int]struct{}{}
	for n != 1 {
		if _, ok := checked[n]; ok {
			return false
		}
		checked[n] = struct{}{}
		nums := make([]int, 0, 30)
		for n != 0 {
			reminder := n % 10
			if reminder != 0 {
				nums = append(nums, reminder)
			}
			n = n / 10
		}
		for _, num := range nums {
			n += common.Pow(num, 2)
		}
	}
	return true
}
