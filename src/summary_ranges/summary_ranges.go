// Package summary_ranges
//
// You are given a sorted unique integer array nums.
//
// A range [a,b] is the set of all integers from a to b (inclusive).
//
// Return the smallest sorted list of ranges that cover all the numbers in the array exactly.
// That is, each element of nums is covered by exactly one of the ranges,
// and there is no integer x such that x is in one of the ranges but not in nums.
//
// Each range [a,b] in the list should be output as:
//
// "a->b" if a != b
// "a" if a == b
package summary_ranges

import (
	"fmt"
)

// summaryRanges
func summaryRanges(nums []int) []string {
	out := make([]string, 0, len(nums))
	if len(nums) == 0 {
		return []string{}
	}

	l, r := 0, 1

	for r < len(nums) {
		if nums[r]-nums[r-1] > 1 {
			out = append(out, printLine(nums[l], nums[r-1]))
			l = r
		}
		r++
	}

	out = append(out, printLine(nums[l], nums[r-1]))

	return out
}

func printLine(l, r int) string {
	if l == r {
		return fmt.Sprintf("%d", l)
	}
	return fmt.Sprintf("%d->%d", l, r)
}
