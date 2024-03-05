package palindrome

import "fmt"

func isPalindrome(f func(x int) bool, x int) bool {
	return f(x)
}

func asString(x int) bool {
	if x == 0 {
		return true
	}
	s := fmt.Sprint(x)
	halfL := len(s) / 2
	str1 := s[:halfL]
	var buf string
	var str2 string

	if (len(s) % 2) == 0 {
		buf = s[halfL:]
	} else {
		buf = s[halfL+1:]
	}
	for _, r := range buf {
		str2 = fmt.Sprintf("%c", r) + str2
	}

	return str1 == str2
}

func asNumber(x int) bool {
	buf := x
	reversed := 0
	for buf > 0 {
		reversed = (reversed * 10) + buf%10
		buf = buf / 10
	}
	return reversed == x
}
