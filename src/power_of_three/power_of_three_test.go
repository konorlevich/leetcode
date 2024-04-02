package power_of_three

import "testing"

var funcs = map[string]func(int) bool{
	"recursive": recursive,
	"iterative": iterative,
	"math":      math,
}

func Test_recursive(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{

		{name: "empty"},

		{name: "3^3",
			n:    27,
			want: true},

		{name: "3^1",
			n:    3,
			want: true},

		{name: "3^0",
			n:    1,
			want: true},

		{name: "3^0",
			n:    1,
			want: true},

		{name: "25",
			n:    25,
			want: false},

		{name: "-1",
			n:    -1,
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			for s, f := range funcs {
				t.Run(s, func(t *testing.T) {
					t.Parallel()
					if got := f(tt.n); got != tt.want {
						t.Errorf("%s() = %v, want %v", s, got, tt.want)
					}
				})
			}
		})
	}
}
