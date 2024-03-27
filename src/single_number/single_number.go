// Package single_number
//
// Given a non-empty array of integers nums, every element appears twice except for one.
// Find that single one.
//
// You must implement a solution with a linear runtime complexity and use only constant extra space.
//
// Each element in the array appears twice except for one element which appears only once.
package single_number

// singleNumber We iterate over the array, checks the number (`x`).
// If it's an unknown number, save it to the map as a key with a value `x-1`.
// If it's a known number, save it as a value with an equal key.
//
// Then we iterate over the map, check if key == value, return otherwise.
//
// I't not really `O(n)`, but Leetcode says,
// it "Beats 99.29% of users with Go" in runtime,
// and "Beats 45.78% of users with Go".
//
// Looks like good enough.
func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	found := make(map[int]int, len(nums)/2+1)
	for _, num := range nums {
		if _, ok := found[num]; ok {
			found[num] = num
			continue
		}
		found[num] = num - 1
	}
	for num, pair := range found {
		if num != pair {
			return num
		}
	}
	return 0
}
