package climbing_stairs

import (
	"fmt"
	"testing"
)

var funcs = map[string]func(n int) int{
	"shortOne":  shortOne,
	"iterative": iterative,
	"recursive": recursive,
}

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		//by the task description, `1 <= n <= 45`,
		//so we can't test with an empty string
		//{name: "empty"},
		{name: "2", n: 2, want: 2},
		{name: "3", n: 3, want: 3},
		{name: "4", n: 4, want: 5},
		{name: "5", n: 5, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			for s, f := range funcs {
				t.Run(s, func(t *testing.T) {
					if got := climbStairs(f, tt.n); got != tt.want {
						t.Errorf("%s() = %v, want %v", s, got, tt.want)
					}
				})
			}

		})
	}
}

func Benchmark_climbStairs(b *testing.B) {
	cases := map[int]int{
		10: 10,
		25: 25,
		45: 45,
	}
	for target, ints := range cases {
		b.Run(fmt.Sprintf("%d", target), func(b *testing.B) {
			for s, f := range funcs {
				b.Run(s, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						climbStairs(f, ints)
					}
				})
			}
		})
	}
}
