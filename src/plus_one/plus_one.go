package plus_one

func plusOne(digits []int) []int {
	lastIndex := len(digits) - 1
	if lastIndex < 0 {
		return []int{1}
	}
	if digits[lastIndex] == 9 {
		return append(plusOne(digits[:lastIndex]), 0)
	}
	digits[lastIndex]++
	return digits
}
