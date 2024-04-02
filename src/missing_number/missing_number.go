// Package missing_number
//
// Given an array nums containing n distinct numbers in the range [0, n],
// return the only number in the range that is missing from the array.
//
// All the numbers of nums are unique.
//
// Follow up: Could you implement a solution using only O(1) extra space complexity and O(n) runtime complexity?
package missing_number

// bruteForce
//
// Straight-forward O(n^2) solution
//
// ATTENTION!
// This function will change the slice
func bruteForce(nums []int) int {
	maxNum := len(nums)
LOOP:
	for needle := 0; needle <= maxNum; needle++ {
		for i, num := range nums {
			if num == needle {
				nums = append(nums[:i], nums[i+1:]...)
				continue LOOP
			}
		}
		return needle
	}
	return maxNum + 1
}

// withTempSlice
//
// A bit faster, it takes O(2n) to find a number.
//
// We create a slice with a length of nums filled with 0s.
// We iterate over nums, filling nth value in a temp slice with n.
// Then we iterate over the temp slice, looking for a value we didn't change (which is 0).
// As an exception, we set 0th value to 0, so we don't confuse it with unchanged value.
func withTempSlice(nums []int) int {
	temp := make([]int, len(nums)+1)
	for _, num := range nums {
		temp[num] = num
		if num == 0 {
			temp[num] = 1
		}
	}
	for i, num := range temp {
		if num == 0 {
			return i
		}
	}
	return len(nums) + 1
}

// sumIt
//
// A great idea I found out there.
//
// The sum of 'n' terms in an AP can be calculated using a straightforward formula: Sn = n/2 [2a + (n-1)d].
// Here, 'a' represents the first term, 'd' is the common difference, and 'n' denotes the number of terms.
//
// So we calculate the sum we are waiting.
// Then we subtract each number from the sum above.
func sumIt(nums []int) (num int) {
	maxNum := len(nums)
	num = (maxNum * (maxNum + 1)) / 2
	for _, n := range nums {
		num -= n
	}
	return num
}
