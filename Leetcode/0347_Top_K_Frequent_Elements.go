package leetcode

import "sort"

// Without bucketing
func topKFrequent(nums []int, k int) []int {
	seen := make(map[int]int)

	for _, n := range nums {
		seen[n]++
	}

	cArr := make([][]int, 0, len(seen))

	for i, v := range seen {
		cArr = append(cArr, []int{i, v})
	}
	sort.Slice(cArr, func(i, j int) bool {
		return cArr[i][1] > cArr[j][1]
	})

	var res []int

	for i := 0; i < k; i++ {
		res = append(res, cArr[i][0])
	}

	return res
}

// With bucketing
func TopKFrequent(nums []int, k int) []int {
	seen := make(map[int]int)

	for _, val := range nums {
		seen[val]++
	}

	buck := make([][]int, len(nums)+1)

	for k, v := range seen {
		buck[v] = append(buck[v], k)
	}

	ans := make([]int, 0, k)

	for i := len(buck) - 1; i >= 0; i-- {
		for _, val := range buck[i] {
			if k > 0 {
				ans = append(ans, val)
				k--
			}
		}
	}

	return ans
}
