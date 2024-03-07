package square_root

import "testing"

func Test_mySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{name: "empty"},
		{name: "1", x: 1, want: 1},
		{name: "2", x: 2, want: 1},
		{name: "3", x: 3, want: 1},
		{name: "4", x: 4, want: 2},
		{name: "5", x: 5, want: 2},
		{name: "6", x: 6, want: 2},
		{name: "7", x: 7, want: 2},
		{name: "8", x: 8, want: 2},
		{name: "9", x: 9, want: 3},
		{name: "2147395599", x: 2147395599, want: 46339},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
