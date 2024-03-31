package excel_sheet_column_number

import (
	"testing"
)

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
