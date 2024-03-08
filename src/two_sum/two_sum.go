package two_sum

// twoSum
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// You can return the answer in any order.
func twoSum(f func([]int, int) []int, nums []int, target int) []int {
	return f(nums, target)
}

// bruteForce
//
// Just iterates over the numbers for each number, looking for a perfect match
func bruteForce(nums []int, target int) []int {
	for k, v := range nums {
		for k1, v1 := range nums {
			if k1 == k {
				continue
			}
			if v1+v == target {
				return []int{k, k1}
			}
		}
	}
	return []int{}
}

// withoutCurrentK
//
// Iterates over the numbers k1 > k , for each number k
func withoutCurrentK(nums []int, target int) []int {
	for k, v := range nums {
		for k1, v1 := range nums[k+1:] {
			if v1+v == target {
				return []int{k, k + k1 + 1}
			}
		}
	}
	return []int{}
}
