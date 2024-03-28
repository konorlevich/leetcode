// Package majority_element
//
// Given an array nums of size n, return the majority element.
//
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.
package majority_element

// iterateWithMap
//
// Straight-forward solution.
// Iterate over the array, save nums count to a map.
// Then iterate over the map, search for a maximum counter.
func iterateWithMap(nums []int) (majority int) {
	if len(nums) == 1 {
		return nums[0]
	}
	known := make(map[int]int, len(nums)/2+len(nums)%2)
	for _, num := range nums {
		if _, ok := known[num]; ok {
			known[num]++
			continue
		}
		known[num] = 1
	}

	maxVal := 0
	for i, i2 := range known {
		if i2 > maxVal {
			majority = i
			maxVal = i2
		}
	}
	return majority
}

// countVotes
//
// Not really obvious solution.
// Iterate over the array, count "votes" for each number.
// If we see the same number, increment counter, decrement otherwise.
// If the counter falls to 0, change the candidate to the current num.
func countVotes(nums []int) int {
	candidate := nums[0]
	voted := 1
	for _, n := range nums[1:] {
		if n == candidate {
			voted++
		} else if voted == 1 {
			candidate = n
		} else {
			voted--
		}
	}
	return candidate
}

// iterateWithMapTillMoreThanHalf
//
// According to the task, "the majority element is the element that appears more than ⌊n / 2⌋ times."
// So we can use the approach from iterateWithMap, but we don't need to iterate over the map.
// On each incrementation we can check if the counter > `len(num)/2`
func iterateWithMapTillMoreThanHalf(nums []int) int {
	known := make(map[int]int, len(nums)/2+len(nums)%2)
	majorityCount := len(nums) / 2
	for _, num := range nums {
		if _, ok := known[num]; !ok {
			known[num] = 1
		} else {
			known[num]++
		}
		if known[num] > majorityCount {
			return num
		}
	}
	return 0
}
