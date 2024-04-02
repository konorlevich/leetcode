package missing_number

import "testing"

var funcs = map[string]func([]int) int{
	"with temp slice": withTempSlice,
	"brute force":     bruteForce,
	"sum":             sumIt,
}

func Test_missingNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "[3,0,1]",
			nums: []int{3, 0, 1},
			want: 2,
		},

		{name: "[0,1]",
			nums: []int{0, 1},
			want: 2,
		},

		{name: "[9,6,4,2,3,5,7,0,1]",
			nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			want: 8,
		},

		{name: "[0,1,2,3,4,5,6,7,8,9]",
			nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 10,
		},
		{name: "[1,2,3,4,5,6,7,8,9]",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			for s, f := range funcs {
				t.Run(s, func(t *testing.T) {
					t.Parallel()
					//some functions may change the slice
					// it's not great, but it's not forbidden
					nums := make([]int, len(tt.nums))
					copy(nums, tt.nums)
					if got := f(nums); got != tt.want {
						t.Errorf("%s\n %v, want %v", s, got, tt.want)
					}
				})
			}

		})
	}
}
