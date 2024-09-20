package leetcode

// Set top and bot
// While bot <= bot {
// Calc the middle point := (top + bot) / 2
//
//	do the checks {
//		if middle index == target val -> return mid index
//		if mid > target -> reduce top window.
//		else -> increment the bottom window.
//		}
//	}
func search(nums []int, target int) int {
	top := len(nums) - 1
	bot := 0
	for bot <= top {
		mid := (top + bot) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			top = mid - 1
		} else { // nums[mid] < targer
			bot = mid + 1
		}
	}

	return -1
}
