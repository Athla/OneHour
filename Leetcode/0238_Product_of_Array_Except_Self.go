package leetcode

func productExceptSelf(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	total := make([]int, n)

	left[0] = 1
	for i := 1; i < len(left); i++ {
		left[i] = left[i-1] * nums[i-1]

	}

	right[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		total[i] = left[i] * right[i]
	}
	return total
}
