package leetcode

func dailyTemperatures(temp []int) []int {
	results := make([]int, len(temp))
	stack := make([]int, 0)

	for i, v := range temp {
		for len(stack) > 0 && temp[stack[len(stack)-1]] < v {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			results[idx] = i - idx
		}

		stack = append(stack, i)
	}

	return results
}
