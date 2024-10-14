// Package max_consecutive_ones
//
// Given a binary array nums, return the maximum number of consecutive 1's in the array.
//
// Constraints:
//
// 1 <= nums.length <= 105
//
// nums[i] is either 0 or 1.
package max_consecutive_ones

func findMaxConsecutiveOnes(nums []int) (res int) {
	localCount := 0
	for _, num := range nums {
		if num == 0 {
			localCount = 0
			continue
		}
		localCount++

		if localCount > res {
			res = localCount
		}
	}

	return res
}
