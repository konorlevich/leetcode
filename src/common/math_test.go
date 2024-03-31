package common

import "testing"

func Test_pow(t *testing.T) {
	tests := []struct {
		name string

		x int
		y int

		wantPowed int
	}{
		{name: "empty", wantPowed: 1},
		{name: "1^10",
			x:         1,
			y:         10,
			wantPowed: 1},
		{name: "2^0",
			x:         2,
			y:         0,
			wantPowed: 1},

		{name: "2^2",
			x:         2,
			y:         2,
			wantPowed: 4},

		{name: "2^20",
			x:         2,
			y:         20,
			wantPowed: 1048576},

		{name: "7^10",
			x:         7,
			y:         10,
			wantPowed: 282475249},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPowed := Pow(tt.x, tt.y); gotPowed != tt.wantPowed {
				t.Errorf("pow() = %v, want %v", gotPowed, tt.wantPowed)
			}
		})
	}
}
