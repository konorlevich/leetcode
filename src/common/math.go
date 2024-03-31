package common

// Pow returns x to the power of y
func Pow(x, y int) (powed int) {
	if x == 1 || y == 0 {
		return 1
	}
	powed = x
	if y == 1 {
		return powed
	}
	for i := 1; i < y; i++ {
		powed *= x
	}
	return
}
