package leetcode

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	n := len(intervals)
	target := newInterval[0]
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if intervals[mid][0] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	result := make([][]int, 0)
	result = append(result, intervals[:left]...)
	result = append(result, newInterval)
	result = append(result, intervals[left:]...)

	res := make([][]int, 0)
	for _, interval := range result {
		if len(res) == 0 || res[len(res)-1][1] < interval[0] {
			res = append(res, interval)
		} else {
			if res[len(res)-1][1] < interval[1] {
				res[len(res)-1][1] = interval[1]
			}
		}
	}

	return res
}
