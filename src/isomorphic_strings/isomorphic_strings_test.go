package isomorphic_strings

import "testing"

var funcs = map[string]func(string, string) bool{
	"with a map":   withMap,
	"with a array": withArray,
}

func Test_isIsomorphic(t *testing.T) {
	tests := []struct {
		name string

		s string
		t string

		want bool
	}{
		{name: "egg | egg",
			s:    "egg",
			t:    "egg",
			want: true,
		},
		{name: "egg | add",
			s:    "egg",
			t:    "add",
			want: true,
		},
		{name: "foo | bar",
			s:    "foo",
			t:    "bar",
			want: false,
		},
		{name: "paper | title",
			s:    "paper",
			t:    "title",
			want: true,
		},
		{name: "badc | baba",
			s:    "badc",
			t:    "baba",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for s, f := range funcs {
				t.Run(s, func(t *testing.T) {
					if got := f(tt.s, tt.t); got != tt.want {
						t.Errorf("%s() = %v, want %v", s, got, tt.want)
					}
				})
			}
		})
	}
}
