package leetcode

func CombinationSum(candidates []int, target int) [][]int {
	var locDfs func(i int, curr []int, total int)
	res := make([][]int, 0)
	locDfs = func(i int, curr []int, total int) {
		if total == target {
			temp := make([]int, len(curr))
			copy(curr, temp)
			res = append(res, temp)
			return
		}
		if i >= len(candidates) || total > target {
			return
		}

		curr = append(curr, candidates[i])
		locDfs(i, curr, total+candidates[i])
		curr = curr[:len(curr)-1]
		locDfs(i+1, curr, total)

	}

	locDfs(0, []int{}, 0)

	return res
}
