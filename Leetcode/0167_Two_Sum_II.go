package leetcode

func twoSumTwo(numbers []int, target int) []int {
	lo, hi := 0, len(numbers)-1

	for lo < hi {
		sum := numbers[lo] + numbers[hi]
		if sum > target {
			hi--
		} else if sum < target {
			lo++
		} else {
			return []int{lo + 1, hi + 1}
		}
	}

	return []int{}
}
