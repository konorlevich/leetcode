// Package squares_of_a_sorted_array
//
// Given an integer array nums sorted in non-decreasing order,
// return an array of the squares of each number sorted in non-decreasing order.
//
// Constraints:
//
//	1 <= nums.length <= 104
//	-104 <= nums[i] <= 104
//	nums is sorted in non-decreasing order.
//
// Follow up:
//
//	Squaring each element and sorting the new array is very trivial,
//	could you find an O(n) solution using a different approach?
package squares_of_a_sorted_array

func sortedSquares(nums []int) []int {
	res := make([]int, 0, len(nums))
	pNum := 0
LOOP:
	for _, num := range nums {
		pNum = num * num
		for i := range res {
			if pNum <= res[i] {
				if i == 0 {
					res = append([]int{pNum}, res...)
					continue LOOP
				}
				res = append(res[:i+1], res[i:]...)
				res[i] = pNum
				continue LOOP
			}
		}
		res = append(res, pNum)

	}
	return res
}
