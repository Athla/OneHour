package leetcode

//	func max(a, b int) int {
//	    if a > b {
//	        return a
//	    }
//	    return b
//	}
func lengthOfLongestSubstring(s string) int {
	chars := make(map[rune]int)
	longest := 0
	start := 0

	for i, v := range s {
		if last, ok := chars[v]; ok && last >= start {
			start = last + 1
		}
		longest = max(longest, i-start+1)
		chars[v] = i
	}

	return longest
}
