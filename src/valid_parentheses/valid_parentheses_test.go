package valid_parentheses

import (
	"fmt"
	"testing"
)

var funcs = map[string]func(s string) bool{
	"byOpen":  byOpen,
	"byClose": byClose,
}

func Test_isValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "empty", want: true},
		{name: "1", s: "()", want: true},
		{name: "2", s: "()[]{}", want: true},
		{name: "3", s: "(]", want: false},
		{name: "4", s: "{([])}", want: true},
		{name: "5", s: "[", want: false},
		{name: "6", s: "]", want: false},
		{name: "7",
			s:    "{[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[()]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]}",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			for s, f := range funcs {
				t.Run(s, func(t *testing.T) {
					t.Parallel()
					if got := f(tt.s); got != tt.want {
						t.Errorf("%s() = %v, want %v", s, got, tt.want)
					}
				})
			}
		})
	}
}

func Benchmark_isValid(b *testing.B) {
	cases := []string{
		"{([])}",
		"{[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[({[({[([{({[({[()]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]})]})]})}])]})]}",
	}
	for i, s := range cases {
		for name, f := range funcs {
			b.Run(name, func(b *testing.B) {
				b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						f(s)
					}
				})
			})

		}
	}
}
