package climbing_stairs

func climbStairs(f func(n int) int, n int) int {
	return f(n)
}

func recursive(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	return recursive(n-1) + recursive(n-2)
}

func iterative(n int) int {
	res := 0

	secondNext := 0
	next := 0
	for i := 1; i <= n; i++ {
		if i == 1 {
			res = 1
		} else if i == 2 {
			res = 2
		} else {
			res = secondNext + next
		}

		next = secondNext
		secondNext = res
	}

	return res
}

func shortOne(n int) int {
	secondLast, last := 0, 1
	for i := 1; i <= n; i++ {
		secondLast, last = last, secondLast+last
	}
	return last
}
