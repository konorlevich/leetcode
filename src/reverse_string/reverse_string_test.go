package reverse_string

import "testing"

func Test_reverseString(t *testing.T) {
	tests := []struct {
		name string
		s    []byte
		want []byte
	}{
		{name: "empty"},
		{name: "Hannah",
			s:    []byte("Hannah"),
			want: []byte("hannaH"),
		},
		{name: "hello",
			s:    []byte("hello"),
			want: []byte("olleh"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.s
			reverseString(s)
		})
	}
}
