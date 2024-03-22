package merge_sorted_array

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{name: "empty"},
		{name: "1",
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6}},
		{name: "2",
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			want:  []int{1}},
		{name: "3",
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
			want:  []int{1}},
		{name: "4",
			nums1: []int{1, 0},
			m:     1,
			nums2: []int{2},
			n:     1,
			want:  []int{1, 2}},
		{name: "5",
			nums1: []int{2, 0},
			m:     1,
			nums2: []int{1},
			n:     1,
			want:  []int{1, 2}},
		{name: "6",
			nums1: []int{-1, 0, 0, 3, 3, 3, 0, 0, 0},
			m:     6,
			nums2: []int{1, 2, 2},
			n:     3,
			want:  []int{-1, 0, 0, 1, 2, 2, 3, 3, 3}},
		{name: "7",
			nums1: []int{0, 1, 1, 2, 2, 0, 0, 0},
			m:     5,
			nums2: []int{0, 3, 3},
			n:     3,
			want:  []int{0, 0, 1, 1, 2, 2, 3, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			if diff := cmp.Diff(tt.want, tt.nums1); diff != "" {
				t.Errorf("merge()\n%s", diff)
			}
		})
	}
}
