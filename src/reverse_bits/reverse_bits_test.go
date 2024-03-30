package reverse_bits

import (
	"fmt"
	"testing"
)

var funcs = map[string]func(num uint32) uint32{
	"iterative":                  iterative,
	"iterative with right shift": iterativeWithRightShift,
}

func Test_reverseBits(t *testing.T) {
	tests := []struct {
		name string
		num  uint32
		want uint32
	}{
		{name: "empty", want: 0},

		{name: "43261596 to 964176192",
			num:  43261596,  //00000010100101000001111010011100
			want: 964176192, //00111001011110000010100101000000
		},

		{name: "43261596 to 964176192",
			num:  4294967293, //11111111111111111111111111111101
			want: 3221225471, //10111111111111111111111111111111
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for s, f := range funcs {
				t.Run(s, func(t *testing.T) {
					if got := f(tt.num); got != tt.want {
						t.Errorf("%s() = %v, want %v", s, got, tt.want)
					}
				})
			}
		})
	}
}

func Benchmark_reverseBits(b *testing.B) {
	cases := []uint32{
		964176192, 0,
	}
	for s, f := range funcs {
		for _, num := range cases {
			b.Run(fmt.Sprintf("%s - %b", s, num), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f(num)
				}
			})
		}

	}
}
