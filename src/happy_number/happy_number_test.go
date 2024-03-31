package happy_number

import (
	"fmt"
	"testing"
)

var funcs = map[string]func(n int) bool{
	"simple":            simple,
	"simple with a map": simpleWithMap,
}

func Test_isHappy(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{name: "19",
			n:    19,
			want: true,
		},
		{name: "1",
			n:    1,
			want: true,
		},
		{name: "2",
			n:    2,
			want: false},
		{name: "20",
			n:    20,
			want: false},
		{name: "2147483647",
			n:    2147483647,
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for s, f := range funcs {
				t.Run(s, func(t *testing.T) {
					if got := f(tt.n); got != tt.want {
						t.Errorf("%s() = %v, want %v", s, got, tt.want)
					}
				})
			}
		})
	}
}

func Benchmark_isHappy(b *testing.B) {
	cases := []int{
		19, 2147483647, 20,
	}
	for _, num := range cases {
		for s, f := range funcs {
			b.Run(fmt.Sprintf("%s - %d", s, num), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f(num)
				}
			})
		}

	}
}
