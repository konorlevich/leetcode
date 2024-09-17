package valid_palindrome

import (
	"testing"

	"github.com/konorlevich/leetcode/src/common/ascii"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "empty", want: true},

		{name: "A man, a plan, a canal: Panama",
			s:    "A man, a plan, a canal: Panama",
			want: true},

		{name: "race a car",
			s:    "race a car",
			want: false},

		{name: "one space",
			s:    " ",
			want: true},

		{name: "0P",
			s:    "0P",
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAlphanumeric(t *testing.T) {
	tests := []struct {
		name string
		s    uint8
		want bool
	}{
		{name: "empty", want: false},
		{name: "NUL (null)", s: uint8(0), want: false},
		{name: "SOH (start of heading)", s: uint8(1), want: false},
		{name: "STX (start of text)", s: uint8(2), want: false},
		{name: "ETX (end of text)", s: uint8(3), want: false},
		{name: "EOT (end of transmission)", s: uint8(4), want: false},
		{name: "ENQ (enquiry)", s: uint8(5), want: false},
		{name: "ACK (acknowledge)", s: uint8(6), want: false},
		{name: "BEL (bell)", s: uint8(7), want: false},
		{name: "BS  (backspace)", s: uint8(8), want: false},
		{name: "TAB (horizontal tab)", s: uint8(9), want: false},
		{name: "LF  (NL line feed, new line)", s: uint8(10), want: false},
		{name: "VT  (vertical tab)", s: uint8(11), want: false},
		{name: "FF  (NP form feed, new page)", s: uint8(12), want: false},
		{name: "CR  (carriage return)", s: uint8(13), want: false},
		{name: "SO  (shift out)", s: uint8(14), want: false},
		{name: "SI  (shift in)", s: uint8(15), want: false},
		{name: "DLE (data link escape)", s: uint8(16), want: false},
		{name: "DC1 (device control 1)", s: uint8(17), want: false},
		{name: "DC2 (device control 2)", s: uint8(18), want: false},
		{name: "DC3 (device control 3)", s: uint8(19), want: false},
		{name: "DC4 (device control 4)", s: uint8(20), want: false},
		{name: "NAK (negative acknowledge)", s: uint8(21), want: false},
		{name: "SYN (synchronous idle)", s: uint8(22), want: false},
		{name: "ETB (end of trans. block)", s: uint8(23), want: false},
		{name: "CAN (cancel)", s: uint8(24), want: false},
		{name: "EM  (end of medium)", s: uint8(25), want: false},
		{name: "SUB (substitute)", s: uint8(26), want: false},
		{name: "ESC (escape)", s: uint8(27), want: false},
		{name: "FS  (file separator)", s: uint8(28), want: false},
		{name: "GS  (group separator)", s: uint8(29), want: false},
		{name: "RS  (record separator)", s: uint8(30), want: false},
		{name: "US  (unit separator)", s: uint8(31), want: false},
		{name: "SPACE", s: uint8(32), want: false},
		{name: "!", s: uint8(33), want: false},
		{name: "\"", s: uint8(34), want: false},
		{name: "#", s: uint8(35), want: false},
		{name: "$", s: uint8(36), want: false},
		{name: "%", s: uint8(37), want: false},
		{name: "&", s: uint8(38), want: false},
		{name: "'", s: uint8(39), want: false},
		{name: "(", s: uint8(40), want: false},
		{name: ")", s: uint8(41), want: false},
		{name: "*", s: uint8(42), want: false},
		{name: "+", s: uint8(43), want: false},
		{name: ",", s: uint8(44), want: false},
		{name: "-", s: uint8(45), want: false},
		{name: ".", s: uint8(46), want: false},
		{name: "/", s: uint8(47), want: false},
		{name: "0", s: uint8(48), want: true},
		{name: "1", s: uint8(49), want: true},
		{name: "2", s: uint8(50), want: true},
		{name: "3", s: uint8(51), want: true},
		{name: "4", s: uint8(52), want: true},
		{name: "5", s: uint8(53), want: true},
		{name: "6", s: uint8(54), want: true},
		{name: "7", s: uint8(55), want: true},
		{name: "8", s: uint8(56), want: true},
		{name: "9", s: uint8(57), want: true},
		{name: ":", s: uint8(58), want: false},
		{name: ";", s: uint8(59), want: false},
		{name: "<", s: uint8(60), want: false},
		{name: "=", s: uint8(61), want: false},
		{name: ">", s: uint8(62), want: false},
		{name: "?", s: uint8(63), want: false},
		{name: "@", s: uint8(64), want: false},
		{name: "A", s: uint8(65), want: true},
		{name: "B", s: uint8(66), want: true},
		{name: "C", s: uint8(67), want: true},
		{name: "D", s: uint8(68), want: true},
		{name: "E", s: uint8(69), want: true},
		{name: "F", s: uint8(70), want: true},
		{name: "G", s: uint8(71), want: true},
		{name: "H", s: uint8(72), want: true},
		{name: "I", s: uint8(73), want: true},
		{name: "J", s: uint8(74), want: true},
		{name: "K", s: uint8(75), want: true},
		{name: "L", s: uint8(76), want: true},
		{name: "M", s: uint8(77), want: true},
		{name: "N", s: uint8(78), want: true},
		{name: "O", s: uint8(79), want: true},
		{name: "P", s: uint8(80), want: true},
		{name: "Q", s: uint8(81), want: true},
		{name: "R", s: uint8(82), want: true},
		{name: "S", s: uint8(83), want: true},
		{name: "T", s: uint8(84), want: true},
		{name: "U", s: uint8(85), want: true},
		{name: "V", s: uint8(86), want: true},
		{name: "W", s: uint8(87), want: true},
		{name: "X", s: uint8(88), want: true},
		{name: "Y", s: uint8(89), want: true},
		{name: "Z", s: uint8(90), want: true},
		{name: "[", s: uint8(91), want: false},
		{name: "\\", s: uint8(92), want: false},
		{name: "]", s: uint8(93), want: false},
		{name: "^", s: uint8(94), want: false},
		{name: "_", s: uint8(95), want: false},
		{name: "`", s: uint8(96), want: false},
		{name: "a", s: uint8(97), want: true},
		{name: "b", s: uint8(98), want: true},
		{name: "c", s: uint8(99), want: true},
		{name: "d", s: uint8(100), want: true},
		{name: "e", s: uint8(101), want: true},
		{name: "f", s: uint8(102), want: true},
		{name: "g", s: uint8(103), want: true},
		{name: "h", s: uint8(104), want: true},
		{name: "i", s: uint8(105), want: true},
		{name: "j", s: uint8(106), want: true},
		{name: "k", s: uint8(107), want: true},
		{name: "l", s: uint8(108), want: true},
		{name: "m", s: uint8(109), want: true},
		{name: "n", s: uint8(110), want: true},
		{name: "o", s: uint8(111), want: true},
		{name: "p", s: uint8(112), want: true},
		{name: "q", s: uint8(113), want: true},
		{name: "r", s: uint8(114), want: true},
		{name: "s", s: uint8(115), want: true},
		{name: "t", s: uint8(116), want: true},
		{name: "u", s: uint8(117), want: true},
		{name: "v", s: uint8(118), want: true},
		{name: "w", s: uint8(119), want: true},
		{name: "x", s: uint8(120), want: true},
		{name: "y", s: uint8(121), want: true},
		{name: "z", s: uint8(122), want: true},
		{name: "{", s: uint8(123), want: false},
		{name: "|", s: uint8(124), want: false},
		{name: "}", s: uint8(125), want: false},
		{name: "~", s: uint8(126), want: false},
		{name: "DEL", s: uint8(127), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isAlphanumeric(tt.s); got != tt.want {
				t.Errorf("isAlphanumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_areSimilar(t *testing.T) {
	tests := []struct {
		name string
		s1   uint8
		s2   uint8
		want bool
	}{
		{name: "empty", want: true},
		{name: "o != p", s1: uint8(111), s2: uint8(112), want: false},
		{name: "a == a", s1: ascii.LetterA, s2: ascii.LetterA, want: true},
		{name: "a == A", s1: ascii.LetterA, s2: ascii.LetterACapital, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := areSimilar(tt.s1, tt.s2); got != tt.want {
				t.Errorf("areSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
