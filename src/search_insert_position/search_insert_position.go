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
