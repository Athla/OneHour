# Rest is a right
- The fifth day of the streak was a rough one for me. I was tired.
- Not physically tired, but the type of tired where you simply can't do shit. Still, I got some thing done!
- Altough not even close to the previous days, I took it to solve and understand the [3Sum](https://leetcode.com/problems/3sum/) problem.
- This one was kinda of new, yet familiar, for me. Mainly because I've solved the TwoSum Ordered, which helped a lot.
- The core part of it was to break down the problem into two steps, to pick the current "key" value that, with the sum of other two values, is nullified and then proceed to do ordered two sum on it. 
- I wasn't able to do it ontime, but got pretty close, then I proceeded to see the editorial and got it fairly easy.
- And I used sort in this case because it looked fairly easy, however, considering constraints of the problem, it would've been better to, probably, use a MergeSort in this case. But hey! I got it done.

### Solution
```go
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
			switch{
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
```

- After that I tried to crunch some more but simply couldn't, so I took the rest of the day to rest. And man it was worth it.
