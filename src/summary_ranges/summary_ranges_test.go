package summary_ranges

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_summaryRanges(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []string
	}{
		{name: "nil", nums: nil, want: []string{}},
		{name: "empty", nums: []int{}, want: []string{}},
		{name: "one", nums: []int{1}, want: []string{"1"}},
		{name: "two", nums: []int{2}, want: []string{"2"}},
		{name: "example 1", nums: []int{0, 1, 2, 4, 5, 7}, want: []string{"0->2", "4->5", "7"}},
		{name: "example2", nums: []int{0, 2, 3, 4, 6, 8, 9}, want: []string{"0", "2->4", "6", "8->9"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, summaryRanges(tt.nums))
		})
	}
}

func Benchmark_summaryRanges(b *testing.B) {
	count := 100_000_000
	data := make([]int, 0, count)
	for i := 0; i < count; i++ {
		if i%10 == 0 {
			continue
		}
		data = append(data, i)
	}

	b.Run(fmt.Sprintf("%d", count), func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			summaryRanges(data)
		}
	})
}
