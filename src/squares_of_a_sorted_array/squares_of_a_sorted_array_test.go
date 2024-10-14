package squares_of_a_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortedSquares(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "[-4,-1,0,3,10]",
			nums: []int{-4, -1, 0, 3, 10}, want: []int{0, 1, 9, 16, 100}},
		{name: "[-7,-3,2,3,11]",
			nums: []int{-7, -3, 2, 3, 11}, want: []int{4, 9, 9, 49, 121}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, sortedSquares(tt.nums))
		})
	}
}
