package median_of_two_sorted_arrays

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{name: "empty"},

		//Explanation: merged array = [1,2,3] and median is 2.
		{name: "nums1 = [1,3], nums2 = [2]",
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  float64(2),
		},

		//Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
		{name: "nums1 = [1,2], nums2 = [3,4]",
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.50000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := findMedianSortedArrays(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
