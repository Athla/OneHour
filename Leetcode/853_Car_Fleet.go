package leetcode

import (
	"sort"
)

type fleet struct {
	Pos, Speed int
}

func carFleet(target int, position []int, speed []int) int {
	pair := []fleet{}
	stack := []float32{}
	for i := range position {
		pair = append(pair, fleet{position[i], speed[i]})
	}

	sort.Slice(pair, func(i, j int) bool {
		return pair[i].Pos < pair[j].Pos
	})

	for i := len(pair) - 1; i >= 0; i-- {
		stack = append(stack, float32(target-pair[i].Pos)/float32(pair[i].Speed))
		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack)
}
