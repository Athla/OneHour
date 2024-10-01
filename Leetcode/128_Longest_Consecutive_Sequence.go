package leetcode

// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }

func longestConsecutive(nums []int) int {
	seen := make(map[int]struct{})
	long := 0

	for _, n := range nums {
		seen[n] = struct{}{}
	}

	for n := range seen {
		if _, ok := seen[n-1]; !ok {
			currSeq := 1
			for {
				if _, ok = seen[n+currSeq]; ok {
					currSeq++
					continue
				}
				long = max(long, currSeq)
				break
			}
		}
	}

	return long
}
