package contains_duplicate

import (
	"testing"
)

var funcs = map[string]func([]int) bool{
	"n^2":    inefficient,
	"linear": withMap,
}

func Test_containsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{name: "empty"},

		{name: "[1,2,3,1]",
			nums: []int{1, 2, 3, 1},
			want: true},

		{name: "[1,2,3,4]",
			nums: []int{1, 2, 3, 4},
			want: false},

		{name: "[1,1,1,3,3,4,3,2,4,2]",
			nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for s, f := range funcs {
				t.Run(s, func(t *testing.T) {
					if got := f(tt.nums); got != tt.want {
						t.Errorf("%s() = %v, want %v", s, got, tt.want)
					}
				})
			}

		})
	}
}

func Benchmark_reverseBits(b *testing.B) {
	for s, f := range funcs {
		b.Run(s, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f([]int{
					1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
					11, 12, 13, 14, 15, 16, 17, 18, 19, 10,
					21, 22, 23, 24, 25, 26, 27, 28, 29, 20,
					31, 32, 33, 34, 35, 36, 37, 38, 39, 30,
					41, 42, 43, 44, 45, 46, 47, 48, 49, 40,
				})
			}
		})
	}
}
