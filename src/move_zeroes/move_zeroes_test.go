package move_zeroes

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var funcs = map[string]func([]int){
	"two pointers": twoPointers,
}

func Test_moveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "[0,1,0,3,12]",
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},

		{name: "[0,0,0,1,0,3,12]",
			nums: []int{0, 0, 0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0, 0, 0},
		},

		{name: "[0]",
			nums: []int{0},
			want: []int{0},
		},

		{name: "[2,1]",
			nums: []int{2, 1},
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for s, f := range funcs {
				t.Run(s, func(t *testing.T) {
					nums := make([]int, len(tt.nums))
					copy(nums, tt.nums)
					f(nums)
					if diff := cmp.Diff(tt.want, nums); diff != "" {
						t.Errorf("%s()\n%s", s, diff)
					}
				})
			}
		})
	}
}
