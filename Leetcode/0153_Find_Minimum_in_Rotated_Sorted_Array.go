package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMin(nums []int) int {
	res := nums[0]
	l, r := 0, len(nums)-1

	for l <= r {
		if nums[l] < nums[r] {
			res = min(nums[l], res)
		}
		m := (l + r) / 2
		res = min(nums[m], res)
		if nums[m] >= nums[l] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return res
}
