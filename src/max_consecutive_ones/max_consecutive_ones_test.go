package max_consecutive_ones

import "testing"

func Test_findMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "0",
			nums: []int{0}, want: 0},
		{name: "1",
			nums: []int{1}, want: 1},
		{
			name: "[1,0,1,0,1,0,1]",
			nums: []int{1, 0, 1, 0, 1, 0, 1}, want: 1,
		},
		{
			name: "[1,1,0,1,1,1]",
			nums: []int{1, 1, 0, 1, 1, 1}, want: 3,
		},
		{
			name: "[1,0,1,1,0,1]",
			nums: []int{1, 0, 1, 1, 0, 1}, want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxConsecutiveOnes(tt.nums); got != tt.want {
				t.Errorf("findMaxConsecutiveOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
