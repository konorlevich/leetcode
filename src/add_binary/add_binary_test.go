package add_binary

import "testing"

func Test_addBinary(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		//by the task description, `1 <= a.length, b.length <= 10^4`,
		//so we can't test with an empty string
		//{name: "empty"},
		{name: "0+0",
			a:    "0",
			b:    "0",
			want: "0"},
		{name: "11+1",
			a: "11", b: "1",
			want: "100"},
		{name: "1010+1011",
			a: "1010", b: "1011",
			want: "10101"},
		{name: "0+1",
			a: "0", b: "1",
			want: "1"},
		{name: "1+111",
			a: "1", b: "111",
			want: "1000"},
		{name: "100+110010",
			a: "100", b: "110010",
			want: "110110"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := addBinary(tt.a, tt.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
