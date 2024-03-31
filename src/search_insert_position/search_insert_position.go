// Package search_insert_position
//
// Given a sorted array of distinct integers and a target value, return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.
//
// You must write an algorithm with O(log n) runtime complexity.
package search_insert_position

func searchInsert(nums []int, target int) int {
	numLen := len(nums)
	if numLen == 0 {
		return 0
	} else if numLen == 1 {
		if nums[0] >= target {
			return 0
		} else {
			return 1
		}
	}
	key := numLen / 2
	if target > nums[key] {
		return key + searchInsert(nums[key:], target)
	} else if target < nums[key] {
		return searchInsert(nums[:key], target)
	}
	return key
}
