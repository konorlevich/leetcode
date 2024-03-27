package remove_element

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_removeElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		wantNums []int
		want     int
	}{
		{name: "empty"},
		{name: "1",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			wantNums: []int{2, 2},
			want:     2,
		},
		{name: "2",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			wantNums: []int{0, 1, 3, 0, 4},
			want:     5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := removeElement(tt.nums, tt.val)
			if got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
			if diff := cmp.Diff(tt.wantNums, tt.nums[:got]); diff != "" {
				t.Errorf("removeElement()\n%s", diff)
			}
		})
	}
}
