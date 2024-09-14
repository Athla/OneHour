package leetcode

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, v := range nums {
		diff := target - v
		if val, ok := seen[diff]; ok == true {
			return []int{val, i}
		} else {
			seen[v] = i
		}

	}
	return []int{}
}
