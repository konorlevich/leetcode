package remove_duplicates_from_sorted_list

import (
	"github.com/google/go-cmp/cmp"
	"github.com/konorlevich/leetcode/src/common/list"
	"testing"
)

var funcs = map[string]func(head *list.Node) *list.Node{
	"iterative": iterative,
	"recursive": recursive,
}

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct {
		name string
		head *list.Node
		want *list.Node
	}{
		{name: "empty"},
		{name: "1,1,2",
			head: list.Create([]int{1, 1, 2}),
			want: list.Create([]int{1, 2}),
		},
		{name: "1,1,2,3,3",
			head: list.Create([]int{1, 1, 2, 3, 3}),
			want: list.Create([]int{1, 2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			for s, f := range funcs {
				t.Run(s, func(t *testing.T) {
					t.Parallel()
					if diff := cmp.Diff(tt.want, f(tt.head)); diff != "" {
						t.Errorf("%s()\n%s", s, diff)
					}
				})
			}

		})
	}
}

func Benchmark_deleteDuplicates(b *testing.B) {
	for s, f := range funcs {
		b.Run(s, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(
					list.Create(
						[]int{
							-96, -86, -85, -81, -75, -75, -74, -72, -63, -63,
							-59, -57, -49, -49, -43, -41, -28, -27, -25, -25,
							-25, -17, -15, -15, -15, -14, -9, -7, -4, -1,
							2, 9, 22, 39, 42, 51, 59, 69, 76, 81,
							85, 86, 86, 90, 93,
						}))
			}
		})
	}
}
