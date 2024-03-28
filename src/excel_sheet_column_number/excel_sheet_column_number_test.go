package excel_sheet_column_number

import "testing"

func Test_titleToNumber(t *testing.T) {
	tests := []struct {
		name        string
		columnTitle string
		want        int
	}{
		{name: "empty"},

		{name: "A",
			columnTitle: "A",
			want:        1},

		{name: "AB",
			columnTitle: "AB",
			want:        28},

		{name: "ZY",
			columnTitle: "ZY",
			want:        701},

		{name: "FXSHRXW",
			columnTitle: "FXSHRXW",
			want:        2147483647},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.columnTitle); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			if gotPowed := pow(tt.x, tt.y); gotPowed != tt.wantPowed {
				t.Errorf("pow() = %v, want %v", gotPowed, tt.wantPowed)
			}
		})
	}
}
