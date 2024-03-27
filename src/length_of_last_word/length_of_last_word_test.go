package length_of_last_word

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		//by the task description, `There will be at least one word in s.`,
		//so we can't test with an empty string
		//{name: "empty"},
		{name: "Hello World",
			s: "Hello World", want: 5},
		{name: "   fly me   to   the moon  ",
			s: "   fly me   to   the moon  ", want: 4},
		{name: "luffy is still joyboy",
			s: "luffy is still joyboy", want: 6},
		{name: " a",
			s: " a", want: 1},
		{name: "a",
			s: "a", want: 1},
		{name: " aa",
			s: " aa", want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := lengthOfLastWord(tt.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
