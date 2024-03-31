// Package merge_sorted_array
//
// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order,
// and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
//
// Merge nums1 and nums2 into a single array sorted in non-decreasing order.
//
// The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
// To accommodate this, nums1 has a length of m + n,
// where the first m elements denote the elements that should be merged,
// and the last n elements are set to 0 and should be ignored.
// nums2 has a length of n.
package merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int) {
	for indexInNums2, numInNums2 := range nums2 {
		if m == 0 {
			nums1[0] = numInNums2
			m++
			continue
		}
		for indexInNums1 := max(m-1, 0); indexInNums1 >= 0; indexInNums1-- {
			numInNums1 := nums1[indexInNums1]
			if numInNums1 <= numInNums2 {
				for i := m; i > indexInNums1; i-- {
					nums1[i] = nums1[i-1]
				}
				nums1[indexInNums1+1] = nums2[indexInNums2]
				m++
				break
			}
			if indexInNums1 == 0 {
				for i := m; i > indexInNums1; i-- {
					nums1[i] = nums1[i-1]
				}
				nums1[indexInNums1] = numInNums2
				m++
				break
			}
		}
	}
}
