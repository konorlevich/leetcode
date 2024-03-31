// Package climbing_stairs
//
// You are climbing a staircase. It takes n steps to reach the top.
//
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
package climbing_stairs

// recursive
// # The solution, in general looks like
//
// T(0) = 0, T(1) = 1, T(2) = 2, T(3) = 3, T(4) = 5
//
// T(n) = T(n-2) + T(n-1)
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
