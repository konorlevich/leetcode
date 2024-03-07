package remove_duplicates_from_sorted_array

import "math"

func removeDuplicates(nums []int) int {
	prev := -math.MaxInt
	prevIndex := -1
	for _, num := range nums {
		if num > prev {
			prev = num
			prevIndex++
			nums[prevIndex] = num
			continue
		}
	}
	return prevIndex + 1
}
