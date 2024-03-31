// Package roman_to_integer
//
// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M. (numbers)
//
// For example, 2 is written as II in Roman numeral, just two ones added together.
// 12 is written as XII, which is simply X + II.
// The number 27 is written as XXVII, which is XX + V + II.
//
// Roman numerals are usually written largest to smallest from left to right.
// However, the numeral for four is not IIII.
// Instead, the number four is written as IV.
// Because the one is before the five we subtract it making four.
// The same principle applies to the number nine, which is written as IX.
// There are six instances where subtraction is used:
//
// - I can be placed before V (5) and X (10) to make 4 and 9.
//
// - X can be placed before L (50) and C (100) to make 40 and 90.
//
// - C can be placed before D (500) and M (1000) to make 400 and 900.
//
// Given a roman numeral, convert it to an integer.
package roman_to_integer

var numbers = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// romanToInt
func romanToInt(s string) (result int) {
	var prev int
	var cur int
	var inc bool
	for s != "" {
		r := s[len(s)-1:]
		s = s[:len(s)-1]
		cur = numbers[r]

		if prev < cur {
			inc = true
		} else if prev > cur {
			inc = false
		}

		if inc {
			result += cur
		} else {
			result -= cur
		}
		prev = cur
	}
	return result
}
