package linked_list_cycle

import (
	"github.com/konorlevich/leetcode/src/common/list"
	"testing"
)

var funcs = map[string]func(head *list.Node) bool{
	"with node buffer":              hasCycleWithNodeBuffer,
	"with slow and fast pointers":   hasCycleWithPointers,
	"with slow and fast pointers 2": hasCycleWithPointers2,
}

func Test_hasCycle(t *testing.T) {
	tests := []struct {
		name string
		head *list.Node
		want bool
	}{

		{name: "head = [3,2,0,-4], pos = 1",
			head: list.CreateWithCycle([]int{3, 2, 0, -4}, 1),
			want: true,
		},

		{name: "head = [3,2,0,-4], pos = -1",
			head: list.CreateWithCycle([]int{3, 2, 0, -4}, -1),
			want: false,
		},

		{name: "head = [1,2], pos = 0",
			head: list.CreateWithCycle([]int{1, 2}, 0),
			want: true,
		},

		{name: "head = [1,2], pos = -1",
			head: list.CreateWithCycle([]int{1, 2}, -1),
			want: false,
		},

		{name: "head = [1], pos = -1",
			head: list.CreateWithCycle([]int{1}, -1),
			want: false,
		},

		{name: "head = [1], pos = 0",
			head: list.CreateWithCycle([]int{1}, 0),
			want: true,
		},

		{name: "head = []",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			for s, f := range funcs {
				t.Run(s, func(t *testing.T) {
					t.Parallel()
					if got := f(tt.head); got != tt.want {
						t.Errorf("%s() = %v, want %v", s, got, tt.want)
					}
				})
			}
		})
	}
}

func Benchmark_hasCycle(b *testing.B) {
	for s, f := range funcs {
		b.Run(s, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(list.CreateWithCycle([]int{
					3, 2, 0, -4, 3,
					2, 0, -4, 3, 2,
					0, -4, 3, 2, 0,
					-4, 3, 2, 0, -4,
					3, 2, 0, -4, 3,
					2, 0, -4, 3, 2,
					0, -4, 3, 2, 0,
					-4,
				}, 35))
			}
		})
	}
}
