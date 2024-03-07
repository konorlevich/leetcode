package remove_element

func removeElement(nums []int, val int) int {
	prevIndex := 0
	for _, num := range nums {
		if num != val {
			nums[prevIndex] = num
			prevIndex++
		}
	}
	return prevIndex
}
