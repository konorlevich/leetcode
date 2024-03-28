package missing_ranges

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_findMissingRanges(t *testing.T) {
	tests := []struct {
		name string

		nums  []int
		lower int
		upper int

		wantMissing [][]int
	}{
		{name: "empty",
			wantMissing: make([][]int, 0)},

		{name: "[], 1, 1",
			lower:       1,
			upper:       1,
			wantMissing: [][]int{{1, 1}}},

		{name: "[1,3,50,75], 0, 99",
			nums:  []int{1, 3, 50, 75},
			lower: 0,
			upper: 99,
			wantMissing: [][]int{
				{0, 0},
				{2, 2},
				{4, 49},
				{51, 74},
				{76, 99}}},

		{name: "[0,1,3,50,75], 0, 99",
			nums:  []int{0, 1, 3, 50, 75},
			lower: 0,
			upper: 99,
			wantMissing: [][]int{
				{2, 2},
				{4, 49},
				{51, 74},
				{76, 99}}},

		{name: "[-1], -1, -1",
			nums:        []int{-1},
			lower:       -1,
			upper:       -1,
			wantMissing: [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.wantMissing, findMissingRanges(tt.nums, tt.lower, tt.upper)); diff != "" {
				t.Errorf("findMissingRanges()\n%s", diff)
			}
		})
	}
}
