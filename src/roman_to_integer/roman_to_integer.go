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
