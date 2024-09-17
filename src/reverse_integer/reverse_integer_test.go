package reverse_integer

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{name: "0", x: 0, want: 0},
		//{name: "min", x: math.MinInt32, want: -8463847412},
		{name: "min", x: math.MinInt32, want: 0},
		//{name: "max", x: math.MaxInt32, want: 7463847412},
		{name: "max", x: math.MaxInt32, want: 0},
		{name: "-123", x: -123, want: -321},
		{name: "-321", x: -321, want: -123},
		{name: "120", x: 120, want: 21},
		{name: "1534236469", x: 1534236469, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverse(tt.x))
		})
	}
}
