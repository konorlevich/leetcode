package remove_duplicates_from_sorted_list

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/konorlevich/leetcode/src/common"
	"testing"
)

var funcs = map[string]func(head *common.ListNode) *common.ListNode{
	"iterative": iterative,
	"recursive": recursive,
}

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct {
		name string
		head *common.ListNode
		want *common.ListNode
	}{
		{name: "empty"},
		{name: "1,1,2",
			head: common.CreateList([]int{1, 1, 2}),
			want: common.CreateList([]int{1, 2}),
		},
		{name: "1,1,2,3,3",
			head: common.CreateList([]int{1, 1, 2, 3, 3}),
			want: common.CreateList([]int{1, 2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for s, f := range funcs {
				t.Run(s, func(t *testing.T) {
					if diff := cmp.Diff(tt.want, deleteDuplicates(f, tt.head)); diff != "" {
						t.Errorf("%s()\n%s", s, diff)
					}
				})
			}

		})
	}
}

func Benchmark_deleteDuplicates(b *testing.B) {
	cases := map[int][]int{
		10: {1, 1, 2, 2, 3, 4, 5, 5, 6, 7},
		25: {
			1, 1, 2, 2, 3, 4, 5, 5, 6, 7,
			8, 8, 9, 10, 11, 11, 12, 12, 12, 13,
			13, 14, 15, 16, 16,
		},
		45: {
			-96, -86, -85, -81, -75, -75, -74, -72, -63, -63,
			-59, -57, -49, -49, -43, -41, -28, -27, -25, -25,
			-25, -17, -15, -15, -15, -14, -9, -7, -4, -1,
			2, 9, 22, 39, 42, 51, 59, 69, 76, 81,
			85, 86, 86, 90, 93,
		},
	}
	for target, ints := range cases {
		list := common.CreateList(ints)
		b.Run(fmt.Sprintf("%d", target), func(b *testing.B) {
			for s, f := range funcs {
				b.Run(s, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						deleteDuplicates(f, list)
					}
				})
			}
		})
	}
}
