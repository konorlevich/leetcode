package pascal_s_triangle

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

var funcs map[string]func(int) [][]int = map[string]func(int) [][]int{
	"brute force": generateBruteForce,
	//"alternative": generateAlternative,
}

func Test_generate(t *testing.T) {
	tests := []struct {
		name    string
		numRows int
		want    [][]int
	}{
		{name: "empty",
			want: [][]int{},
		},
		{name: "1",
			numRows: 1,
			want:    [][]int{{1}},
		},
		{name: "5",
			numRows: 5,
			want: [][]int{
				{1},
				{1, 1}, {1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1}},
		},
		{name: "6",
			numRows: 6,
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
				{1, 5, 10, 10, 5, 1}},
		},
		{name: "7",
			numRows: 7,
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
				{1, 5, 10, 10, 5, 1},
				{1, 6, 15, 20, 15, 6, 1}},
		},
		{name: "8",
			numRows: 8,
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
				{1, 5, 10, 10, 5, 1},
				{1, 6, 15, 20, 15, 6, 1},
				{1, 7, 21, 35, 35, 21, 7, 1},
			},
		},
		{name: "9",
			numRows: 9,
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
				{1, 5, 10, 10, 5, 1},
				{1, 6, 15, 20, 15, 6, 1},
				{1, 7, 21, 35, 35, 21, 7, 1},
				{1, 8, 28, 56, 70, 56, 28, 8, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			for s, f := range funcs {
				t.Run(s, func(t *testing.T) {
					if diff := cmp.Diff(tt.want, f(tt.numRows)); diff != "" {
						t.Errorf("generate()\n%s", diff)
					}
				})
			}
		})
	}
}
