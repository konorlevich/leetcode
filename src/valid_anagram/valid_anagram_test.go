package valid_anagram

import "testing"

var funcs = map[string]func(string, string) bool{
	"iterate with a map":               iterateWithMap,
	"sum of letters to the power of 4": sumOfPower,
}

func Test_isAnagram(t *testing.T) {
	tests := []struct {
		name string

		s string
		t string

		want bool
	}{
		{name: "empty",
			want: true},

		{name: "anagram|nagaram",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{name: "anagram|nagara",
			s:    "anagram",
			t:    "nagara",
			want: false,
		},

		{name: "rat|cat",
			s:    "rat",
			t:    "cat",
			want: false,
		},

		{name: "tat|cat",
			s:    "tat",
			t:    "cat",
			want: false,
		},

		{name: "ac|bb",
			s:    "ac",
			t:    "bb",
			want: false,
		},

		{name: "fe|ja",
			s:    "fe",
			t:    "ja",
			want: false,
		},

		{name: "nl|cx",
			s:    "nl",
			t:    "cx",
			want: false,
		},

		{name: "hqbqo|lsnma",
			s:    "hqbqo",
			t:    "lsnma",
			want: false,
		},

		{name: "vbnxkji|wqdtegp",
			s:    "vbnxkji",
			t:    "wqdtegp",
			want: false,
		},

		//unicode strings
		{name: "фрф|рфф",
			s:    "фрф",
			t:    "фрф",
			want: true,
		},
		{name: "фрф|рфр",
			s:    "фрф",
			t:    "рфр",
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

func Benchmark_isAnagram(b *testing.B) {
	for s, f := range funcs {
		b.Run(s, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(
					"asdfgasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdf",
					"asdfgasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdf")
			}
		})
	}
}
