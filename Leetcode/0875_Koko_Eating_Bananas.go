package leetcode

import "math"

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func minEatingSpeed(piles []int, h int) int {
	check := func(piles []int, speed, limit int) bool {
		var hours float64
		for _, bananas := range piles {
			hours += math.Ceil(float64(bananas) / float64(speed))
		}
		return int(hours) <= limit
	}

	left := 1
	right := 0
	for _, curr := range piles {
		right = max(right, curr)
	}
	for left <= right {
		mid := (left + right) / 2
		if check(piles, mid, h) {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	return left
}
