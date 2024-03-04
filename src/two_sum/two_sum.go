package two_sum

func twoSum(f func([]int, int) []int, nums []int, target int) []int {
	return f(nums, target)
}

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
