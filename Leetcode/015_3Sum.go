package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	var total int
	res := [][]int{}
	sort.Ints(nums)

	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}

		curr := v
		l, r := i+1, len(nums)-1
		for l < r {
			total = curr + (nums[l] + nums[r])
			switch {
			case total > 0:
				r--
			case total < 0:
				l++
			default:
				res = append(res, []int{v, nums[l], nums[r]})
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}
	}
	return res
}
