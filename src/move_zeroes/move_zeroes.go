// Package move_zeroes
//
// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//
// Note that you must do this in-place without making a copy of the array.
//
// Follow up: Could you minimize the total number of operations done?
package move_zeroes

func twoPointers(nums []int) {
	left := 0
	for left < len(nums) && nums[left] != 0 {
		left++
	}
	right := left + 1
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
