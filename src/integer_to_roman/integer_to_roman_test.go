package integer_to_roman

import "testing"

func Test_intToRoman(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{name: "1",
			num: 1, want: "I"},
		{name: "2",
			num: 2, want: "II"},
		{name: "3749",
			num: 3749, want: "MMMDCCXLIX"},
		{name: "58",
			num: 58, want: "LVIII"},
		{name: "1994",
			num: 1994, want: "MCMXCIV"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.num); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
