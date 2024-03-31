// Package median_of_two_sorted_arrays
//
// Given two sorted arrays `nums1` and `nums2` of size `m` and `n` respectively,
// return the median of the two sorted arrays.
//
// The overall run time complexity should be `O(log(m+n))`.
package median_of_two_sorted_arrays

// findMedianSortedArrays
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	newNum := make([]int, 0, len(nums1)+len(nums2))
	i1, i2 := 0, 0

	for i1 <= len(nums1) && i2 <= len(nums2) {
		if i1 >= len(nums1) {
			newNum = append(newNum, nums2[i2:]...)
			break
		}
		if i2 >= len(nums2) {
			newNum = append(newNum, nums1[i1:]...)
			break
		}
		if nums1[i1] > nums2[i2] {
			newNum = append(newNum, nums2[i2])
			i2++
		} else {
			newNum = append(newNum, nums1[i1])
			i1++
		}
	}
	halfLen := len(newNum) / 2
	if len(newNum)%2 == 1 {
		return float64(newNum[halfLen])
	}
	return float64(newNum[halfLen-1]+newNum[halfLen]) / float64(2)
}
