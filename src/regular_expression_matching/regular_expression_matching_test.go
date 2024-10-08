package regular_expression_matching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isMatch(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want bool
	}{
		{
			name: "aa with a",
			s:    "aa", p: "a",
			want: false,
		},
		{
			name: "aa with a*",
			s:    "aa", p: "a*",
			want: true,
		},
		{
			name: "ab with .*",
			s:    "ab", p: ".*",
			want: true,
		},
		{
			name: "ab with ab",
			s:    "ab", p: "ab",
			want: true,
		},
		{
			name: "ab with .*ab",
			s:    "abab", p: ".*ab",
			want: true,
		},
		{
			name: "Hello World with Hel*o .*",
			s:    "Hello World", p: "Hel*o .*",
			want: true,
		},
		{
			name: "mississippi with mis*is*p*.",
			s:    "mississippi", p: "mis*is*p*.",
			want: false,
		},
		{
			name: "aaa with ab*a",
			s:    "aaa", p: "ab*a",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isMatch(tt.s, tt.p))
		})
	}
}
