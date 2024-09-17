package string_to_integer_atoi

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myAtoi(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		wantN int
	}{
		{name: "empty string",
			s: "", wantN: 0},
		{name: "42",
			s: "42", wantN: 42},
		{name: "-042",
			s: "-042", wantN: -42},
		{name: "1337c0d3",
			s: "1337c0d3", wantN: 1337},
		{name: "3.14159",
			s: "3.14159", wantN: 3},
		{name: "0-1",
			s: "0-1", wantN: 0},
		{name: ".1",
			s: ".1", wantN: 0},
		{name: "words and 987",
			s: "words and 987", wantN: 0},
		{name: "more than max int32",
			s: fmt.Sprintf("%d", math.MaxInt32+1), wantN: math.MaxInt32},
		{name: "less than min int32",
			s: fmt.Sprintf("%d", math.MinInt32-1), wantN: math.MinInt32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantN, myAtoi(tt.s))
		})
	}
}
