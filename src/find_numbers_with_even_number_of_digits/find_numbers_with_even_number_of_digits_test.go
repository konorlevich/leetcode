package find_numbers_with_even_number_of_digits

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
)

var funcs = map[string]func([]int) int{
	"calculate":         calculateNumbers,
	"convert to string": convertNumbers,
	"compare":           compareNumbers,
}

func Test_findNumbers(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "[12,345,2,6,7896]",
			nums: []int{12, 345, 2, 6, 7896}, want: 2},
		{name: "[555,901,482,1771]",
			nums: []int{555, 901, 482, 1771}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for s, f := range funcs {
				t.Run(s, func(t *testing.T) {
					assert.Equal(t, tt.want, f(tt.nums))
				})
			}
		})
	}
}

func Benchmark_findNumbers(b *testing.B) {
	nums := make([]int, 0, 5000000)
	for i := 0; i < 5000000; i++ {
		nums = append(nums, rand.Int())
	}
	for s, f := range funcs {
		b.Run(s, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
					f(nums)
				})
			}
		})
	}
}
