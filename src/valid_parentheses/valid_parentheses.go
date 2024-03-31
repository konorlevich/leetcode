// Package valid_parentheses
//
// Given a string `s` containing just the characters in brackets, determine if the input string is valid.
//
// An input string is valid if:
//
// - Open `brackets` must be closed by the same type of `brackets`.
//
// - Open `brackets` must be closed in the correct order.
//
// - Every close bracket has a corresponding open bracket of the same type.
package valid_parentheses

var (
	brackets = map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
)

// isValid
func byOpen(s string) bool {
	opened := make([]rune, 0, len(s)/2)
	for _, r := range s {
		if _, ok := brackets[r]; ok {
			opened = append(opened, r)
			continue
		}
		if len(opened) == 0 {
			return false
		}
		if r != brackets[opened[len(opened)-1]] {
			return false
		}
		opened = append([]rune{}, opened[:len(opened)-1]...)
	}
	return len(opened) == 0
}

func byClose(s string) bool {
	var closed []rune
	for _, r := range s {
		if c, ok := brackets[r]; ok {
			closed = append(closed, c)
			continue
		}
		if len(closed) == 0 {
			return false
		}
		if r != closed[len(closed)-1] {
			return false
		}
		closed = append([]rune{}, closed[:len(closed)-1]...)
	}
	return len(closed) == 0
}
