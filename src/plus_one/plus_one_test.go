package plus_one

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_plusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		//by the task description, `1 <= digits.length <= 100`,
		//so we can't test with an empty string
		//{name: "empty"},
		{name: "1",
			digits: []int{1, 2, 3},
			want:   []int{1, 2, 4}},
		{name: "2",
			digits: []int{4, 3, 2, 1},
			want:   []int{4, 3, 2, 2}},
		{name: "3",
			digits: []int{9},
			want:   []int{1, 0}},
		{name: "4",
			digits: []int{0},
			want:   []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if diff := cmp.Diff(tt.want, plusOne(tt.digits)); diff != "" {
				t.Errorf("plusOne()\n%s", diff)
			}
		})
	}
}
