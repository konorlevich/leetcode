package reverse_linked_list

import (
	"github.com/google/go-cmp/cmp"
	"github.com/konorlevich/leetcode/src/common/list"
	"testing"
)

var funcs = map[string]func(*list.Node) *list.Node{
	"iterative": iterative,
	"recursive": recursive,
}

func Test_reverseList(t *testing.T) {
	tests := []struct {
		name string
		head *list.Node
		want *list.Node
	}{
		{name: "empty"},

		{name: "1,2,3,4,5 => 5,4,3,2,1",
			head: list.Create([]int{1, 2, 3, 4, 5}),
			want: list.Create([]int{5, 4, 3, 2, 1}),
		},

		{name: "1,2 => 2,1",
			head: list.Create([]int{1, 2}),
			want: list.Create([]int{2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for s, f := range funcs {
				t.Run(s, func(t *testing.T) {
					if diff := cmp.Diff(tt.want, f(tt.head)); diff != "" {
						t.Errorf("%s()\n%s", s, diff)
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
				f(list.Create([]int{
					1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
					1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
					1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
					1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
					1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
					1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
					1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
					1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
					1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
				}))
			}
		})

	}
}
