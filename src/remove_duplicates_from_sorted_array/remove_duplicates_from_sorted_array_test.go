package remove_duplicates_from_sorted_array

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name     string
		nums     []int
		wantNums []int
		want     int
	}{
		{name: "empty"},
		{name: "1",
			nums:     []int{1, 1, 2},
			wantNums: []int{1, 2},
			want:     2,
		},
		{name: "2",
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			wantNums: []int{0, 1, 2, 3, 4},
			want:     5,
		},
		{name: "3",
			nums:     []int{-1, 0, 0, 0, 0, 3, 3},
			wantNums: []int{-1, 0, 3},
			want:     3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := tt.nums
			got := removeDuplicates(nums)
			if got != tt.want {
				t.Errorf("wrong count = %d, want %d", got, tt.want)
			}
			if diff := cmp.Diff(tt.wantNums, nums[:max(0, got)]); diff != "" {
				t.Errorf("wrong nums: \n%s", diff)
			}
		})
	}
}
