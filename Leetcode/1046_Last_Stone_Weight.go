package leetcode

import "sort"

func Abs(x, y int) int {
	ans := x - y
	if ans < 0 {
		return -ans
	}
	return ans
}

func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)
		x, y := stones[len(stones)-1], stones[len(stones)-2]
		if x == y {
			stones = stones[:len(stones)-2]
			continue
		} else {
			stones = stones[:len(stones)-2]
			stones = append(stones, Abs(x, y))
			continue
		}
	}
	if len(stones) == 0 {
		return 0
	}

	return stones[0]
}
