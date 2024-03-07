package search_insert_position

import "testing"

func Test_searchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{name: "empty"},
		{name: "1",
			nums: []int{1, 3, 5, 6}, target: 5, want: 2},
		{name: "2",
			nums: []int{1, 3, 5, 6}, target: 2, want: 1},
		{name: "3",
			nums: []int{1, 3, 5, 6}, target: 7, want: 4},
		{name: "4",
			nums: []int{1}, target: 1, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.nums, tt.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
