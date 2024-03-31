// Package add_binary
//
// Given two binary strings a and b, return their sum as a binary string.
package add_binary

func addBinary(a string, b string) string {
	lastA := len(a) - 1
	lastB := len(b) - 1
	// if one of strings empty, return the other one or empty string
	if lastA < 0 {
		return max(b, "")
	}
	if lastB < 0 {
		return max(a, "")
	}

	if a[lastA] == '0' && b[lastB] == '0' { // two zeros gives us a zero
		return addBinary(a[:lastA], b[:lastB]) + "0"
	} else if a[lastA] != b[lastB] { // 1 and 0 gives us 1
		return addBinary(a[:lastA], b[:lastB]) + "1"
	}
	// two 1s gives us 0 and 1 for the next level
	return addBinary(a[:lastA], addBinary(b[:lastB], "1")) + "0"
}
