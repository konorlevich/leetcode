package roman_to_integer

import "testing"

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "empty"},
		{name: "1", s: "I", want: 1},
		{name: "2", s: "II", want: 2},
		{name: "3", s: "III", want: 3},
		{name: "4", s: "IV", want: 4},
		{name: "9", s: "IX", want: 9},
		{name: "58", s: "LVIII", want: 58},
		{name: "1994", s: "MCMXCIV", want: 1994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}

		})
	}
}
