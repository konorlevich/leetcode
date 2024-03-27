package common

import (
	"reflect"
	"testing"
)

func Test_CreateList(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want *ListNode
	}{
		{name: "empty"},
		{
			name: "valid", nums: []int{1, 2, 3},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := CreateList(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createList() = %v, want %v", got, tt.want)
			}
		})
	}
}
