// Package contains_duplicate
//
// Given an integer array nums, return true if any value appears at least twice in the array,
// and return false if every element is distinct.
package contains_duplicate

// inefficient straight-forward solution.
//
// Iterate over the array on each num. Return true if we found an equal number.
func inefficient(nums []int) bool {
	lastIdx := len(nums) - 1
	for i, num1 := range nums {
		if i == lastIdx {
			break
		}
		for _, num2 := range nums[i+1:] {
			if num1 == num2 {
				return true
			}
		}
	}
	return false
}

// we use empty struct because it has a minimum size of zero bytes
var emptyStruct = struct{}{}

// withMap
//
// We iterate over the array, saving checking if we saw this number.
// If not, save it to the map.
func withMap(nums []int) bool {
	checked := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, ok := checked[num]; ok {
			return true
		}
		checked[num] = emptyStruct
	}
	return false
}
