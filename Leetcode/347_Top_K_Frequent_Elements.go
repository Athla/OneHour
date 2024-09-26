package leetcode

import "sort"

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
