package majority_element

import (
	"testing"
)

var funcs = map[string]func(nums []int) int{
	"iterate with a map":              iterateWithMap,
	"count votes":                     countVotes,
	"iterate with a map, check count": iterateWithMapTillMoreThanHalf,
}

func Test_majorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{

		{name: "[1]",
			nums: []int{1},
			want: 1,
		},

		{name: "[3,2,3]",
			nums: []int{3, 2, 3},
			want: 3,
		},

		{name: "[2,2,1,1,1,2,2]",
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			for s, f := range funcs {
				t.Run(s, func(t *testing.T) {
					t.Parallel()
					f := f
					if got := f(tt.nums); got != tt.want {
						t.Errorf("%s() = %v, want %v", s, got, tt.want)
					}
				})
			}

		})
	}
}

func Benchmark_majorityElement(b *testing.B) {
	for s, f := range funcs {
		b.Run(s, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f([]int{2, 2, 1, 1, 1, 2, 2})
			}
		})
	}
}
