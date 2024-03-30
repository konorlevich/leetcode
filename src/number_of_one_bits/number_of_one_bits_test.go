package number_of_one_bits

import "testing"

func Test_hammingWeight(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "empty"},

		{name: "1011",
			n:    11,
			want: 3,
		},

		{name: "10000000",
			n:    128,
			want: 1,
		},

		{name: "1111111111111111111111111111101",
			n:    2147483645,
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := iterativeTill0(tt.n); got != tt.want {
				t.Errorf("iterativeTill0() = %v, want %v", got, tt.want)
			}
		})
	}
}
