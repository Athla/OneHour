package leetcode

func maxProfit(prices []int) int {
	lo := 0
	hi := 1
	prof := 0

	for hi < len(prices) {
		if prices[hi] > prices[lo] {
			tmp := prices[hi] - prices[lo]
			if tmp > prof {
				prof = tmp
			}

		} else {
			lo = hi
		}
		hi += 1

	}

	return prof
}
