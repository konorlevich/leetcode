package add_binary

func addBinary(a string, b string) string {
	lastA := len(a) - 1
	lastB := len(b) - 1
	if lastA < 0 {
		return max(b, "")
	}
	if lastB < 0 {
		return max(a, "")
	}
	if a[lastA] == '0' && b[lastB] == '0' {
		return addBinary(a[:lastA], b[:lastB]) + "0"
	} else if a[lastA] != b[lastB] {
		return addBinary(a[:lastA], b[:lastB]) + "1"
	}
	return addBinary(a[:lastA], addBinary(b[:lastB], "1")) + "0"
}
