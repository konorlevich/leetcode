// Package container_with_most_water
//
// You are given an integer array height of length n.
// There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
//
// Find two lines that together with the x-axis form a container, such that the container contains the most water.
//
// Return the maximum amount of water a container can store.
//
// Notice that you may not slant the container.
//
// Constraints:
//
// n == height.length
//
// 2 <= n <= 105
//
// 0 <= height[i] <= 104
package container_with_most_water

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxA := (r - l) * (min(height[l], height[r]))
	for r > l {
		if height[r] < height[l] {
			r--
		} else {
			l++
		}

		s := (r - l) * (min(height[l], height[r]))
		if s > maxA {
			maxA = s
		}
	}

	return maxA
}
