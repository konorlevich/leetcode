package square_root

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	var left = 0
	var right = x

	for left < right {
		mid := left + ((right - left) / 2)
		ii := mid * mid
		if ii == x {
			return mid
		} else if ii < x {
			left = mid + 1
			continue
		}
		right = mid
	}
	return left - 1
}
