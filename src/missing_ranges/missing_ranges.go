// Package missing_ranges
//
// You are given an inclusive range `[lower, upper]` and a sorted unique integer array `nums`,
// where all elements are within the inclusive range.
//
// A number `x` is considered missing if `x` is in the range `[lower, upper]` and `x` is not in nums.
//
// Return the shortest sorted list of ranges that exactly covers all the missing numbers.
// That is, no element of `nums` is included in any of the ranges,
// and each missing number is covered by one of the ranges.
//
// All the values of `nums` are unique.
package missing_ranges

// findMissingRanges
//
// Straight-forward solution.
//
// At first, we check if we have any numbers and/or boundaries.
// If no numbers and lower, upper are empty, return an empty array.
// If no numbers, return an array with range [lower, upper].
// Then we check if the first number > lower. If no, add a range [lower, `first num`-1] to the resulting array.
// Then we iterate over the `nums` array, checking if number n and n+1 are consequent.
// If no, add a range to the resulting array.
// Then we check if the last number < upper. If no, add a range [`last num`+1, upper] to the resulting array.
func findMissingRanges(nums []int, lower int, upper int) (missing [][]int) {
	missing = make([][]int, 0, len(nums)+1)
	if lower == 0 && upper == 0 && len(nums) == 0 {
		return missing
	}
	if len(nums) == 0 {
		missing = append(missing, []int{lower, upper})
		return
	}
	if nums[0] > lower {
		missing = append(missing, []int{lower, nums[0] - 1})
	}
	for i, num := range nums {
		if i == len(nums)-1 {
			break
		}
		if nums[i+1]-num > 1 {
			missing = append(missing, []int{num + 1, nums[i+1] - 1})
		}
	}
	if nums[len(nums)-1] < upper {
		missing = append(missing, []int{nums[len(nums)-1] + 1, upper})
	}
	return missing
}
